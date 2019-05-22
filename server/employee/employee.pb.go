// Code generated by protoc-gen-go. DO NOT EDIT.
// source: employee.proto

package employee

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Nothing struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nothing) Reset()         { *m = Nothing{} }
func (m *Nothing) String() string { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()    {}
func (*Nothing) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb50a19aa79a6eac, []int{0}
}

func (m *Nothing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nothing.Unmarshal(m, b)
}
func (m *Nothing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nothing.Marshal(b, m, deterministic)
}
func (m *Nothing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nothing.Merge(m, src)
}
func (m *Nothing) XXX_Size() int {
	return xxx_messageInfo_Nothing.Size(m)
}
func (m *Nothing) XXX_DiscardUnknown() {
	xxx_messageInfo_Nothing.DiscardUnknown(m)
}

var xxx_messageInfo_Nothing proto.InternalMessageInfo

type Employee struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Designation          string   `protobuf:"bytes,3,opt,name=designation,proto3" json:"designation,omitempty"`
	Salary               float32  `protobuf:"fixed32,4,opt,name=salary,proto3" json:"salary,omitempty"`
	Department           *Dept    `protobuf:"bytes,5,opt,name=department,proto3" json:"department,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Employee) Reset()         { *m = Employee{} }
func (m *Employee) String() string { return proto.CompactTextString(m) }
func (*Employee) ProtoMessage()    {}
func (*Employee) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb50a19aa79a6eac, []int{1}
}

func (m *Employee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Employee.Unmarshal(m, b)
}
func (m *Employee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Employee.Marshal(b, m, deterministic)
}
func (m *Employee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Employee.Merge(m, src)
}
func (m *Employee) XXX_Size() int {
	return xxx_messageInfo_Employee.Size(m)
}
func (m *Employee) XXX_DiscardUnknown() {
	xxx_messageInfo_Employee.DiscardUnknown(m)
}

var xxx_messageInfo_Employee proto.InternalMessageInfo

func (m *Employee) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Employee) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Employee) GetDesignation() string {
	if m != nil {
		return m.Designation
	}
	return ""
}

func (m *Employee) GetSalary() float32 {
	if m != nil {
		return m.Salary
	}
	return 0
}

func (m *Employee) GetDepartment() *Dept {
	if m != nil {
		return m.Department
	}
	return nil
}

type EmployeeRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmployeeRequest) Reset()         { *m = EmployeeRequest{} }
func (m *EmployeeRequest) String() string { return proto.CompactTextString(m) }
func (*EmployeeRequest) ProtoMessage()    {}
func (*EmployeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb50a19aa79a6eac, []int{2}
}

func (m *EmployeeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmployeeRequest.Unmarshal(m, b)
}
func (m *EmployeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmployeeRequest.Marshal(b, m, deterministic)
}
func (m *EmployeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmployeeRequest.Merge(m, src)
}
func (m *EmployeeRequest) XXX_Size() int {
	return xxx_messageInfo_EmployeeRequest.Size(m)
}
func (m *EmployeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmployeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmployeeRequest proto.InternalMessageInfo

func (m *EmployeeRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type EmployeeUpdateRequest struct {
	Id                   int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Emp                  *Employee `protobuf:"bytes,2,opt,name=emp,proto3" json:"emp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EmployeeUpdateRequest) Reset()         { *m = EmployeeUpdateRequest{} }
func (m *EmployeeUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*EmployeeUpdateRequest) ProtoMessage()    {}
func (*EmployeeUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb50a19aa79a6eac, []int{3}
}

