// Code generated by protoc-gen-go.
// source: analyzer.proto
// DO NOT EDIT!

/*
Package pulumirpc is a generated protocol buffer package.

It is generated from these files:
	analyzer.proto
	engine.proto
	errors.proto
	language.proto
	plugin.proto
	provider.proto
	resource.proto

It has these top-level messages:
	AnalyzeRequest
	AnalyzeResponse
	AnalyzeFailure
	LogRequest
	ErrorCause
	GetRequiredPluginsRequest
	GetRequiredPluginsResponse
	RunRequest
	RunResponse
	PluginInfo
	PluginDependency
	ConfigureRequest
	ConfigureErrorMissingKeys
	InvokeRequest
	InvokeResponse
	CheckRequest
	CheckResponse
	CheckFailure
	DiffRequest
	DiffResponse
	CreateRequest
	CreateResponse
	ReadRequest
	ReadResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	RegisterResourceRequest
	RegisterResourceResponse
	RegisterResourceOutputsRequest
*/
package pulumirpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import google_protobuf1 "github.com/golang/protobuf/ptypes/struct"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AnalyzeRequest struct {
	Type       string                   `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Properties *google_protobuf1.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
}

func (m *AnalyzeRequest) Reset()                    { *m = AnalyzeRequest{} }
func (m *AnalyzeRequest) String() string            { return proto.CompactTextString(m) }
func (*AnalyzeRequest) ProtoMessage()               {}
func (*AnalyzeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AnalyzeRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AnalyzeRequest) GetProperties() *google_protobuf1.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type AnalyzeResponse struct {
	Failures []*AnalyzeFailure `protobuf:"bytes,1,rep,name=failures" json:"failures,omitempty"`
}

func (m *AnalyzeResponse) Reset()                    { *m = AnalyzeResponse{} }
func (m *AnalyzeResponse) String() string            { return proto.CompactTextString(m) }
func (*AnalyzeResponse) ProtoMessage()               {}
func (*AnalyzeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AnalyzeResponse) GetFailures() []*AnalyzeFailure {
	if m != nil {
		return m.Failures
	}
	return nil
}

type AnalyzeFailure struct {
	Property string `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	Reason   string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
}

func (m *AnalyzeFailure) Reset()                    { *m = AnalyzeFailure{} }
func (m *AnalyzeFailure) String() string            { return proto.CompactTextString(m) }
func (*AnalyzeFailure) ProtoMessage()               {}
func (*AnalyzeFailure) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AnalyzeFailure) GetProperty() string {
	if m != nil {
		return m.Property
	}
	return ""
}

func (m *AnalyzeFailure) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*AnalyzeRequest)(nil), "pulumirpc.AnalyzeRequest")
	proto.RegisterType((*AnalyzeResponse)(nil), "pulumirpc.AnalyzeResponse")
	proto.RegisterType((*AnalyzeFailure)(nil), "pulumirpc.AnalyzeFailure")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Analyzer service

type AnalyzerClient interface {
	// Analyze analyzes a single resource object, and returns any errors that it finds.
	Analyze(ctx context.Context, in *AnalyzeRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error)
	// GetPluginInfo returns generic information about this plugin, like its version.
	GetPluginInfo(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PluginInfo, error)
}

type analyzerClient struct {
	cc *grpc.ClientConn
}

func NewAnalyzerClient(cc *grpc.ClientConn) AnalyzerClient {
	return &analyzerClient{cc}
}

