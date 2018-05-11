// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

/*
Package task_go_proto is a generated protocol buffer package.

It is generated from these files:
	task.proto

It has these top-level messages:
	Spec
	ListSpec
	ProcessListSpec
	CopySpec
	TaskReqMsg
	TaskRespMsg
	Log
	ListLog
	ProcessListLog
	CopyLog
*/
package task_go_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Specifies the task operation that a task performs.
type Type int32

const (
	Type_UNSET_TYPE   Type = 0
	Type_LIST         Type = 1
	Type_PROCESS_LIST Type = 2
	Type_COPY         Type = 3
)

var Type_name = map[int32]string{
	0: "UNSET_TYPE",
	1: "LIST",
	2: "PROCESS_LIST",
	3: "COPY",
}
var Type_value = map[string]int32{
	"UNSET_TYPE":   0,
	"LIST":         1,
	"PROCESS_LIST": 2,
	"COPY":         3,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Specifies the current status of the cloud ingest task.
type Status int32

const (
	Status_UNSET_STATUS Status = 0
	Status_UNQUEUED     Status = 1
	Status_QUEUED       Status = 2
	Status_FAILED       Status = 3
	Status_SUCCESS      Status = 4
)

var Status_name = map[int32]string{
	0: "UNSET_STATUS",
	1: "UNQUEUED",
	2: "QUEUED",
	3: "FAILED",
	4: "SUCCESS",
}
var Status_value = map[string]int32{
	"UNSET_STATUS": 0,
	"UNQUEUED":     1,
	"QUEUED":       2,
	"FAILED":       3,
	"SUCCESS":      4,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Specifies the classes of failures that a task can have.
type FailureType int32

const (
	FailureType_UNSET_FAILURE_TYPE     FailureType = 0
	FailureType_UNKNOWN_FAILURE        FailureType = 1
	FailureType_FILE_MODIFIED_FAILURE  FailureType = 2
	FailureType_HASH_MISMATCH_FAILURE  FailureType = 3
	FailureType_PRECONDITION_FAILURE   FailureType = 4
	FailureType_FILE_NOT_FOUND_FAILURE FailureType = 5
	FailureType_PERMISSION_FAILURE     FailureType = 6
)

var FailureType_name = map[int32]string{
	0: "UNSET_FAILURE_TYPE",
	1: "UNKNOWN_FAILURE",
	2: "FILE_MODIFIED_FAILURE",
	3: "HASH_MISMATCH_FAILURE",
	4: "PRECONDITION_FAILURE",
	5: "FILE_NOT_FOUND_FAILURE",
	6: "PERMISSION_FAILURE",
}
var FailureType_value = map[string]int32{
	"UNSET_FAILURE_TYPE":     0,
	"UNKNOWN_FAILURE":        1,
	"FILE_MODIFIED_FAILURE":  2,
	"HASH_MISMATCH_FAILURE":  3,
	"PRECONDITION_FAILURE":   4,
	"FILE_NOT_FOUND_FAILURE": 5,
	"PERMISSION_FAILURE":     6,
}

func (x FailureType) String() string {
	return proto.EnumName(FailureType_name, int32(x))
}
func (FailureType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Contains information about a task. A task is a unit of work, one of:
// 1) listing the contents of a single directory
// 2) processing a list file
// 3) copying a single file
// Tasks might be incremental and require multiple request-response round trips
// to complete.
type Spec struct {
	// Types that are valid to be assigned to Spec:
	//	*Spec_ListSpec
	//	*Spec_ProcessListSpec
	//	*Spec_CopySpec
	Spec isSpec_Spec `protobuf_oneof:"spec"`
}

func (m *Spec) Reset()                    { *m = Spec{} }
func (m *Spec) String() string            { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()               {}
func (*Spec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isSpec_Spec interface {
	isSpec_Spec()
}

type Spec_ListSpec struct {
	ListSpec *ListSpec `protobuf:"bytes,1,opt,name=list_spec,json=listSpec,oneof"`
}
type Spec_ProcessListSpec struct {
	ProcessListSpec *ProcessListSpec `protobuf:"bytes,2,opt,name=process_list_spec,json=processListSpec,oneof"`
}
type Spec_CopySpec struct {
	CopySpec *CopySpec `protobuf:"bytes,3,opt,name=copy_spec,json=copySpec,oneof"`
}

func (*Spec_ListSpec) isSpec_Spec()        {}
func (*Spec_ProcessListSpec) isSpec_Spec() {}
func (*Spec_CopySpec) isSpec_Spec()        {}

func (m *Spec) GetSpec() isSpec_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Spec) GetListSpec() *ListSpec {
	if x, ok := m.GetSpec().(*Spec_ListSpec); ok {
		return x.ListSpec
	}
	return nil
}

func (m *Spec) GetProcessListSpec() *ProcessListSpec {
	if x, ok := m.GetSpec().(*Spec_ProcessListSpec); ok {
		return x.ProcessListSpec
	}
	return nil
}

func (m *Spec) GetCopySpec() *CopySpec {
	if x, ok := m.GetSpec().(*Spec_CopySpec); ok {
		return x.CopySpec
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Spec) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Spec_OneofMarshaler, _Spec_OneofUnmarshaler, _Spec_OneofSizer, []interface{}{
		(*Spec_ListSpec)(nil),
		(*Spec_ProcessListSpec)(nil),
		(*Spec_CopySpec)(nil),
	}
}

func _Spec_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Spec)
	// spec
	switch x := m.Spec.(type) {
	case *Spec_ListSpec:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ListSpec); err != nil {
			return err
		}
	case *Spec_ProcessListSpec:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProcessListSpec); err != nil {
			return err
		}
	case *Spec_CopySpec:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CopySpec); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Spec.Spec has unexpected type %T", x)
	}
	return nil
}