func (m *EmployeeUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmployeeUpdateRequest.Unmarshal(m, b)
}
func (m *EmployeeUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmployeeUpdateRequest.Marshal(b, m, deterministic)
}
func (m *EmployeeUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmployeeUpdateRequest.Merge(m, src)
}
func (m *EmployeeUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_EmployeeUpdateRequest.Size(m)
}
func (m *EmployeeUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmployeeUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmployeeUpdateRequest proto.InternalMessageInfo

func (m *EmployeeUpdateRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EmployeeUpdateRequest) GetEmp() *Employee {
	if m != nil {
		return m.Emp
	}
	return nil
}

type Dept struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dept) Reset()         { *m = Dept{} }
func (m *Dept) String() string { return proto.CompactTextString(m) }
func (*Dept) ProtoMessage()    {}
func (*Dept) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb50a19aa79a6eac, []int{4}
}

func (m *Dept) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dept.Unmarshal(m, b)
}
func (m *Dept) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dept.Marshal(b, m, deterministic)
}
func (m *Dept) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dept.Merge(m, src)
}
func (m *Dept) XXX_Size() int {
	return xxx_messageInfo_Dept.Size(m)
}
func (m *Dept) XXX_DiscardUnknown() {
	xxx_messageInfo_Dept.DiscardUnknown(m)
}

var xxx_messageInfo_Dept proto.InternalMessageInfo

func (m *Dept) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Dept) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeptRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeptRequest) Reset()         { *m = DeptRequest{} }
func (m *DeptRequest) String() string { return proto.CompactTextString(m) }
func (*DeptRequest) ProtoMessage()    {}
func (*DeptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb50a19aa79a6eac, []int{5}
}

