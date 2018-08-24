// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel/vessel.proto

package go_micro_srv_vessel

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwnerId              string   `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_vessel_6d8d19059a7c47c5, []int{0}
}
func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (dst *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(dst, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_vessel_6d8d19059a7c47c5, []int{1}
}
func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (dst *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(dst, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Created              bool      `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels,proto3" json:"vessels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_vessel_6d8d19059a7c47c5, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func init() {
	proto.RegisterType((*Vessel)(nil), "go.micro.srv.vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "go.micro.srv.vessel.Specification")
	proto.RegisterType((*Response)(nil), "go.micro.srv.vessel.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	FindAvailabel(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
	CreateVessel(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	c           client.Client
	serviceName string
}

func NewVesselServiceClient(serviceName string, c client.Client) VesselServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.vessel"
	}
	return &vesselServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vesselServiceClient) FindAvailabel(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.FindAvailabel", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vesselServiceClient) CreateVessel(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.CreateVessel", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	FindAvailabel(context.Context, *Specification, *Response) error
	CreateVessel(context.Context, *Vessel, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VesselService{hdlr}, opts...))
}

type VesselService struct {
	VesselServiceHandler
}

func (h *VesselService) FindAvailabel(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailabel(ctx, in, out)
}

func (h *VesselService) CreateVessel(ctx context.Context, in *Vessel, out *Response) error {
	return h.VesselServiceHandler.CreateVessel(ctx, in, out)
}

func init() { proto.RegisterFile("proto/vessel/vessel.proto", fileDescriptor_vessel_6d8d19059a7c47c5) }

var fileDescriptor_vessel_6d8d19059a7c47c5 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x75, 0x0b, 0x94, 0x32, 0x8a, 0x87, 0xf1, 0xb2, 0xa0, 0x24, 0x4d, 0x4f, 0x9c, 0x6a, 0x02,
	0xf1, 0x03, 0x8c, 0x89, 0x89, 0xc6, 0x53, 0x31, 0x7a, 0x24, 0xcb, 0x76, 0xc4, 0x4d, 0xda, 0x6e,
	0xd3, 0x36, 0x05, 0xff, 0xc2, 0x4f, 0xf0, 0x1f, 0xfc, 0x41, 0x93, 0x29, 0x68, 0x30, 0x04, 0x4e,
	0x3b, 0xf3, 0xe6, 0xed, 0xcb, 0x9b, 0xb7, 0x0b, 0x83, 0xbc, 0xb0, 0x95, 0xbd, 0xae, 0xa9, 0x2c,
	0x29, 0xd9, 0x1c, 0x21, 0x63, 0x78, 0xb1, 0xb4, 0x61, 0x6a, 0x74, 0x61, 0xc3, 0xb2, 0xa8, 0xc3,
	0x66, 0x14, 0x7c, 0x09, 0x70, 0x5f, 0xb8, 0xc4, 0x73, 0x70, 0x4c, 0x2c, 0x85, 0x2f, 0xc6, 0xbd,
	0xc8, 0x31, 0x31, 0x0e, 0xc1, 0xd3, 0x2a, 0x57, 0xda, 0x54, 0x1f, 0xd2, 0xf1, 0xc5, 0xb8, 0x13,
	0xfd, 0xf6, 0x38, 0x02, 0x48, 0xd5, 0x7a, 0xbe, 0x22, 0xb3, 0x7c, 0xaf, 0x64, 0x8b, 0xa7, 0xbd,
	0x54, 0xad, 0x5f, 0x19, 0x40, 0x84, 0x76, 0xa6, 0x52, 0x92, 0x6d, 0x16, 0xe3, 0x1a, 0xaf, 0xa0,
	0xa7, 0x6a, 0x65, 0x12, 0xb5, 0x48, 0x48, 0x76, 0x7c, 0x31, 0xf6, 0xa2, 0x3f, 0x00, 0x07, 0xe0,
	0xd9, 0x55, 0x46, 0xc5, 0xdc, 0xc4, 0xd2, 0xe5, 0x5b, 0x5d, 0xee, 0x1f, 0xe2, 0xe0, 0x11, 0xfa,
	0xb3, 0x9c, 0xb4, 0x79, 0x33, 0x5a, 0x55, 0xc6, 0x66, 0x3b, 0xc6, 0xc4, 0x41, 0x63, 0xce, 0x3f,
	0x63, 0xc1, 0xa7, 0x00, 0x2f, 0xa2, 0x32, 0xb7, 0x59, 0x49, 0x38, 0x05, 0xb7, 0x49, 0x81, 0x55,
	0x4e, 0x27, 0x97, 0xe1, 0x9e, 0x84, 0xc2, 0x26, 0x9d, 0x68, 0x43, 0x45, 0x09, 0x5d, 0x5d, 0x90,
	0xaa, 0x28, 0xe6, 0xb5, 0xbd, 0x68, 0xdb, 0xe2, 0x0d, 0x74, 0x1b, 0x4e, 0x29, 0x1d, 0xbf, 0x75,
	0x4c, 0x6f, 0xcb, 0x9d, 0x7c, 0x0b, 0xe8, 0x37, 0xd8, 0x8c, 0x8a, 0xda, 0x68, 0xc2, 0x67, 0xe8,
	0xdf, 0x9b, 0x2c, 0xbe, 0x6d, 0xc2, 0xa1, 0x04, 0x83, 0xbd, 0x42, 0x3b, 0xa1, 0x0c, 0x47, 0x7b,
	0x39, 0xdb, 0x5d, 0x83, 0x13, 0x7c, 0x82, 0xb3, 0x3b, 0x76, 0xba, 0x79, 0xee, 0x43, 0xee, 0x8e,
	0xaa, 0x2d, 0x5c, 0xfe, 0x53, 0xd3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x7b, 0x8f, 0x22,
	0x70, 0x02, 0x00, 0x00,
}
