// Code generated by protoc-gen-gogo.
// source: translator.proto
// DO NOT EDIT!

/*
Package translatorpb is a generated protocol buffer package.

It is generated from these files:
	translator.proto

It has these top-level messages:
	TranslateRequest
	TranslateResponse
*/
package translatorpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TranslateRequest struct {
	Message   string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Language  string   `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	Arguments []string `protobuf:"bytes,3,rep,name=arguments" json:"arguments,omitempty"`
}

func (m *TranslateRequest) Reset()                    { *m = TranslateRequest{} }
func (m *TranslateRequest) String() string            { return proto.CompactTextString(m) }
func (*TranslateRequest) ProtoMessage()               {}
func (*TranslateRequest) Descriptor() ([]byte, []int) { return fileDescriptorTranslator, []int{0} }

func (m *TranslateRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *TranslateRequest) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *TranslateRequest) GetArguments() []string {
	if m != nil {
		return m.Arguments
	}
	return nil
}

type TranslateResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
}

func (m *TranslateResponse) Reset()                    { *m = TranslateResponse{} }
func (m *TranslateResponse) String() string            { return proto.CompactTextString(m) }
func (*TranslateResponse) ProtoMessage()               {}
func (*TranslateResponse) Descriptor() ([]byte, []int) { return fileDescriptorTranslator, []int{1} }

func (m *TranslateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *TranslateResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*TranslateRequest)(nil), "translator.TranslateRequest")
	proto.RegisterType((*TranslateResponse)(nil), "translator.TranslateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TranslatorService service

type TranslatorServiceClient interface {
	Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateResponse, error)
}

type translatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewTranslatorServiceClient(cc *grpc.ClientConn) TranslatorServiceClient {
	return &translatorServiceClient{cc}
}

func (c *translatorServiceClient) Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateResponse, error) {
	out := new(TranslateResponse)
	err := grpc.Invoke(ctx, "/translator.TranslatorService/Translate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TranslatorService service

type TranslatorServiceServer interface {
	Translate(context.Context, *TranslateRequest) (*TranslateResponse, error)
}

func RegisterTranslatorServiceServer(s *grpc.Server, srv TranslatorServiceServer) {
	s.RegisterService(&_TranslatorService_serviceDesc, srv)
}

func _TranslatorService_Translate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServiceServer).Translate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/translator.TranslatorService/Translate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServiceServer).Translate(ctx, req.(*TranslateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TranslatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "translator.TranslatorService",
	HandlerType: (*TranslatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Translate",
			Handler:    _TranslatorService_Translate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "translator.proto",
}

func init() { proto.RegisterFile("translator.proto", fileDescriptorTranslator) }

var fileDescriptorTranslator = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x29, 0x4a, 0xcc,
	0x2b, 0xce, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0xa5, 0x71, 0x09, 0x84, 0x40, 0x79, 0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42,
	0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x30, 0xae, 0x90, 0x14, 0x17, 0x47, 0x4e, 0x62, 0x5e, 0x7a, 0x29, 0x48, 0x8a, 0x09, 0x2c,
	0x05, 0xe7, 0x0b, 0xc9, 0x70, 0x71, 0x26, 0x16, 0xa5, 0x97, 0xe6, 0xa6, 0xe6, 0x95, 0x14, 0x4b,
	0x30, 0x2b, 0x30, 0x6b, 0x70, 0x06, 0x21, 0x04, 0x94, 0xdc, 0xb8, 0x04, 0x91, 0xec, 0x29, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0xc5, 0x63, 0x91, 0x38, 0x17, 0x7b, 0x6a, 0x51, 0x51, 0x7c, 0x6e, 0x71,
	0x3a, 0xd4, 0x1e, 0xb6, 0xd4, 0xa2, 0x22, 0xdf, 0xe2, 0x74, 0xa3, 0x78, 0x84, 0x39, 0xf9, 0x45,
	0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x5e, 0x5c, 0x9c, 0x70, 0xc3, 0x85, 0x64, 0xf4,
	0x90, 0x3c, 0x8c, 0xee, 0x37, 0x29, 0x59, 0x1c, 0xb2, 0x10, 0x17, 0x29, 0x31, 0x38, 0xf1, 0x45,
	0xf1, 0x20, 0x54, 0x14, 0x24, 0x25, 0xb1, 0x81, 0xc3, 0xcc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xe3, 0xf8, 0xf8, 0x8a, 0x47, 0x01, 0x00, 0x00,
}
