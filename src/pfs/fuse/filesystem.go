package fuse

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strings"
	"sync/atomic"

	"go.pedge.io/protolog"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/pachyderm/pachyderm/src/pfs"
	"github.com/pachyderm/pachyderm/src/pfs/pfsutil"
	"golang.org/x/net/context"
)

type filesystem struct {
	apiClient pfs.APIClient
	Filesystem
}

func newFilesystem(
	apiClient pfs.APIClient,
	shard uint64,
	modulus uint64,
) *filesystem {
	return &filesystem{
		apiClient,
		Filesystem{
			shard,
			modulus,
		},
	}
}

func (f *filesystem) Root() (result fs.Node, retErr error) {
	defer func() {
		protolog.Info(&Root{&f.Filesystem, getNode(result), errorToString(retErr)})
	}()
	return &directory{f, Node{"", "", "", true}}, nil
}

type directory struct {
	fs *filesystem
	Node
}

func (d *directory) Attr(ctx context.Context, a *fuse.Attr) (retErr error) {
	defer func() {
		protolog.Info(&DirectoryAttr{&d.Node, &Attr{uint32(a.Mode)}, errorToString(retErr)})
	}()
	if d.Write {
		a.Mode = os.ModeDir | 0775
	} else {
		a.Mode = os.ModeDir | 0555
	}
	return nil
}

func (d *directory) Lookup(ctx context.Context, name string) (result fs.Node, retErr error) {
	defer func() {
		protolog.Info(&DirectoryLookup{&d.Node, name, getNode(result), errorToString(retErr)})
	}()
	if d.RepoName == "" {
		return d.lookUpRepo(ctx, name)
	}
	if d.CommitID == "" {
		return d.lookUpCommit(ctx, name)
	}
	return d.lookUpFile(ctx, name)
}

func (d *directory) ReadDirAll(ctx context.Context) (result []fuse.Dirent, retErr error) {
	defer func() {
		var dirents []*Dirent
		for _, dirent := range result {
			dirents = append(dirents, &Dirent{dirent.Inode, dirent.Name})
		}
		protolog.Info(&DirectoryReadDirAll{&d.Node, dirents, errorToString(retErr)})
	}()
	if d.RepoName == "" {
		return d.readRepos(ctx)
	}
	if d.CommitID == "" {
		return d.readCommits(ctx)
	}
	return d.readFiles(ctx)
}

func (d *directory) Create(ctx context.Context, request *fuse.CreateRequest, response *fuse.CreateResponse) (result fs.Node, _ fs.Handle, retErr error) {
	defer func() {
		protolog.Info(&DirectoryCreate{&d.Node, getNode(result), errorToString(retErr)})
	}()
	if d.CommitID == "" {
		return nil, 0, fuse.EPERM
	}
	directory := *d
	directory.Path = path.Join(directory.Path, request.Name)
	localResult := &file{directory, 0, 0}
	handle, err := localResult.Open(ctx, nil, nil)
	if err != nil {
		return nil, nil, err
	}
	return localResult, handle, nil
}

func (d *directory) Mkdir(ctx context.Context, request *fuse.MkdirRequest) (result fs.Node, retErr error) {
	defer func() {
		protolog.Info(&DirectoryMkdir{&d.Node, getNode(result), errorToString(retErr)})
	}()
	if d.CommitID == "" {
		return nil, fuse.EPERM
	}
	if err := pfsutil.MakeDirectory(d.fs.apiClient, d.RepoName, d.CommitID, path.Join(d.Path, request.Name)); err != nil {
		return nil, err
	}
	localResult := *d
	localResult.Path = path.Join(localResult.Path, request.Name)
	return &localResult, nil
}

type file struct {
	directory
	handles int32
	size    int64
}

func (f *file) Attr(ctx context.Context, a *fuse.Attr) (retErr error) {
	defer func() {
		protolog.Info(&FileAttr{&f.Node, &Attr{uint32(a.Mode)}, errorToString(retErr)})
	}()
	fileInfo, err := pfsutil.InspectFile(
		f.fs.apiClient,
		f.RepoName,
		f.CommitID,
		f.Path,
	)
	if err != nil {
		return err
	}
	if fileInfo != nil {
		a.Size = fileInfo.SizeBytes
	}
	a.Mode = 0666
	return nil
}

func (f *file) Read(ctx context.Context, request *fuse.ReadRequest, response *fuse.ReadResponse) (retErr error) {
	defer func() {
		protolog.Info(&FileRead{&f.Node, errorToString(retErr)})
	}()
	buffer := bytes.NewBuffer(make([]byte, 0, request.Size))
	if err := pfsutil.GetFile(f.fs.apiClient, f.RepoName, f.CommitID, f.Path, request.Offset, int64(request.Size), buffer); err != nil {
		return err
	}
	response.Data = buffer.Bytes()
	return nil
}

