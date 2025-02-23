package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"google.golang.org/grpc"

	"go.pedge.io/dockervolume"
	"go.pedge.io/env"
	"go.pedge.io/pkg/map"

	"github.com/pachyderm/pachyderm/src/pfs"
	"github.com/pachyderm/pachyderm/src/pfs/fuse"
	"github.com/pachyderm/pachyderm/src/pkg/uuid"
)

const (
	defaultShard   = 0
	defaultModulus = 1
)

type appEnv struct {
	PachydermPfsd1Port string `env:"PACHYDERM_PFSD_1_PORT"`
	PfsAddress         string `env:"PFS_ADDRESS,default=0.0.0.0:650"`
	BaseMountpoint     string `env:"BASE_MOUNTPOINT,default=/tmp/pfs-volume-driver"`
	GRPCPort           int    `env:"GRPC_PORT,default=2150"`
	HTTPPort           int    `env:"HTTP_PORT,default=1950"`
	VolumeDriverName   string `env:"VOLUME_DRIVER_NAME,default=pfs"`
}

func main() {
	env.Main(do, &appEnv{})
}

func do(appEnvObj interface{}) error {
	server, err := newServer(appEnvObj.(*appEnv))
	if err != nil {
		return err
	}
	return server.Serve()
}

func newServer(appEnv *appEnv) (dockervolume.Server, error) {
	clientConn, err := grpc.Dial(getPFSAddress(appEnv), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	apiClient := pfs.NewAPIClient(clientConn)
	mounter := fuse.NewMounter(getPFSAddress(appEnv), apiClient)
	return dockervolume.NewTCPServer(
		newVolumeDriver(
			mounter,
			appEnv.BaseMountpoint,
		),
		appEnv.VolumeDriverName,
		fmt.Sprintf(":%d", appEnv.HTTPPort),
		dockervolume.ServerOptions{
			GRPCPort: uint16(appEnv.GRPCPort),
		},
	), nil
}

func getPFSAddress(appEnv *appEnv) string {
	address := appEnv.PachydermPfsd1Port
	if address == "" {
		return appEnv.PfsAddress
	}
	return strings.Replace(address, "tcp://", "", -1)
}

type volumeDriver struct {
	mounter        fuse.Mounter
	baseMountpoint string
}

func newVolumeDriver(
	mounter fuse.Mounter,
	baseMountpoint string,
) *volumeDriver {
	return &volumeDriver{
		mounter,
		baseMountpoint,
	}
}

func (v *volumeDriver) Create(_ string, _ pkgmap.StringStringMap) error {
	return nil
}

func (v *volumeDriver) Remove(_ string, _ pkgmap.StringStringMap, _ string) error {
	return nil
}

func (v *volumeDriver) Mount(name string, opts pkgmap.StringStringMap) (string, error) {
	mount, err := getMount(opts, v.baseMountpoint)
	if err != nil {
		return "", err
	}
	if err := mount.init(); err != nil {
		return "", err
	}
	if err := v.mounter.Mount(
		mount.mountpoint,
		mount.shard,
		mount.modulus,
		nil,
	); err != nil {
		return "", err
	}
	return mount.mountpoint, nil
}

func (v *volumeDriver) Unmount(_ string, _ pkgmap.StringStringMap, mountpoint string) error {
	return v.mounter.Unmount(mountpoint)
}

type mount struct {
	repository string
	commitID   string
	shard      uint64
	modulus    uint64
	mountpoint string
}

func getMount(opts pkgmap.StringStringMap, baseMountpoint string) (*mount, error) {
	repository, err := opts.GetRequiredString("repository")
	if err != nil {
		return nil, err
	}
	commitID, err := opts.GetRequiredString("commit_id")
	if err != nil {
		return nil, err
	}
	shard, err := opts.GetUint64("shard")
	if err != nil {
		return nil, err
	}
	if shard == 0 {
		shard = defaultShard
	}
	modulus, err := opts.GetUint64("modulus")
	if err != nil {
		return nil, err
	}
	if modulus == 0 {
		modulus = defaultModulus
	}
	return &mount{
		repository,
		commitID,
		shard,
		modulus,
		filepath.Join(baseMountpoint, fmt.Sprintf("%s-%s-%d-%d-%s", repository, commitID, shard, modulus, uuid.New())),
	}, nil
}

func (m *mount) init() error {
	return os.MkdirAll(m.mountpoint, 0777)
}