func _Spec_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Spec)
	switch tag {
	case 1: // spec.list_spec
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ListSpec)
		err := b.DecodeMessage(msg)
		m.Spec = &Spec_ListSpec{msg}
		return true, err
	case 2: // spec.process_list_spec
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProcessListSpec)
		err := b.DecodeMessage(msg)
		m.Spec = &Spec_ProcessListSpec{msg}
		return true, err
	case 3: // spec.copy_spec
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CopySpec)
		err := b.DecodeMessage(msg)
		m.Spec = &Spec_CopySpec{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Spec_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Spec)
	// spec
	switch x := m.Spec.(type) {
	case *Spec_ListSpec:
		s := proto.Size(x.ListSpec)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Spec_ProcessListSpec:
		s := proto.Size(x.ProcessListSpec)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Spec_CopySpec:
		s := proto.Size(x.CopySpec)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Contains the information about a list task. A list task is responsible for
// listing the contents of a directory.
type ListSpec struct {
	DstListResultBucket   string `protobuf:"bytes,1,opt,name=dst_list_result_bucket,json=dstListResultBucket" json:"dst_list_result_bucket,omitempty"`
	DstListResultObject   string `protobuf:"bytes,2,opt,name=dst_list_result_object,json=dstListResultObject" json:"dst_list_result_object,omitempty"`
	SrcDirectory          string `protobuf:"bytes,3,opt,name=src_directory,json=srcDirectory" json:"src_directory,omitempty"`
	ExpectedGenerationNum int64  `protobuf:"varint,4,opt,name=expected_generation_num,json=expectedGenerationNum" json:"expected_generation_num,omitempty"`
}

func (m *ListSpec) Reset()                    { *m = ListSpec{} }
func (m *ListSpec) String() string            { return proto.CompactTextString(m) }
func (*ListSpec) ProtoMessage()               {}
func (*ListSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListSpec) GetDstListResultBucket() string {
	if m != nil {
		return m.DstListResultBucket
	}
	return ""
}

func (m *ListSpec) GetDstListResultObject() string {
	if m != nil {
		return m.DstListResultObject
	}
	return ""
}

func (m *ListSpec) GetSrcDirectory() string {
	if m != nil {
		return m.SrcDirectory
	}
	return ""
}

func (m *ListSpec) GetExpectedGenerationNum() int64 {
	if m != nil {
		return m.ExpectedGenerationNum
	}
	return 0
}

// Contains the information about a process list task. A process list task is
// responsible for listing a small chunk of the list task.
type ProcessListSpec struct {
	DstListResultBucket string `protobuf:"bytes,1,opt,name=dst_list_result_bucket,json=dstListResultBucket" json:"dst_list_result_bucket,omitempty"`
	DstListResultObject string `protobuf:"bytes,2,opt,name=dst_list_result_object,json=dstListResultObject" json:"dst_list_result_object,omitempty"`
	SrcDirectory        string `protobuf:"bytes,3,opt,name=src_directory,json=srcDirectory" json:"src_directory,omitempty"`
	ByteOffset          int64  `protobuf:"varint,4,opt,name=byte_offset,json=byteOffset" json:"byte_offset,omitempty"`
}

func (m *ProcessListSpec) Reset()                    { *m = ProcessListSpec{} }
func (m *ProcessListSpec) String() string            { return proto.CompactTextString(m) }
func (*ProcessListSpec) ProtoMessage()               {}
func (*ProcessListSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ProcessListSpec) GetDstListResultBucket() string {
	if m != nil {
		return m.DstListResultBucket
	}
	return ""
}

func (m *ProcessListSpec) GetDstListResultObject() string {
	if m != nil {
		return m.DstListResultObject
	}
	return ""
}

func (m *ProcessListSpec) GetSrcDirectory() string {
	if m != nil {
		return m.SrcDirectory
	}
	return ""
}

func (m *ProcessListSpec) GetByteOffset() int64 {
	if m != nil {
		return m.ByteOffset
	}
	return 0
}

// Contains the information about a copy task. A copy task is responsible for
// copying a single file.
type CopySpec struct {
	SrcFile               string `protobuf:"bytes,1,opt,name=src_file,json=srcFile" json:"src_file,omitempty"`
	DstBucket             string `protobuf:"bytes,2,opt,name=dst_bucket,json=dstBucket" json:"dst_bucket,omitempty"`
	DstObject             string `protobuf:"bytes,3,opt,name=dst_object,json=dstObject" json:"dst_object,omitempty"`
	ExpectedGenerationNum int64  `protobuf:"varint,4,opt,name=expected_generation_num,json=expectedGenerationNum" json:"expected_generation_num,omitempty"`
	// Fields for bandwidth management.
	Bandwidth int64 `protobuf:"varint,5,opt,name=bandwidth" json:"bandwidth,omitempty"`
	// Fields only for managing resumable copies.
	FileBytes         int64  `protobuf:"varint,6,opt,name=file_bytes,json=fileBytes" json:"file_bytes,omitempty"`
	FileMTime         int64  `protobuf:"varint,7,opt,name=file_m_time,json=fileMTime" json:"file_m_time,omitempty"`
	BytesCopied       int64  `protobuf:"varint,8,opt,name=bytes_copied,json=bytesCopied" json:"bytes_copied,omitempty"`
	Crc32C            uint32 `protobuf:"varint,9,opt,name=crc32c" json:"crc32c,omitempty"`
	BytesToCopy       int64  `protobuf:"varint,10,opt,name=bytes_to_copy,json=bytesToCopy" json:"bytes_to_copy,omitempty"`
	ResumableUploadId string `protobuf:"bytes,11,opt,name=resumable_upload_id,json=resumableUploadId" json:"resumable_upload_id,omitempty"`
}

func (m *CopySpec) Reset()                    { *m = CopySpec{} }
func (m *CopySpec) String() string            { return proto.CompactTextString(m) }
func (*CopySpec) ProtoMessage()               {}
func (*CopySpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CopySpec) GetSrcFile() string {
	if m != nil {
		return m.SrcFile
	}
	return ""
}

func (m *CopySpec) GetDstBucket() string {
	if m != nil {
		return m.DstBucket
	}
	return ""
}

func (m *CopySpec) GetDstObject() string {
	if m != nil {
		return m.DstObject
	}
	return ""
}

func (m *CopySpec) GetExpectedGenerationNum() int64 {
	if m != nil {
		return m.ExpectedGenerationNum
	}
	return 0
}

func (m *CopySpec) GetBandwidth() int64 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *CopySpec) GetFileBytes() int64 {
	if m != nil {
		return m.FileBytes
	}
	return 0
}