func (f *file) Open(ctx context.Context, request *fuse.OpenRequest, response *fuse.OpenResponse) (_ fs.Handle, retErr error) {
	defer func() {
		protolog.Info(&FileRead{&f.Node, errorToString(retErr)})
	}()
	atomic.AddInt32(&f.handles, 1)
	return f, nil
}

func (f *file) Write(ctx context.Context, request *fuse.WriteRequest, response *fuse.WriteResponse) (retErr error) {
	defer func() {
		protolog.Info(&FileWrite{&f.Node, errorToString(retErr)})
	}()
	written, err := pfsutil.PutFile(f.fs.apiClient, f.RepoName, f.CommitID, f.Path, request.Offset, bytes.NewReader(request.Data))
	if err != nil {
		return err
	}
	response.Size = written
	if f.size < request.Offset+int64(written) {
		f.size = request.Offset + int64(written)
	}
	return nil
}

func (d *directory) lookUpRepo(ctx context.Context, name string) (fs.Node, error) {
	repoInfo, err := pfsutil.InspectRepo(d.fs.apiClient, name)
	if err != nil {
		return nil, err
	}
	if repoInfo == nil {
		return nil, fuse.ENOENT
	}
	result := *d
	result.RepoName = name
	return &result, nil
}

func (d *directory) lookUpCommit(ctx context.Context, name string) (fs.Node, error) {
	commitInfo, err := pfsutil.InspectCommit(
		d.fs.apiClient,
		d.RepoName,
		name,
	)
	if err != nil {
		return nil, err
	}
	if commitInfo == nil {
		return nil, fuse.ENOENT
	}
	result := *d
	result.CommitID = name
	return &result, nil
}

func (d *directory) lookUpFile(ctx context.Context, name string) (fs.Node, error) {
	fileInfo, err := pfsutil.InspectFile(
		d.fs.apiClient,
		d.RepoName,
		d.CommitID,
		path.Join(d.Path, name),
	)
	if err != nil {
		return nil, err
	}
	if fileInfo.FileType == pfs.FileType_FILE_TYPE_NONE {
		return nil, fuse.ENOENT
	}
	directory := *d
	directory.Path = fileInfo.File.Path
	switch fileInfo.FileType {
	case pfs.FileType_FILE_TYPE_REGULAR:
		directory.Path = fileInfo.File.Path
		return &file{
			directory,
			0,
			int64(fileInfo.SizeBytes),
		}, nil
	case pfs.FileType_FILE_TYPE_DIR:
		return &directory, nil
	default:
		return nil, fmt.Errorf("Unrecognized FileType.")
	}
}

func (d *directory) readRepos(ctx context.Context) ([]fuse.Dirent, error) {
	repoInfos, err := pfsutil.ListRepo(d.fs.apiClient)
	if err != nil {
		return nil, err
	}
	var result []fuse.Dirent
	for _, repoInfo := range repoInfos {
		result = append(result, fuse.Dirent{Name: repoInfo.Repo.Name, Type: fuse.DT_Dir})
	}
	return result, nil
}

func (d *directory) readCommits(ctx context.Context) ([]fuse.Dirent, error) {
	commitInfos, err := pfsutil.ListCommit(d.fs.apiClient, d.RepoName)
	if err != nil {
		return nil, err
	}
	result := make([]fuse.Dirent, 0, len(commitInfos))
	for _, commitInfo := range commitInfos {
		result = append(result, fuse.Dirent{Name: commitInfo.Commit.Id, Type: fuse.DT_Dir})
	}
	return result, nil
}

func (d *directory) readFiles(ctx context.Context) ([]fuse.Dirent, error) {
	fileInfos, err := pfsutil.ListFile(d.fs.apiClient, d.RepoName, d.CommitID, d.Path, d.fs.Shard, d.fs.Modulus)
	if err != nil {
		return nil, err
	}
	var result []fuse.Dirent
	for _, fileInfo := range fileInfos {
		shortPath := strings.TrimPrefix(fileInfo.File.Path, d.Path)
		switch fileInfo.FileType {
		case pfs.FileType_FILE_TYPE_REGULAR:
			result = append(result, fuse.Dirent{Name: shortPath, Type: fuse.DT_File})
		case pfs.FileType_FILE_TYPE_DIR:
			result = append(result, fuse.Dirent{Name: shortPath, Type: fuse.DT_Dir})
		default:
			continue
		}
	}
	return result, nil
}

// TODO this code is duplicate elsewhere, we should put it somehwere.
func errorToString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func getNode(node fs.Node) *Node {
	switch n := node.(type) {
	default:
		return nil
	case *directory:
		return &n.Node
	case *file:
		return &n.Node
	}
}