func (m *DeptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeptRequest.Unmarshal(m, b)
}
func (m *DeptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeptRequest.Marshal(b, m, deterministic)
}
func (m *DeptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeptRequest.Merge(m, src)
}
func (m *DeptRequest) XXX_Size() int {
	return xxx_messageInfo_DeptRequest.Size(m)
}
func (m *DeptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeptRequest proto.InternalMessageInfo

func (m *DeptRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeptUpdateRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Department           *Dept    `protobuf:"bytes,2,opt,name=department,proto3" json:"department,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeptUpdateRequest) Reset()         { *m = DeptUpdateRequest{} }
func (m *DeptUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*DeptUpdateRequest) ProtoMessage()    {}
func (*DeptUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb50a19aa79a6eac, []int{6}
}

func (m *DeptUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeptUpdateRequest.Unmarshal(m, b)
}
func (m *DeptUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeptUpdateRequest.Marshal(b, m, deterministic)
}
func (m *DeptUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeptUpdateRequest.Merge(m, src)
}
func (m *DeptUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_DeptUpdateRequest.Size(m)
}
func (m *DeptUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeptUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeptUpdateRequest proto.InternalMessageInfo

func (m *DeptUpdateRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeptUpdateRequest) GetDepartment() *Dept {
	if m != nil {
		return m.Department
	}
	return nil
}

type ResponseMessage struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseMessage) Reset()         { *m = ResponseMessage{} }
func (m *ResponseMessage) String() string { return proto.CompactTextString(m) }
func (*ResponseMessage) ProtoMessage()    {}
func (*ResponseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb50a19aa79a6eac, []int{7}
}

func (m *ResponseMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseMessage.Unmarshal(m, b)
}
func (m *ResponseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseMessage.Marshal(b, m, deterministic)
}
func (m *ResponseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMessage.Merge(m, src)
}
func (m *ResponseMessage) XXX_Size() int {
	return xxx_messageInfo_ResponseMessage.Size(m)
}
func (m *ResponseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMessage proto.InternalMessageInfo

func (m *ResponseMessage) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Nothing)(nil), "Nothing")
	proto.RegisterType((*Employee)(nil), "Employee")
	proto.RegisterType((*EmployeeRequest)(nil), "EmployeeRequest")
	proto.RegisterType((*EmployeeUpdateRequest)(nil), "EmployeeUpdateRequest")
	proto.RegisterType((*Dept)(nil), "Dept")
	proto.RegisterType((*DeptRequest)(nil), "DeptRequest")
	proto.RegisterType((*DeptUpdateRequest)(nil), "DeptUpdateRequest")
	proto.RegisterType((*ResponseMessage)(nil), "ResponseMessage")
}

func init() { proto.RegisterFile("employee.proto", fileDescriptor_eb50a19aa79a6eac) }

var fileDescriptor_eb50a19aa79a6eac = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdb, 0xca, 0xd3, 0x40,
	0x10, 0xee, 0xa6, 0xed, 0xdf, 0x3f, 0x13, 0x49, 0xe3, 0x82, 0x25, 0x54, 0x84, 0xb8, 0x2a, 0x84,
	0x22, 0x4b, 0x89, 0xe2, 0x85, 0x78, 0x19, 0x29, 0x88, 0x7a, 0x91, 0xe2, 0x03, 0x44, 0x33, 0xc4,
	0x40, 0x73, 0x30, 0xbb, 0x0a, 0x7d, 0x8a, 0x3e, 0xa8, 0x2f, 0x21, 0xd9, 0x1c, 0x9a, 0xc6, 0xb4,
	0xfc, 0x77, 0xd9, 0x99, 0xf9, 0x66, 0xbf, 0xc3, 0x06, 0x4c, 0x4c, 0x8b, 0x43, 0x7e, 0x44, 0xe4,
	0x45, 0x99, 0xcb, 0x9c, 0xe9, 0xb0, 0xf8, 0x9a, 0xcb, 0x9f, 0x49, 0x16, 0xb3, 0x13, 0x81, 0xfb,
	0x8f, 0x4d, 0x97, 0x9a, 0xa0, 0x25, 0x91, 0x4d, 0x1c, 0xe2, 0xce, 0x03, 0x2d, 0x89, 0x28, 0x85,
	0x59, 0x16, 0xa6, 0x68, 0x6b, 0x0e, 0x71, 0xf5, 0x40, 0x7d, 0x53, 0x07, 0x8c, 0x08, 0x45, 0x12,
	0x67, 0xa1, 0x4c, 0xf2, 0xcc, 0x9e, 0xaa, 0x56, 0xbf, 0x44, 0x57, 0x70, 0x27, 0xc2, 0x43, 0x58,
	0x1e, 0xed, 0x99, 0x43, 0x5c, 0x2d, 0x68, 0x4e, 0xf4, 0x15, 0x40, 0x84, 0x45, 0x58, 0xca, 0x14,
	0x33, 0x69, 0xcf, 0x1d, 0xe2, 0x1a, 0xde, 0x9c, 0xfb, 0x58, 0xc8, 0xa0, 0xd7, 0x60, 0xcf, 0x61,
	0xd9, 0x12, 0x0a, 0xf0, 0xd7, 0x6f, 0x14, 0x72, 0xc8, 0x8b, 0xf9, 0xf0, 0xa4, 0x1d, 0xf9, 0x56,
	0x44, 0xa1, 0xbc, 0x36, 0x48, 0x9f, 0xc2, 0x14, 0xd3, 0x42, 0xf1, 0x37, 0x3c, 0x9d, 0x77, 0x7b,
	0xab, 0x2a, 0xdb, 0xc0, 0xac, 0xba, 0xfc, 0x21, 0xaa, 0xd9, 0x33, 0x30, 0x14, 0xd1, 0x2b, 0x84,
	0x3e, 0xc1, 0xe3, 0xaa, 0x7d, 0x9b, 0xcc, 0xa5, 0x7e, 0xed, 0x9a, 0xfe, 0x17, 0xb0, 0x0c, 0x50,
	0x14, 0x79, 0x26, 0xf0, 0x0b, 0x0a, 0x11, 0xc6, 0x48, 0x2d, 0x98, 0xa6, 0x22, 0x56, 0xab, 0xf4,
	0xa0, 0xfa, 0xf4, 0x4e, 0xda, 0xd9, 0xa5, 0x3d, 0x96, 0x7f, 0x92, 0x1f, 0x48, 0x5f, 0x83, 0xb1,
	0x43, 0xd9, 0x85, 0x69, 0xf1, 0x81, 0x8d, 0xeb, 0xb3, 0x01, 0x6c, 0x42, 0x37, 0xb0, 0xec, 0x4d,
	0x7f, 0x4e, 0x84, 0xa4, 0xf7, 0xbc, 0x79, 0x15, 0x17, 0x93, 0x5b, 0x52, 0x6d, 0xde, 0xf7, 0x36,
	0x9f, 0xbb, 0x6b, 0x8b, 0x0f, 0xb8, 0xb2, 0x09, 0x7d, 0x07, 0xa6, 0x8f, 0x07, 0x94, 0x78, 0x83,
	0xca, 0x18, 0xee, 0x03, 0x98, 0xb5, 0x81, 0x1d, 0x6e, 0xc5, 0x47, 0x63, 0x1e, 0x43, 0x7b, 0x7f,
	0x49, 0x1d, 0x51, 0xeb, 0x86, 0x03, 0x8b, 0x1d, 0x4a, 0x15, 0xf0, 0x23, 0xde, 0xcb, 0x6e, 0x5d,
	0x5b, 0xce, 0x26, 0x94, 0x29, 0xbf, 0xaa, 0xc3, 0x40, 0x7d, 0x3b, 0xb1, 0x25, 0xf4, 0x25, 0x2c,
	0xf6, 0xcd, 0x96, 0xba, 0x3a, 0xca, 0xfc, 0x2d, 0x58, 0xb5, 0x62, 0xbf, 0x8b, 0x71, 0x70, 0xe9,
	0x18, 0xea, 0x3d, 0x58, 0xb5, 0xac, 0x1e, 0x8a, 0xf2, 0xff, 0xde, 0xd1, 0x18, 0xf6, 0xfb, 0x9d,
	0xfa, 0x91, 0xdf, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x74, 0x38, 0x5c, 0x48, 0xda, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmployeeServiceClient is the client API for EmployeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmployeeServiceClient interface {
	GetEmployee(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*Employee, error)
	GetEmployeeList(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (EmployeeService_GetEmployeeListClient, error)
	SetEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*ResponseMessage, error)
	DeleteEmployee(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*ResponseMessage, error)
	UpdateEmployee(ctx context.Context, in *EmployeeUpdateRequest, opts ...grpc.CallOption) (*ResponseMessage, error)
}

type employeeServiceClient struct {
	cc *grpc.ClientConn
}

func NewEmployeeServiceClient(cc *grpc.ClientConn) EmployeeServiceClient {
	return &employeeServiceClient{cc}
}

func (c *employeeServiceClient) GetEmployee(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*Employee, error) {
	out := new(Employee)
	err := c.cc.Invoke(ctx, "/EmployeeService/GetEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) GetEmployeeList(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (EmployeeService_GetEmployeeListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EmployeeService_serviceDesc.Streams[0], "/EmployeeService/GetEmployeeList", opts...)
	if err != nil {
		return nil, err
	}
	x := &employeeServiceGetEmployeeListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EmployeeService_GetEmployeeListClient interface {
	Recv() (*Employee, error)
	grpc.ClientStream
}

type employeeServiceGetEmployeeListClient struct {
	grpc.ClientStream
}

func (x *employeeServiceGetEmployeeListClient) Recv() (*Employee, error) {
	m := new(Employee)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *employeeServiceClient) SetEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/EmployeeService/SetEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) DeleteEmployee(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/EmployeeService/DeleteEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) UpdateEmployee(ctx context.Context, in *EmployeeUpdateRequest, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/EmployeeService/UpdateEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeServiceServer is the server API for EmployeeService service.
type EmployeeServiceServer interface {
	GetEmployee(context.Context, *EmployeeRequest) (*Employee, error)
	GetEmployeeList(*Nothing, EmployeeService_GetEmployeeListServer) error
	SetEmployee(context.Context, *Employee) (*ResponseMessage, error)
	DeleteEmployee(context.Context, *EmployeeRequest) (*ResponseMessage, error)
	UpdateEmployee(context.Context, *EmployeeUpdateRequest) (*ResponseMessage, error)
}

// UnimplementedEmployeeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEmployeeServiceServer struct {
}

func (*UnimplementedEmployeeServiceServer) GetEmployee(ctx context.Context, req *EmployeeRequest) (*Employee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployee not implemented")
}
func (*UnimplementedEmployeeServiceServer) GetEmployeeList(req *Nothing, srv EmployeeService_GetEmployeeListServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEmployeeList not implemented")
}
func (*UnimplementedEmployeeServiceServer) SetEmployee(ctx context.Context, req *Employee) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetEmployee not implemented")
}
func (*UnimplementedEmployeeServiceServer) DeleteEmployee(ctx context.Context, req *EmployeeRequest) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmployee not implemented")
}
func (*UnimplementedEmployeeServiceServer) UpdateEmployee(ctx context.Context, req *EmployeeUpdateRequest) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployee not implemented")
}

func RegisterEmployeeServiceServer(s *grpc.Server, srv EmployeeServiceServer) {
	s.RegisterService(&_EmployeeService_serviceDesc, srv)
}

func _EmployeeService_GetEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).GetEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmployeeService/GetEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).GetEmployee(ctx, req.(*EmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_GetEmployeeList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Nothing)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EmployeeServiceServer).GetEmployeeList(m, &employeeServiceGetEmployeeListServer{stream})
}

type EmployeeService_GetEmployeeListServer interface {
	Send(*Employee) error
	grpc.ServerStream
}

type employeeServiceGetEmployeeListServer struct {
	grpc.ServerStream
}

func (x *employeeServiceGetEmployeeListServer) Send(m *Employee) error {
	return x.ServerStream.SendMsg(m)
}

func _EmployeeService_SetEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).SetEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmployeeService/SetEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).SetEmployee(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_DeleteEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).DeleteEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmployeeService/DeleteEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).DeleteEmployee(ctx, req.(*EmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_UpdateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).UpdateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmployeeService/UpdateEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).UpdateEmployee(ctx, req.(*EmployeeUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmployeeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "EmployeeService",
	HandlerType: (*EmployeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEmployee",
			Handler:    _EmployeeService_GetEmployee_Handler,
		},
		{
			MethodName: "SetEmployee",
			Handler:    _EmployeeService_SetEmployee_Handler,
		},
		{
			MethodName: "DeleteEmployee",
			Handler:    _EmployeeService_DeleteEmployee_Handler,
		},
		{
			MethodName: "UpdateEmployee",
			Handler:    _EmployeeService_UpdateEmployee_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEmployeeList",
			Handler:       _EmployeeService_GetEmployeeList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "employee.proto",
}

// DeptServiceClient is the client API for DeptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeptServiceClient interface {
	GetDept(ctx context.Context, in *DeptRequest, opts ...grpc.CallOption) (*Dept, error)
	GetDeptList(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (DeptService_GetDeptListClient, error)
	SetDept(ctx context.Context, in *Dept, opts ...grpc.CallOption) (*ResponseMessage, error)
	DeleteDepartment(ctx context.Context, in *DeptRequest, opts ...grpc.CallOption) (*ResponseMessage, error)
	UpdateDepartment(ctx context.Context, in *DeptUpdateRequest, opts ...grpc.CallOption) (*ResponseMessage, error)
}

type deptServiceClient struct {
	cc *grpc.ClientConn
}

func NewDeptServiceClient(cc *grpc.ClientConn) DeptServiceClient {
	return &deptServiceClient{cc}
}

func (c *deptServiceClient) GetDept(ctx context.Context, in *DeptRequest, opts ...grpc.CallOption) (*Dept, error) {
	out := new(Dept)
	err := c.cc.Invoke(ctx, "/DeptService/GetDept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deptServiceClient) GetDeptList(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (DeptService_GetDeptListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DeptService_serviceDesc.Streams[0], "/DeptService/GetDeptList", opts...)
	if err != nil {
		return nil, err
	}
	x := &deptServiceGetDeptListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DeptService_GetDeptListClient interface {
	Recv() (*Dept, error)
	grpc.ClientStream
}

type deptServiceGetDeptListClient struct {
	grpc.ClientStream
}

func (x *deptServiceGetDeptListClient) Recv() (*Dept, error) {
	m := new(Dept)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *deptServiceClient) SetDept(ctx context.Context, in *Dept, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/DeptService/SetDept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deptServiceClient) DeleteDepartment(ctx context.Context, in *DeptRequest, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/DeptService/DeleteDepartment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deptServiceClient) UpdateDepartment(ctx context.Context, in *DeptUpdateRequest, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/DeptService/UpdateDepartment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeptServiceServer is the server API for DeptService service.
type DeptServiceServer interface {
	GetDept(context.Context, *DeptRequest) (*Dept, error)
	GetDeptList(*Nothing, DeptService_GetDeptListServer) error
	SetDept(context.Context, *Dept) (*ResponseMessage, error)
	DeleteDepartment(context.Context, *DeptRequest) (*ResponseMessage, error)
	UpdateDepartment(context.Context, *DeptUpdateRequest) (*ResponseMessage, error)
}

// UnimplementedDeptServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDeptServiceServer struct {
}

func (*UnimplementedDeptServiceServer) GetDept(ctx context.Context, req *DeptRequest) (*Dept, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDept not implemented")
}
func (*UnimplementedDeptServiceServer) GetDeptList(req *Nothing, srv DeptService_GetDeptListServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDeptList not implemented")
}
func (*UnimplementedDeptServiceServer) SetDept(ctx context.Context, req *Dept) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDept not implemented")
}
func (*UnimplementedDeptServiceServer) DeleteDepartment(ctx context.Context, req *DeptRequest) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDepartment not implemented")
}
func (*UnimplementedDeptServiceServer) UpdateDepartment(ctx context.Context, req *DeptUpdateRequest) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDepartment not implemented")
}

func RegisterDeptServiceServer(s *grpc.Server, srv DeptServiceServer) {
	s.RegisterService(&_DeptService_serviceDesc, srv)
}

func _DeptService_GetDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeptServiceServer).GetDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DeptService/GetDept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeptServiceServer).GetDept(ctx, req.(*DeptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeptService_GetDeptList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Nothing)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DeptServiceServer).GetDeptList(m, &deptServiceGetDeptListServer{stream})
}

type DeptService_GetDeptListServer interface {
	Send(*Dept) error
	grpc.ServerStream
}

type deptServiceGetDeptListServer struct {
	grpc.ServerStream
}

func (x *deptServiceGetDeptListServer) Send(m *Dept) error {
	return x.ServerStream.SendMsg(m)
}

func _DeptService_SetDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Dept)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeptServiceServer).SetDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DeptService/SetDept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeptServiceServer).SetDept(ctx, req.(*Dept))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeptService_DeleteDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeptServiceServer).DeleteDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DeptService/DeleteDepartment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeptServiceServer).DeleteDepartment(ctx, req.(*DeptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeptService_UpdateDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeptServiceServer).UpdateDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DeptService/UpdateDepartment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeptServiceServer).UpdateDepartment(ctx, req.(*DeptUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeptService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DeptService",
	HandlerType: (*DeptServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDept",
			Handler:    _DeptService_GetDept_Handler,
		},
		{
			MethodName: "SetDept",
			Handler:    _DeptService_SetDept_Handler,
		},
		{
			MethodName: "DeleteDepartment",
			Handler:    _DeptService_DeleteDepartment_Handler,
		},
		{
			MethodName: "UpdateDepartment",
			Handler:    _DeptService_UpdateDepartment_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDeptList",
			Handler:       _DeptService_GetDeptList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "employee.proto",
}
