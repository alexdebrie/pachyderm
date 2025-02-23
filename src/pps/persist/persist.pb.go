// Code generated by protoc-gen-go.
// source: pps/persist/persist.proto
// DO NOT EDIT!

/*
Package persist is a generated protocol buffer package.

It is generated from these files:
	pps/persist/persist.proto

It has these top-level messages:
	JobInfo
	JobInfos
	JobStatus
	JobStatuses
	JobOutput
	JobLog
	JobLogs
	PipelineInfo
	PipelineInfos
*/
package persist

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/google-protobuf"
import google_protobuf1 "go.pedge.io/google-protobuf"
import pfs "github.com/pachyderm/pachyderm/src/pfs"
import pachyderm_pps "github.com/pachyderm/pachyderm/src/pps"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type JobInfo struct {
	JobId string `protobuf:"bytes,1,opt,name=job_id" json:"job_id,omitempty"`
	// Types that are valid to be assigned to Spec:
	//	*JobInfo_Transform
	//	*JobInfo_PipelineName
	Spec         isJobInfo_Spec              `protobuf_oneof:"spec"`
	Input        *pfs.Commit                 `protobuf:"bytes,4,opt,name=input" json:"input,omitempty"`
	OutputParent *pfs.Commit                 `protobuf:"bytes,5,opt,name=output_parent" json:"output_parent,omitempty"`
	CreatedAt    *google_protobuf1.Timestamp `protobuf:"bytes,6,opt,name=created_at" json:"created_at,omitempty"`
}

func (m *JobInfo) Reset()         { *m = JobInfo{} }
func (m *JobInfo) String() string { return proto.CompactTextString(m) }
func (*JobInfo) ProtoMessage()    {}

type isJobInfo_Spec interface {
	isJobInfo_Spec()
}

type JobInfo_Transform struct {
	Transform *pachyderm_pps.Transform `protobuf:"bytes,2,opt,name=transform,oneof"`
}
type JobInfo_PipelineName struct {
	PipelineName string `protobuf:"bytes,3,opt,name=pipeline_name,oneof"`
}

func (*JobInfo_Transform) isJobInfo_Spec()    {}
func (*JobInfo_PipelineName) isJobInfo_Spec() {}

func (m *JobInfo) GetSpec() isJobInfo_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *JobInfo) GetTransform() *pachyderm_pps.Transform {
	if x, ok := m.GetSpec().(*JobInfo_Transform); ok {
		return x.Transform
	}
	return nil
}

func (m *JobInfo) GetPipelineName() string {
	if x, ok := m.GetSpec().(*JobInfo_PipelineName); ok {
		return x.PipelineName
	}
	return ""
}

func (m *JobInfo) GetInput() *pfs.Commit {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *JobInfo) GetOutputParent() *pfs.Commit {
	if m != nil {
		return m.OutputParent
	}
	return nil
}

func (m *JobInfo) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*JobInfo) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _JobInfo_OneofMarshaler, _JobInfo_OneofUnmarshaler, []interface{}{
		(*JobInfo_Transform)(nil),
		(*JobInfo_PipelineName)(nil),
	}
}

func _JobInfo_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*JobInfo)
	// spec
	switch x := m.Spec.(type) {
	case *JobInfo_Transform:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Transform); err != nil {
			return err
		}
	case *JobInfo_PipelineName:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.PipelineName)
	case nil:
	default:
		return fmt.Errorf("JobInfo.Spec has unexpected type %T", x)
	}
	return nil
}

func _JobInfo_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*JobInfo)
	switch tag {
	case 2: // spec.transform
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(pachyderm_pps.Transform)
		err := b.DecodeMessage(msg)
		m.Spec = &JobInfo_Transform{msg}
		return true, err
	case 3: // spec.pipeline_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Spec = &JobInfo_PipelineName{x}
		return true, err
	default:
		return false, nil
	}
}

type JobInfos struct {
	JobInfo []*JobInfo `protobuf:"bytes,1,rep,name=job_info" json:"job_info,omitempty"`
}