func (m *CopySpec) GetFileMTime() int64 {
	if m != nil {
		return m.FileMTime
	}
	return 0
}

func (m *CopySpec) GetBytesCopied() int64 {
	if m != nil {
		return m.BytesCopied
	}
	return 0
}

func (m *CopySpec) GetCrc32C() uint32 {
	if m != nil {
		return m.Crc32C
	}
	return 0
}

func (m *CopySpec) GetBytesToCopy() int64 {
	if m != nil {
		return m.BytesToCopy
	}
	return 0
}

func (m *CopySpec) GetResumableUploadId() string {
	if m != nil {
		return m.ResumableUploadId
	}
	return ""
}

// Contains the message sent from the DCP to an Agent to issue a task request.
type TaskReqMsg struct {
	TaskRelRsrcName string `protobuf:"bytes,1,opt,name=task_rel_rsrc_name,json=taskRelRsrcName" json:"task_rel_rsrc_name,omitempty"`
	Spec            *Spec  `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
}

func (m *TaskReqMsg) Reset()                    { *m = TaskReqMsg{} }
func (m *TaskReqMsg) String() string            { return proto.CompactTextString(m) }
func (*TaskReqMsg) ProtoMessage()               {}
func (*TaskReqMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TaskReqMsg) GetTaskRelRsrcName() string {
	if m != nil {
		return m.TaskRelRsrcName
	}
	return ""
}

func (m *TaskReqMsg) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// Contains the message sent from the Agent to the DCP in response to a task
// request.
type TaskRespMsg struct {
	TaskRelRsrcName string      `protobuf:"bytes,1,opt,name=task_rel_rsrc_name,json=taskRelRsrcName" json:"task_rel_rsrc_name,omitempty"`
	Status          string      `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	FailureType     FailureType `protobuf:"varint,3,opt,name=failure_type,json=failureType,enum=cloud_ingest_task.FailureType" json:"failure_type,omitempty"`
	FailureMessage  string      `protobuf:"bytes,4,opt,name=failure_message,json=failureMessage" json:"failure_message,omitempty"`
	Log             *Log        `protobuf:"bytes,5,opt,name=log" json:"log,omitempty"`
	ReqSpec         *Spec       `protobuf:"bytes,6,opt,name=req_spec,json=reqSpec" json:"req_spec,omitempty"`
	RespSpec        *Spec       `protobuf:"bytes,7,opt,name=resp_spec,json=respSpec" json:"resp_spec,omitempty"`
}

func (m *TaskRespMsg) Reset()                    { *m = TaskRespMsg{} }
func (m *TaskRespMsg) String() string            { return proto.CompactTextString(m) }
func (*TaskRespMsg) ProtoMessage()               {}
func (*TaskRespMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TaskRespMsg) GetTaskRelRsrcName() string {
	if m != nil {
		return m.TaskRelRsrcName
	}
	return ""
}

func (m *TaskRespMsg) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TaskRespMsg) GetFailureType() FailureType {
	if m != nil {
		return m.FailureType
	}
	return FailureType_UNSET_FAILURE_TYPE
}