func (c *analyzerClient) Analyze(ctx context.Context, in *AnalyzeRequest, opts ...grpc.CallOption) (*AnalyzeResponse, error) {
	out := new(AnalyzeResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.Analyzer/Analyze", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) GetPluginInfo(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PluginInfo, error) {
	out := new(PluginInfo)
	err := grpc.Invoke(ctx, "/pulumirpc.Analyzer/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Analyzer service

type AnalyzerServer interface {
	// Analyze analyzes a single resource object, and returns any errors that it finds.
	Analyze(context.Context, *AnalyzeRequest) (*AnalyzeResponse, error)
	// GetPluginInfo returns generic information about this plugin, like its version.
	GetPluginInfo(context.Context, *google_protobuf.Empty) (*PluginInfo, error)
}

func RegisterAnalyzerServer(s *grpc.Server, srv AnalyzerServer) {
	s.RegisterService(&_Analyzer_serviceDesc, srv)
}

func _Analyzer_Analyze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).Analyze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.Analyzer/Analyze",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).Analyze(ctx, req.(*AnalyzeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.Analyzer/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).GetPluginInfo(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Analyzer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pulumirpc.Analyzer",
	HandlerType: (*AnalyzerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Analyze",
			Handler:    _Analyzer_Analyze_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _Analyzer_GetPluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analyzer.proto",
}

func init() { proto.RegisterFile("analyzer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x1b, 0x95, 0xda, 0x4e, 0xb5, 0xc2, 0x80, 0xb5, 0xae, 0x1e, 0x42, 0x4e, 0x39, 0x6d,
	0x21, 0x22, 0x5e, 0x55, 0xfc, 0x7b, 0x93, 0x78, 0xf6, 0x90, 0x96, 0x49, 0x08, 0xa4, 0xd9, 0x75,
	0xff, 0x1c, 0xe2, 0xa7, 0xf0, 0x23, 0x8b, 0xbb, 0x6b, 0x2c, 0xda, 0xdb, 0x0c, 0xef, 0xf1, 0x9b,
	0x37, 0x33, 0x30, 0x2d, 0xda, 0xa2, 0xe9, 0x3e, 0x48, 0x71, 0xa9, 0x84, 0x11, 0x38, 0x96, 0xb6,
	0xb1, 0xeb, 0x5a, 0xc9, 0x15, 0x3b, 0x90, 0x8d, 0xad, 0xea, 0xd6, 0x0b, 0xec, 0xac, 0x12, 0xa2,
	0x6a, 0x68, 0xe1, 0xba, 0xa5, 0x2d, 0x17, 0xb4, 0x96, 0xa6, 0x0b, 0xe2, 0xf9, 0x5f, 0x51, 0x1b,
	0x65, 0x57, 0xc6, 0xab, 0xc9, 0x1b, 0x4c, 0x6f, 0xfc, 0x94, 0x9c, 0xde, 0x2d, 0x69, 0x83, 0x08,
	0x7b, 0xa6, 0x93, 0x34, 0x8f, 0xe2, 0x28, 0x1d, 0xe7, 0xae, 0xc6, 0x2b, 0x00, 0xa9, 0x84, 0x24,
	0x65, 0x6a, 0xd2, 0xf3, 0x9d, 0x38, 0x4a, 0x27, 0xd9, 0x09, 0xf7, 0x60, 0xfe, 0x03, 0xe6, 0xaf,
	0x0e, 0x9c, 0x6f, 0x58, 0x93, 0x27, 0x38, 0xea, 0xf1, 0x5a, 0x8a, 0x56, 0x13, 0x5e, 0xc2, 0xa8,
	0x2c, 0xea, 0xc6, 0x2a, 0xd2, 0xf3, 0x28, 0xde, 0x4d, 0x27, 0xd9, 0x29, 0xef, 0x17, 0xe3, 0xc1,
	0xfd, 0xe0, 0x1d, 0x79, 0x6f, 0x4d, 0xee, 0xfa, 0xa0, 0x41, 0x43, 0x06, 0xa3, 0x30, 0xa9, 0x0b,
	0x61, 0xfb, 0x1e, 0x67, 0x30, 0x54, 0x54, 0x68, 0xd1, 0xba, 0xb0, 0xe3, 0x3c, 0x74, 0xd9, 0x67,
	0x04, 0xa3, 0x80, 0x51, 0x78, 0x0b, 0xfb, 0xa1, 0xc6, 0x2d, 0x11, 0xc2, 0x3d, 0x18, 0xdb, 0x26,
	0xf9, 0x5d, 0x92, 0x01, 0x5e, 0xc3, 0xe1, 0x23, 0x99, 0x17, 0xf7, 0x8d, 0xe7, 0xb6, 0x14, 0x38,
	0xfb, 0x77, 0x96, 0xfb, 0xef, 0x67, 0xb0, 0xe3, 0x0d, 0xcc, 0xaf, 0x3d, 0x19, 0x2c, 0x87, 0xce,
	0x78, 0xf1, 0x15, 0x00, 0x00, 0xff, 0xff, 0xff, 0x6d, 0x9c, 0x46, 0xee, 0x01, 0x00, 0x00,
}