func (m *JobInfos) Reset()         { *m = JobInfos{} }
func (m *JobInfos) String() string { return proto.CompactTextString(m) }
func (*JobInfos) ProtoMessage()    {}

func (m *JobInfos) GetJobInfo() []*JobInfo {
	if m != nil {
		return m.JobInfo
	}
	return nil
}

type JobStatus struct {
	Id        string                      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	JobId     string                      `protobuf:"bytes,2,opt,name=job_id" json:"job_id,omitempty"`
	Type      pachyderm_pps.JobStatusType `protobuf:"varint,3,opt,name=type,enum=pachyderm.pps.JobStatusType" json:"type,omitempty"`
	Timestamp *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Message   string                      `protobuf:"bytes,5,opt,name=message" json:"message,omitempty"`
}

func (m *JobStatus) Reset()         { *m = JobStatus{} }
func (m *JobStatus) String() string { return proto.CompactTextString(m) }
func (*JobStatus) ProtoMessage()    {}

func (m *JobStatus) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type JobStatuses struct {
	JobStatus []*JobStatus `protobuf:"bytes,1,rep,name=job_status" json:"job_status,omitempty"`
}

func (m *JobStatuses) Reset()         { *m = JobStatuses{} }
func (m *JobStatuses) String() string { return proto.CompactTextString(m) }
func (*JobStatuses) ProtoMessage()    {}

func (m *JobStatuses) GetJobStatus() []*JobStatus {
	if m != nil {
		return m.JobStatus
	}
	return nil
}

// maybe name JobOutputCommit? we name the fields
// input and output, not input_commit and output_commit
type JobOutput struct {
	// if we wanted to be able to have multuple output pfs repos,
	// we would need JobOutputCommit to have a separate id
	JobId  string      `protobuf:"bytes,1,opt,name=job_id" json:"job_id,omitempty"`
	Output *pfs.Commit `protobuf:"bytes,2,opt,name=output" json:"output,omitempty"`
}

func (m *JobOutput) Reset()         { *m = JobOutput{} }
func (m *JobOutput) String() string { return proto.CompactTextString(m) }
func (*JobOutput) ProtoMessage()    {}

func (m *JobOutput) GetOutput() *pfs.Commit {
	if m != nil {
		return m.Output
	}
	return nil
}

type JobLog struct {
	Id           string                      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	JobId        string                      `protobuf:"bytes,2,opt,name=job_id" json:"job_id,omitempty"`
	Timestamp    *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	OutputStream pachyderm_pps.OutputStream  `protobuf:"varint,4,opt,name=output_stream,enum=pachyderm.pps.OutputStream" json:"output_stream,omitempty"`
	Value        []byte                      `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *JobLog) Reset()         { *m = JobLog{} }
func (m *JobLog) String() string { return proto.CompactTextString(m) }
func (*JobLog) ProtoMessage()    {}

func (m *JobLog) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type JobLogs struct {
	JobLog []*JobLog `protobuf:"bytes,1,rep,name=job_log" json:"job_log,omitempty"`
}

func (m *JobLogs) Reset()         { *m = JobLogs{} }
func (m *JobLogs) String() string { return proto.CompactTextString(m) }
func (*JobLogs) ProtoMessage()    {}

func (m *JobLogs) GetJobLog() []*JobLog {
	if m != nil {
		return m.JobLog
	}
	return nil
}

type PipelineInfo struct {
	PipelineName string                      `protobuf:"bytes,1,opt,name=pipeline_name" json:"pipeline_name,omitempty"`
	Transform    *pachyderm_pps.Transform    `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	Input        *pfs.Repo                   `protobuf:"bytes,3,opt,name=input" json:"input,omitempty"`
	Output       *pfs.Repo                   `protobuf:"bytes,4,opt,name=output" json:"output,omitempty"`
	CreatedAt    *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=created_at" json:"created_at,omitempty"`
}

func (m *PipelineInfo) Reset()         { *m = PipelineInfo{} }
func (m *PipelineInfo) String() string { return proto.CompactTextString(m) }
func (*PipelineInfo) ProtoMessage()    {}