func (m *TaskRespMsg) GetFailureMessage() string {
	if m != nil {
		return m.FailureMessage
	}
	return ""
}

func (m *TaskRespMsg) GetLog() *Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *TaskRespMsg) GetReqSpec() *Spec {
	if m != nil {
		return m.ReqSpec
	}
	return nil
}

func (m *TaskRespMsg) GetRespSpec() *Spec {
	if m != nil {
		return m.RespSpec
	}
	return nil
}

// Contains log information for a task. This message is suitable for the "Log"
// field in the LogEntries Spanner table. Note that this info is eventually
// dumped into the user's GCS bucket.
type Log struct {
	// Types that are valid to be assigned to Log:
	//	*Log_ListLog
	//	*Log_ProcessListLog
	//	*Log_CopyLog
	Log isLog_Log `protobuf_oneof:"log"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isLog_Log interface {
	isLog_Log()
}

type Log_ListLog struct {
	ListLog *ListLog `protobuf:"bytes,1,opt,name=list_log,json=listLog,oneof"`
}
type Log_ProcessListLog struct {
	ProcessListLog *ProcessListLog `protobuf:"bytes,2,opt,name=process_list_log,json=processListLog,oneof"`
}
type Log_CopyLog struct {
	CopyLog *CopyLog `protobuf:"bytes,3,opt,name=copy_log,json=copyLog,oneof"`
}

func (*Log_ListLog) isLog_Log()        {}
func (*Log_ProcessListLog) isLog_Log() {}
func (*Log_CopyLog) isLog_Log()        {}

func (m *Log) GetLog() isLog_Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *Log) GetListLog() *ListLog {
	if x, ok := m.GetLog().(*Log_ListLog); ok {
		return x.ListLog
	}
	return nil
}

func (m *Log) GetProcessListLog() *ProcessListLog {
	if x, ok := m.GetLog().(*Log_ProcessListLog); ok {
		return x.ProcessListLog
	}
	return nil
}

func (m *Log) GetCopyLog() *CopyLog {
	if x, ok := m.GetLog().(*Log_CopyLog); ok {
		return x.CopyLog
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Log) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Log_OneofMarshaler, _Log_OneofUnmarshaler, _Log_OneofSizer, []interface{}{
		(*Log_ListLog)(nil),
		(*Log_ProcessListLog)(nil),
		(*Log_CopyLog)(nil),
	}
}

func _Log_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Log)
	// log
	switch x := m.Log.(type) {
	case *Log_ListLog:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ListLog); err != nil {
			return err
		}
	case *Log_ProcessListLog:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProcessListLog); err != nil {
			return err
		}
	case *Log_CopyLog:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CopyLog); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Log.Log has unexpected type %T", x)
	}
	return nil
}

func _Log_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Log)
	switch tag {
	case 1: // log.list_log
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ListLog)
		err := b.DecodeMessage(msg)
		m.Log = &Log_ListLog{msg}
		return true, err
	case 2: // log.process_list_log
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProcessListLog)
		err := b.DecodeMessage(msg)
		m.Log = &Log_ProcessListLog{msg}
		return true, err
	case 3: // log.copy_log
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CopyLog)
		err := b.DecodeMessage(msg)
		m.Log = &Log_CopyLog{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Log_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Log)
	// log
	switch x := m.Log.(type) {
	case *Log_ListLog:
		s := proto.Size(x.ListLog)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Log_ProcessListLog:
		s := proto.Size(x.ProcessListLog)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Log_CopyLog:
		s := proto.Size(x.CopyLog)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Contains log fields for a List task.
type ListLog struct {
	FilesFound int64 `protobuf:"varint,1,opt,name=files_found,json=filesFound" json:"files_found,omitempty"`
	BytesFound int64 `protobuf:"varint,2,opt,name=bytes_found,json=bytesFound" json:"bytes_found,omitempty"`
	DirsFound  int64 `protobuf:"varint,3,opt,name=dirs_found,json=dirsFound" json:"dirs_found,omitempty"`
}

func (m *ListLog) Reset()                    { *m = ListLog{} }
func (m *ListLog) String() string            { return proto.CompactTextString(m) }
func (*ListLog) ProtoMessage()               {}
func (*ListLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListLog) GetFilesFound() int64 {
	if m != nil {
		return m.FilesFound
	}
	return 0
}

func (m *ListLog) GetBytesFound() int64 {
	if m != nil {
		return m.BytesFound
	}
	return 0
}

func (m *ListLog) GetDirsFound() int64 {
	if m != nil {
		return m.DirsFound
	}
	return 0
}

// Contains log fields for a ProcessList task.
type ProcessListLog struct {
	EntriesProcessed int64 `protobuf:"varint,1,opt,name=entries_processed,json=entriesProcessed" json:"entries_processed,omitempty"`
	StartingOffset   int64 `protobuf:"varint,2,opt,name=starting_offset,json=startingOffset" json:"starting_offset,omitempty"`
	EndingOffset     int64 `protobuf:"varint,3,opt,name=ending_offset,json=endingOffset" json:"ending_offset,omitempty"`
}

func (m *ProcessListLog) Reset()                    { *m = ProcessListLog{} }
func (m *ProcessListLog) String() string            { return proto.CompactTextString(m) }
func (*ProcessListLog) ProtoMessage()               {}
func (*ProcessListLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ProcessListLog) GetEntriesProcessed() int64 {
	if m != nil {
		return m.EntriesProcessed
	}
	return 0
}

func (m *ProcessListLog) GetStartingOffset() int64 {
	if m != nil {
		return m.StartingOffset
	}
	return 0
}

func (m *ProcessListLog) GetEndingOffset() int64 {
	if m != nil {
		return m.EndingOffset
	}
	return 0
}

// Contains log fields for a Copy task.
type CopyLog struct {
	SrcFile     string `protobuf:"bytes,1,opt,name=src_file,json=srcFile" json:"src_file,omitempty"`
	SrcBytes    int64  `protobuf:"varint,2,opt,name=src_bytes,json=srcBytes" json:"src_bytes,omitempty"`
	SrcMTime    int64  `protobuf:"varint,3,opt,name=src_m_time,json=srcMTime" json:"src_m_time,omitempty"`
	SrcCrc32C   uint32 `protobuf:"varint,4,opt,name=src_crc32c,json=srcCrc32c" json:"src_crc32c,omitempty"`
	DstFile     string `protobuf:"bytes,5,opt,name=dst_file,json=dstFile" json:"dst_file,omitempty"`
	DstBytes    int64  `protobuf:"varint,6,opt,name=dst_bytes,json=dstBytes" json:"dst_bytes,omitempty"`
	DstMTime    int64  `protobuf:"varint,7,opt,name=dst_m_time,json=dstMTime" json:"dst_m_time,omitempty"`
	DstCrc32C   uint32 `protobuf:"varint,8,opt,name=dst_crc32c,json=dstCrc32c" json:"dst_crc32c,omitempty"`
	BytesCopied int64  `protobuf:"varint,9,opt,name=bytes_copied,json=bytesCopied" json:"bytes_copied,omitempty"`
}

func (m *CopyLog) Reset()                    { *m = CopyLog{} }
func (m *CopyLog) String() string            { return proto.CompactTextString(m) }
func (*CopyLog) ProtoMessage()               {}
func (*CopyLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CopyLog) GetSrcFile() string {
	if m != nil {
		return m.SrcFile
	}
	return ""
}

func (m *CopyLog) GetSrcBytes() int64 {
	if m != nil {
		return m.SrcBytes
	}
	return 0
}

func (m *CopyLog) GetSrcMTime() int64 {
	if m != nil {
		return m.SrcMTime
	}
	return 0
}

func (m *CopyLog) GetSrcCrc32C() uint32 {
	if m != nil {
		return m.SrcCrc32C
	}
	return 0
}

func (m *CopyLog) GetDstFile() string {
	if m != nil {
		return m.DstFile
	}
	return ""
}

func (m *CopyLog) GetDstBytes() int64 {
	if m != nil {
		return m.DstBytes
	}
	return 0
}

func (m *CopyLog) GetDstMTime() int64 {
	if m != nil {
		return m.DstMTime
	}
	return 0
}

func (m *CopyLog) GetDstCrc32C() uint32 {
	if m != nil {
		return m.DstCrc32C
	}
	return 0
}

func (m *CopyLog) GetBytesCopied() int64 {
	if m != nil {
		return m.BytesCopied
	}
	return 0
}

func init() {
	proto.RegisterType((*Spec)(nil), "cloud_ingest_task.Spec")
	proto.RegisterType((*ListSpec)(nil), "cloud_ingest_task.ListSpec")
	proto.RegisterType((*ProcessListSpec)(nil), "cloud_ingest_task.ProcessListSpec")
	proto.RegisterType((*CopySpec)(nil), "cloud_ingest_task.CopySpec")
	proto.RegisterType((*TaskReqMsg)(nil), "cloud_ingest_task.TaskReqMsg")
	proto.RegisterType((*TaskRespMsg)(nil), "cloud_ingest_task.TaskRespMsg")
	proto.RegisterType((*Log)(nil), "cloud_ingest_task.Log")
	proto.RegisterType((*ListLog)(nil), "cloud_ingest_task.ListLog")
	proto.RegisterType((*ProcessListLog)(nil), "cloud_ingest_task.ProcessListLog")
	proto.RegisterType((*CopyLog)(nil), "cloud_ingest_task.CopyLog")
	proto.RegisterEnum("cloud_ingest_task.Type", Type_name, Type_value)
	proto.RegisterEnum("cloud_ingest_task.Status", Status_name, Status_value)
	proto.RegisterEnum("cloud_ingest_task.FailureType", FailureType_name, FailureType_value)
}

func init() { proto.RegisterFile("task.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcb, 0x6e, 0xdb, 0x46,
	0x17, 0x36, 0x45, 0x59, 0x97, 0x23, 0x59, 0xa2, 0x27, 0x7f, 0x14, 0xe5, 0x9e, 0xe8, 0x5f, 0xd4,
	0x48, 0x00, 0x2f, 0x9c, 0xa2, 0x05, 0x8a, 0x6e, 0x6c, 0x5d, 0x6a, 0xa1, 0xba, 0x95, 0xa4, 0x50,
	0xa4, 0x9b, 0x01, 0x45, 0x8e, 0x54, 0x26, 0x94, 0xc8, 0x70, 0x46, 0x68, 0xf5, 0x04, 0x5d, 0x74,
	0xdf, 0xc7, 0x29, 0xd0, 0x07, 0xe8, 0xa2, 0x5d, 0xf5, 0x71, 0x8a, 0x33, 0x33, 0x94, 0x64, 0xc7,
	0x36, 0xd0, 0xae, 0xba, 0xd2, 0xf0, 0x7c, 0xe7, 0x7c, 0xe7, 0x9b, 0xcb, 0xf9, 0x20, 0x00, 0xe1,
	0xf1, 0xf7, 0xa7, 0x49, 0x1a, 0x8b, 0x98, 0x1c, 0xfb, 0x51, 0xbc, 0x0e, 0x68, 0xb8, 0x5a, 0x30,
	0x2e, 0x28, 0x02, 0xad, 0xbf, 0x0c, 0xc8, 0x3b, 0x09, 0xf3, 0xc9, 0x17, 0x50, 0x8e, 0x42, 0x2e,
	0x28, 0x4f, 0x98, 0xdf, 0x34, 0x5e, 0x18, 0x27, 0x95, 0xb3, 0xc7, 0xa7, 0x1f, 0xe5, 0x9f, 0x0e,
	0x42, 0x2e, 0x30, 0xff, 0xf2, 0xc0, 0x2e, 0x45, 0x7a, 0x4d, 0x26, 0x70, 0x9c, 0xa4, 0xb1, 0xcf,
	0x38, 0xa7, 0x3b, 0x8e, 0x9c, 0xe4, 0x68, 0xdd, 0xc0, 0x31, 0x51, 0xb9, 0x7b, 0x54, 0xf5, 0xe4,
	0x6a, 0x08, 0xd5, 0xf8, 0x71, 0xb2, 0x51, 0x4c, 0xe6, 0xad, 0x6a, 0xda, 0x71, 0xb2, 0xc9, 0xd4,
	0xf8, 0x7a, 0x7d, 0x51, 0x80, 0x3c, 0x96, 0xb5, 0xfe, 0x30, 0xa0, 0xb4, 0x25, 0x7c, 0x03, 0x8d,
	0x80, 0x0b, 0x25, 0x2f, 0x65, 0x7c, 0x1d, 0x09, 0x3a, 0x5b, 0xfb, 0xef, 0x99, 0x90, 0x7b, 0x2d,
	0xdb, 0xf7, 0x02, 0x2e, 0x30, 0xd9, 0x96, 0xd8, 0x85, 0x84, 0x6e, 0x2a, 0x8a, 0x67, 0xef, 0x98,
	0x2f, 0xe4, 0xe6, 0xae, 0x17, 0x8d, 0x25, 0x44, 0xfe, 0x0f, 0x47, 0x3c, 0xf5, 0x69, 0x10, 0xa6,
	0xcc, 0x17, 0x71, 0xba, 0x91, 0xf2, 0xcb, 0x76, 0x95, 0xa7, 0x7e, 0x27, 0x8b, 0x91, 0xcf, 0xe0,
	0x01, 0xfb, 0x31, 0x61, 0xbe, 0x60, 0x01, 0x5d, 0xb0, 0x15, 0x4b, 0x3d, 0x11, 0xc6, 0x2b, 0xba,
	0x5a, 0x2f, 0x9b, 0xf9, 0x17, 0xc6, 0x89, 0x69, 0xdf, 0xcf, 0xe0, 0xaf, 0xb6, 0xe8, 0x68, 0xbd,
	0x6c, 0xfd, 0x66, 0x40, 0xfd, 0xda, 0xf1, 0xfd, 0xd7, 0xb6, 0xf6, 0x1c, 0x2a, 0xb3, 0x8d, 0x60,
	0x34, 0x9e, 0xcf, 0x39, 0x13, 0x7a, 0x3b, 0x80, 0xa1, 0xb1, 0x8c, 0xb4, 0x7e, 0x32, 0xa1, 0x94,
	0x5d, 0x1c, 0x79, 0x08, 0x25, 0xa4, 0x9c, 0x87, 0x11, 0xd3, 0x72, 0x8b, 0x3c, 0xf5, 0x7b, 0x61,
	0xc4, 0xc8, 0x53, 0x00, 0x94, 0xa8, 0xf7, 0xa2, 0x64, 0x95, 0x03, 0x9e, 0xed, 0x40, 0xc3, 0x5a,
	0xb5, 0xb9, 0x85, 0xb5, 0xd6, 0x7f, 0x79, 0xc2, 0xe4, 0x09, 0x94, 0x67, 0xde, 0x2a, 0xf8, 0x21,
	0x0c, 0xc4, 0xf7, 0xcd, 0x43, 0x99, 0xb9, 0x0b, 0x60, 0x53, 0x94, 0x4a, 0x71, 0x3b, 0xbc, 0x59,
	0x50, 0x30, 0x46, 0x2e, 0x30, 0x40, 0x9e, 0x41, 0x45, 0xc2, 0x4b, 0x2a, 0xc2, 0x25, 0x6b, 0x16,
	0x77, 0xf8, 0xd0, 0x0d, 0x97, 0x8c, 0xbc, 0x84, 0xaa, 0xac, 0xa4, 0x7e, 0x9c, 0x84, 0x2c, 0x68,
	0x96, 0x64, 0x82, 0x3c, 0x2f, 0xde, 0x96, 0x21, 0xd2, 0x80, 0x82, 0x9f, 0xfa, 0x6f, 0xce, 0xfc,
	0x66, 0xf9, 0x85, 0x71, 0x72, 0x64, 0xeb, 0x2f, 0xd2, 0x82, 0x23, 0x55, 0x2a, 0x62, 0xac, 0xde,
	0x34, 0x61, 0xaf, 0xd6, 0x8d, 0xf1, 0x40, 0xc9, 0x29, 0xdc, 0xc3, 0xbb, 0x5c, 0x7a, 0xb3, 0x88,
	0xd1, 0x75, 0x12, 0xc5, 0x5e, 0x40, 0xc3, 0xa0, 0x59, 0x91, 0x67, 0x73, 0xbc, 0x85, 0xa6, 0x12,
	0xe9, 0x07, 0xad, 0x39, 0x80, 0xeb, 0xf1, 0xf7, 0x36, 0xfb, 0x30, 0xe4, 0x0b, 0xf2, 0x1a, 0x08,
	0x0e, 0x15, 0x4d, 0x59, 0x44, 0x53, 0xbc, 0x94, 0x95, 0xb7, 0xcc, 0x2e, 0xa5, 0x2e, 0x64, 0x5e,
	0x64, 0xf3, 0xd4, 0x1f, 0x79, 0x4b, 0x46, 0x5e, 0xab, 0x21, 0xd3, 0x53, 0xfe, 0xe0, 0x86, 0xd9,
	0xc4, 0xeb, 0xb5, 0xd5, 0x24, 0xfe, 0x99, 0x83, 0x8a, 0x6a, 0xc4, 0x93, 0x7f, 0xdc, 0xa9, 0x01,
	0x05, 0x2e, 0x3c, 0xb1, 0xe6, 0xfa, 0x09, 0xe8, 0x2f, 0x72, 0x0e, 0xd5, 0xb9, 0x17, 0x46, 0xeb,
	0x94, 0x51, 0xb1, 0x49, 0x98, 0x7c, 0x01, 0xb5, 0xb3, 0x67, 0x37, 0x28, 0xe9, 0xa9, 0x34, 0x77,
	0x93, 0x30, 0xbb, 0x32, 0xdf, 0x7d, 0x90, 0x4f, 0xa0, 0x9e, 0x51, 0x2c, 0x19, 0xe7, 0xde, 0x82,
	0xc9, 0xb7, 0x51, 0xb6, 0x6b, 0x3a, 0x3c, 0x54, 0x51, 0x72, 0x02, 0x66, 0x14, 0x2f, 0xe4, 0x73,
	0xa8, 0x9c, 0x35, 0x6e, 0xb2, 0xc5, 0x78, 0x61, 0x63, 0x0a, 0x39, 0x83, 0x52, 0xca, 0x3e, 0x28,
	0xdf, 0x2a, 0xdc, 0x7d, 0x36, 0xc5, 0x94, 0x7d, 0x90, 0x33, 0xf0, 0x29, 0x94, 0x53, 0xc6, 0x13,
	0x55, 0x54, 0xbc, 0xbb, 0xa8, 0x84, 0x99, 0xb8, 0x6a, 0xfd, 0x6e, 0x80, 0x39, 0x88, 0x17, 0xe4,
	0x73, 0x90, 0x46, 0x4c, 0x51, 0xa0, 0xf2, 0xed, 0x47, 0xb7, 0xf8, 0xf6, 0x20, 0x5e, 0x5c, 0x1e,
	0xd8, 0xc5, 0x48, 0x2d, 0xc9, 0x10, 0xac, 0x2b, 0xae, 0x8d, 0x04, 0xea, 0x3a, 0x5f, 0xde, 0x6d,
	0xda, 0x8a, 0xa7, 0x96, 0x5c, 0x89, 0xa0, 0x0e, 0x69, 0xd9, 0x48, 0x63, 0xde, 0xaa, 0x03, 0xdf,
	0xa9, 0xd6, 0xe1, 0xab, 0xe5, 0xc5, 0xa1, 0x3c, 0xdc, 0xd6, 0x3b, 0x28, 0x66, 0x54, 0xcf, 0xd5,
	0x18, 0x71, 0x3a, 0x8f, 0xd7, 0xab, 0x40, 0xee, 0xca, 0xb4, 0xe5, 0xe0, 0xf1, 0x1e, 0x46, 0x32,
	0x8f, 0xc9, 0x12, 0x72, 0x3b, 0x8f, 0xd1, 0x09, 0x68, 0x0e, 0x61, 0x9a, 0xe1, 0xa6, 0x9a, 0x43,
	0x8c, 0x48, 0xb8, 0xf5, 0xb3, 0x01, 0xb5, 0xab, 0x1b, 0x22, 0xaf, 0xe1, 0x98, 0xad, 0x44, 0x1a,
	0x32, 0x4e, 0xf5, 0xc6, 0x58, 0xd6, 0xd9, 0xd2, 0xc0, 0x24, 0x8b, 0xe3, 0xc3, 0xe1, 0xc2, 0x4b,
	0x45, 0xb8, 0x5a, 0x64, 0x3e, 0xa7, 0x34, 0xd4, 0xb2, 0xb0, 0xf2, 0x3a, 0x74, 0x4c, 0xb6, 0x0a,
	0xf6, 0xd2, 0x94, 0x94, 0xaa, 0x0a, 0x6a, 0x43, 0xfc, 0x25, 0x07, 0x45, 0x7d, 0x2e, 0x77, 0xf9,
	0xe1, 0x63, 0x28, 0x23, 0xa4, 0xac, 0x47, 0xb5, 0xc3, 0x5c, 0xe5, 0x3c, 0x4f, 0x00, 0x10, 0xd4,
	0xc6, 0x63, 0x6e, 0x51, 0xe5, 0x3b, 0x4f, 0x15, 0xaa, 0x8d, 0x25, 0x2f, 0x8d, 0x05, 0xc9, 0xda,
	0xca, 0x5b, 0x1e, 0x42, 0x09, 0xad, 0x54, 0x36, 0x3d, 0x54, 0x4d, 0x03, 0x2e, 0xb2, 0xa6, 0xd2,
	0x84, 0xf7, 0xfc, 0x0e, 0x73, 0xb7, 0x4d, 0x11, 0xbc, 0xe2, 0x76, 0x88, 0x6e, 0x9b, 0x22, 0xaa,
	0x9b, 0x96, 0x54, 0xd3, 0x80, 0x0b, 0xdd, 0xf4, 0xba, 0x17, 0x96, 0x3f, 0xf2, 0xc2, 0x57, 0x5f,
	0x42, 0x5e, 0xce, 0x69, 0x0d, 0x60, 0x3a, 0x72, 0xba, 0x2e, 0x75, 0xdf, 0x4e, 0xba, 0xd6, 0x01,
	0x29, 0x41, 0x7e, 0xd0, 0x77, 0x5c, 0xcb, 0x20, 0x16, 0x54, 0x27, 0xf6, 0xb8, 0xdd, 0x75, 0x1c,
	0x2a, 0x23, 0x39, 0xc4, 0xda, 0xe3, 0xc9, 0x5b, 0xcb, 0x7c, 0x35, 0x84, 0x82, 0xa3, 0xac, 0xc2,
	0x82, 0xaa, 0xaa, 0x77, 0xdc, 0x73, 0x77, 0xea, 0x58, 0x07, 0xa4, 0x0a, 0xa5, 0xe9, 0xe8, 0x9b,
	0x69, 0x77, 0xda, 0xed, 0x58, 0x06, 0x01, 0x28, 0xe8, 0x75, 0x0e, 0xd7, 0xbd, 0xf3, 0xfe, 0xa0,
	0xdb, 0xb1, 0x4c, 0x52, 0x81, 0xa2, 0x33, 0x6d, 0x23, 0xbb, 0x95, 0x7f, 0xf5, 0xab, 0x01, 0x95,
	0x3d, 0x27, 0x21, 0x0d, 0x20, 0x8a, 0x14, 0xd3, 0xa7, 0x76, 0x37, 0x13, 0x77, 0x0f, 0xea, 0xd3,
	0xd1, 0xd7, 0xa3, 0xf1, 0xb7, 0xa3, 0x0c, 0xb1, 0x0c, 0xf2, 0x10, 0xee, 0xf7, 0xfa, 0x83, 0x2e,
	0x1d, 0x8e, 0x3b, 0xfd, 0x5e, 0xbf, 0xdb, 0xd9, 0x42, 0x39, 0x84, 0x2e, 0xcf, 0x9d, 0x4b, 0x3a,
	0xec, 0x3b, 0xc3, 0x73, 0xb7, 0x7d, 0xb9, 0x85, 0x4c, 0xd2, 0x84, 0xff, 0x4d, 0xec, 0x6e, 0x7b,
	0x3c, 0xea, 0xf4, 0xdd, 0xfe, 0x78, 0xc7, 0x97, 0x27, 0x8f, 0xa0, 0x21, 0xf9, 0x46, 0x63, 0x97,
	0xf6, 0xc6, 0xd3, 0xd1, 0x8e, 0xf0, 0x10, 0x85, 0x4d, 0xba, 0xf6, 0xb0, 0xef, 0x38, 0xfb, 0x35,
	0x85, 0x8b, 0xfa, 0x77, 0x47, 0xd2, 0x75, 0x17, 0x31, 0x95, 0x7f, 0x07, 0x67, 0x05, 0xf9, 0xf3,
	0xe6, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x7f, 0xcb, 0xf7, 0x23, 0x0a, 0x00, 0x00,
}