func (m *PipelineInfo) GetTransform() *pachyderm_pps.Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *PipelineInfo) GetInput() *pfs.Repo {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *PipelineInfo) GetOutput() *pfs.Repo {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *PipelineInfo) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type PipelineInfos struct {
	PipelineInfo []*PipelineInfo `protobuf:"bytes,1,rep,name=pipeline_info" json:"pipeline_info,omitempty"`
}

func (m *PipelineInfos) Reset()         { *m = PipelineInfos{} }
func (m *PipelineInfos) String() string { return proto.CompactTextString(m) }
func (*PipelineInfos) ProtoMessage()    {}

func (m *PipelineInfos) GetPipelineInfo() []*PipelineInfo {
	if m != nil {
		return m.PipelineInfo
	}
	return nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for API service

type APIClient interface {
	// job_id cannot be set
	// timestamp cannot be set
	CreateJobInfo(ctx context.Context, in *JobInfo, opts ...grpc.CallOption) (*JobInfo, error)
	GetJobInfo(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobInfo, error)
	// ordered by time, latest to earliest
	ListJobInfos(ctx context.Context, in *pachyderm_pps.ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error)
	// should only be called when rolling back if a Job does not start!
	DeleteJobInfo(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// id cannot be set
	// timestamp cannot be set
	CreateJobStatus(ctx context.Context, in *JobStatus, opts ...grpc.CallOption) (*JobStatus, error)
	// ordered by time, latest to earliest
	GetJobStatuses(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobStatuses, error)
	CreateJobOutput(ctx context.Context, in *JobOutput, opts ...grpc.CallOption) (*JobOutput, error)
	GetJobOutput(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobOutput, error)
	// id cannot be set
	CreateJobLog(ctx context.Context, in *JobLog, opts ...grpc.CallOption) (*JobLog, error)
	// ordered by time, latest to earliest
	GetJobLogs(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobLogs, error)
	// timestamp cannot be set
	CreatePipelineInfo(ctx context.Context, in *PipelineInfo, opts ...grpc.CallOption) (*PipelineInfo, error)
	GetPipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*PipelineInfo, error)
	// ordered by time, latest to earliest
	ListPipelineInfos(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PipelineInfos, error)
	DeletePipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) CreateJobInfo(ctx context.Context, in *JobInfo, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreateJobInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetJobInfo(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/GetJobInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListJobInfos(ctx context.Context, in *pachyderm_pps.ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error) {
	out := new(JobInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/ListJobInfos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeleteJobInfo(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/DeleteJobInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateJobStatus(ctx context.Context, in *JobStatus, opts ...grpc.CallOption) (*JobStatus, error) {
	out := new(JobStatus)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreateJobStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetJobStatuses(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobStatuses, error) {
	out := new(JobStatuses)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/GetJobStatuses", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateJobOutput(ctx context.Context, in *JobOutput, opts ...grpc.CallOption) (*JobOutput, error) {
	out := new(JobOutput)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreateJobOutput", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetJobOutput(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobOutput, error) {
	out := new(JobOutput)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/GetJobOutput", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateJobLog(ctx context.Context, in *JobLog, opts ...grpc.CallOption) (*JobLog, error) {
	out := new(JobLog)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreateJobLog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetJobLogs(ctx context.Context, in *pachyderm_pps.Job, opts ...grpc.CallOption) (*JobLogs, error) {
	out := new(JobLogs)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/GetJobLogs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreatePipelineInfo(ctx context.Context, in *PipelineInfo, opts ...grpc.CallOption) (*PipelineInfo, error) {
	out := new(PipelineInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/CreatePipelineInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetPipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*PipelineInfo, error) {
	out := new(PipelineInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/GetPipelineInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListPipelineInfos(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PipelineInfos, error) {
	out := new(PipelineInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/ListPipelineInfos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeletePipelineInfo(ctx context.Context, in *pachyderm_pps.Pipeline, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.persist.API/DeletePipelineInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	// job_id cannot be set
	// timestamp cannot be set
	CreateJobInfo(context.Context, *JobInfo) (*JobInfo, error)
	GetJobInfo(context.Context, *pachyderm_pps.Job) (*JobInfo, error)
	// ordered by time, latest to earliest
	ListJobInfos(context.Context, *pachyderm_pps.ListJobRequest) (*JobInfos, error)
	// should only be called when rolling back if a Job does not start!
	DeleteJobInfo(context.Context, *pachyderm_pps.Job) (*google_protobuf.Empty, error)
	// id cannot be set
	// timestamp cannot be set
	CreateJobStatus(context.Context, *JobStatus) (*JobStatus, error)
	// ordered by time, latest to earliest
	GetJobStatuses(context.Context, *pachyderm_pps.Job) (*JobStatuses, error)
	CreateJobOutput(context.Context, *JobOutput) (*JobOutput, error)
	GetJobOutput(context.Context, *pachyderm_pps.Job) (*JobOutput, error)
	// id cannot be set
	CreateJobLog(context.Context, *JobLog) (*JobLog, error)
	// ordered by time, latest to earliest
	GetJobLogs(context.Context, *pachyderm_pps.Job) (*JobLogs, error)
	// timestamp cannot be set
	CreatePipelineInfo(context.Context, *PipelineInfo) (*PipelineInfo, error)
	GetPipelineInfo(context.Context, *pachyderm_pps.Pipeline) (*PipelineInfo, error)
	// ordered by time, latest to earliest
	ListPipelineInfos(context.Context, *google_protobuf.Empty) (*PipelineInfos, error)
	DeletePipelineInfo(context.Context, *pachyderm_pps.Pipeline) (*google_protobuf.Empty, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_CreateJobInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(JobInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreateJobInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetJobInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetJobInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_ListJobInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.ListJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).ListJobInfos(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_DeleteJobInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).DeleteJobInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_CreateJobStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(JobStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreateJobStatus(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetJobStatuses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetJobStatuses(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_CreateJobOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(JobOutput)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreateJobOutput(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetJobOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetJobOutput(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_CreateJobLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(JobLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreateJobLog(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetJobLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetJobLogs(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_CreatePipelineInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(PipelineInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).CreatePipelineInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_GetPipelineInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Pipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).GetPipelineInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_ListPipelineInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).ListPipelineInfos(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_DeletePipelineInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(pachyderm_pps.Pipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).DeletePipelineInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pachyderm.pps.persist.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJobInfo",
			Handler:    _API_CreateJobInfo_Handler,
		},
		{
			MethodName: "GetJobInfo",
			Handler:    _API_GetJobInfo_Handler,
		},
		{
			MethodName: "ListJobInfos",
			Handler:    _API_ListJobInfos_Handler,
		},
		{
			MethodName: "DeleteJobInfo",
			Handler:    _API_DeleteJobInfo_Handler,
		},
		{
			MethodName: "CreateJobStatus",
			Handler:    _API_CreateJobStatus_Handler,
		},
		{
			MethodName: "GetJobStatuses",
			Handler:    _API_GetJobStatuses_Handler,
		},
		{
			MethodName: "CreateJobOutput",
			Handler:    _API_CreateJobOutput_Handler,
		},
		{
			MethodName: "GetJobOutput",
			Handler:    _API_GetJobOutput_Handler,
		},
		{
			MethodName: "CreateJobLog",
			Handler:    _API_CreateJobLog_Handler,
		},
		{
			MethodName: "GetJobLogs",
			Handler:    _API_GetJobLogs_Handler,
		},
		{
			MethodName: "CreatePipelineInfo",
			Handler:    _API_CreatePipelineInfo_Handler,
		},
		{
			MethodName: "GetPipelineInfo",
			Handler:    _API_GetPipelineInfo_Handler,
		},
		{
			MethodName: "ListPipelineInfos",
			Handler:    _API_ListPipelineInfos_Handler,
		},
		{
			MethodName: "DeletePipelineInfo",
			Handler:    _API_DeletePipelineInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
