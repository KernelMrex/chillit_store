// Code generated by protoc-gen-go. DO NOT EDIT.
// source: places.proto

package places

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

type Place struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Place) Reset()         { *m = Place{} }
func (m *Place) String() string { return proto.CompactTextString(m) }
func (*Place) ProtoMessage()    {}
func (*Place) Descriptor() ([]byte, []int) {
	return fileDescriptor_0937d2e70aaf1027, []int{0}
}

func (m *Place) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Place.Unmarshal(m, b)
}
func (m *Place) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Place.Marshal(b, m, deterministic)
}
func (m *Place) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Place.Merge(m, src)
}
func (m *Place) XXX_Size() int {
	return xxx_messageInfo_Place.Size(m)
}
func (m *Place) XXX_DiscardUnknown() {
	xxx_messageInfo_Place.DiscardUnknown(m)
}

var xxx_messageInfo_Place proto.InternalMessageInfo

func (m *Place) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Place) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Place) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Place) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type GetPlacesRequest struct {
	Offset               uint64   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Amount               uint64   `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	ShortDescription     bool     `protobuf:"varint,3,opt,name=shortDescription,proto3" json:"shortDescription,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlacesRequest) Reset()         { *m = GetPlacesRequest{} }
func (m *GetPlacesRequest) String() string { return proto.CompactTextString(m) }
func (*GetPlacesRequest) ProtoMessage()    {}
func (*GetPlacesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0937d2e70aaf1027, []int{1}
}

func (m *GetPlacesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlacesRequest.Unmarshal(m, b)
}
func (m *GetPlacesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlacesRequest.Marshal(b, m, deterministic)
}
func (m *GetPlacesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlacesRequest.Merge(m, src)
}
func (m *GetPlacesRequest) XXX_Size() int {
	return xxx_messageInfo_GetPlacesRequest.Size(m)
}
func (m *GetPlacesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlacesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlacesRequest proto.InternalMessageInfo

func (m *GetPlacesRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetPlacesRequest) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *GetPlacesRequest) GetShortDescription() bool {
	if m != nil {
		return m.ShortDescription
	}
	return false
}

type GetPlacesResponse struct {
	Places               []*Place `protobuf:"bytes,1,rep,name=places,proto3" json:"places,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPlacesResponse) Reset()         { *m = GetPlacesResponse{} }
func (m *GetPlacesResponse) String() string { return proto.CompactTextString(m) }
func (*GetPlacesResponse) ProtoMessage()    {}
func (*GetPlacesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0937d2e70aaf1027, []int{2}
}

func (m *GetPlacesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPlacesResponse.Unmarshal(m, b)
}
func (m *GetPlacesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPlacesResponse.Marshal(b, m, deterministic)
}
func (m *GetPlacesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPlacesResponse.Merge(m, src)
}
func (m *GetPlacesResponse) XXX_Size() int {
	return xxx_messageInfo_GetPlacesResponse.Size(m)
}
func (m *GetPlacesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPlacesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPlacesResponse proto.InternalMessageInfo

func (m *GetPlacesResponse) GetPlaces() []*Place {
	if m != nil {
		return m.Places
	}
	return nil
}

type AddPlaceRequest struct {
	Place                *Place   `protobuf:"bytes,1,opt,name=place,proto3" json:"place,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPlaceRequest) Reset()         { *m = AddPlaceRequest{} }
func (m *AddPlaceRequest) String() string { return proto.CompactTextString(m) }
func (*AddPlaceRequest) ProtoMessage()    {}
func (*AddPlaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0937d2e70aaf1027, []int{3}
}

func (m *AddPlaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPlaceRequest.Unmarshal(m, b)
}
func (m *AddPlaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPlaceRequest.Marshal(b, m, deterministic)
}
func (m *AddPlaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPlaceRequest.Merge(m, src)
}
func (m *AddPlaceRequest) XXX_Size() int {
	return xxx_messageInfo_AddPlaceRequest.Size(m)
}
func (m *AddPlaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPlaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddPlaceRequest proto.InternalMessageInfo

func (m *AddPlaceRequest) GetPlace() *Place {
	if m != nil {
		return m.Place
	}
	return nil
}

type AddPlaceResponse struct {
	// Types that are valid to be assigned to Result:
	//	*AddPlaceResponse_Error
	//	*AddPlaceResponse_Id
	Result               isAddPlaceResponse_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *AddPlaceResponse) Reset()         { *m = AddPlaceResponse{} }
func (m *AddPlaceResponse) String() string { return proto.CompactTextString(m) }
func (*AddPlaceResponse) ProtoMessage()    {}
func (*AddPlaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0937d2e70aaf1027, []int{4}
}

func (m *AddPlaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPlaceResponse.Unmarshal(m, b)
}
func (m *AddPlaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPlaceResponse.Marshal(b, m, deterministic)
}
func (m *AddPlaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPlaceResponse.Merge(m, src)
}
func (m *AddPlaceResponse) XXX_Size() int {
	return xxx_messageInfo_AddPlaceResponse.Size(m)
}
func (m *AddPlaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPlaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddPlaceResponse proto.InternalMessageInfo

type isAddPlaceResponse_Result interface {
	isAddPlaceResponse_Result()
}

type AddPlaceResponse_Error struct {
	Error string `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type AddPlaceResponse_Id struct {
	Id uint64 `protobuf:"varint,2,opt,name=id,proto3,oneof"`
}

func (*AddPlaceResponse_Error) isAddPlaceResponse_Result() {}

func (*AddPlaceResponse_Id) isAddPlaceResponse_Result() {}

func (m *AddPlaceResponse) GetResult() isAddPlaceResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *AddPlaceResponse) GetError() string {
	if x, ok := m.GetResult().(*AddPlaceResponse_Error); ok {
		return x.Error
	}
	return ""
}

func (m *AddPlaceResponse) GetId() uint64 {
	if x, ok := m.GetResult().(*AddPlaceResponse_Id); ok {
		return x.Id
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AddPlaceResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AddPlaceResponse_Error)(nil),
		(*AddPlaceResponse_Id)(nil),
	}
}

func init() {
	proto.RegisterType((*Place)(nil), "Place")
	proto.RegisterType((*GetPlacesRequest)(nil), "GetPlacesRequest")
	proto.RegisterType((*GetPlacesResponse)(nil), "GetPlacesResponse")
	proto.RegisterType((*AddPlaceRequest)(nil), "AddPlaceRequest")
	proto.RegisterType((*AddPlaceResponse)(nil), "AddPlaceResponse")
}

func init() { proto.RegisterFile("places.proto", fileDescriptor_0937d2e70aaf1027) }

var fileDescriptor_0937d2e70aaf1027 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x51, 0x6b, 0xfa, 0x30,
	0x14, 0xc5, 0xad, 0xda, 0xfe, 0xf5, 0xfa, 0x67, 0xab, 0x97, 0x21, 0x41, 0xc6, 0x28, 0x7d, 0x92,
	0x3d, 0x44, 0xd0, 0x4f, 0xb0, 0x31, 0x36, 0x1f, 0x47, 0xf6, 0x09, 0x9c, 0xb9, 0xb2, 0x80, 0x6b,
	0xba, 0x24, 0xfd, 0xfe, 0xa3, 0x49, 0x9c, 0xa5, 0x3e, 0x9e, 0x73, 0x73, 0x39, 0xbf, 0x73, 0x09,
	0xfc, 0xaf, 0x4f, 0xfb, 0x03, 0x59, 0x5e, 0x1b, 0xed, 0x74, 0xa9, 0x20, 0x7d, 0x6f, 0x35, 0xde,
	0xc0, 0x50, 0x49, 0x96, 0x14, 0xc9, 0x6a, 0x2c, 0x86, 0x4a, 0xe2, 0x1d, 0xa4, 0x4e, 0xb9, 0x13,
	0xb1, 0x61, 0x91, 0xac, 0xa6, 0x22, 0x08, 0x64, 0xf0, 0x6f, 0x2f, 0xa5, 0x21, 0x6b, 0xd9, 0xc8,
	0xfb, 0x67, 0x89, 0x05, 0xcc, 0x24, 0xd9, 0x83, 0x51, 0xb5, 0x53, 0xba, 0x62, 0x63, 0x3f, 0xed,
	0x5a, 0x65, 0x05, 0xf9, 0x1b, 0x39, 0x9f, 0x66, 0x05, 0xfd, 0x34, 0x64, 0x1d, 0x2e, 0x20, 0xd3,
	0xc7, 0xa3, 0x25, 0x17, 0x93, 0xa3, 0x6a, 0xfd, 0xfd, 0xb7, 0x6e, 0x2a, 0xe7, 0xe3, 0xc7, 0x22,
	0x2a, 0x7c, 0x84, 0xdc, 0x7e, 0x69, 0xe3, 0x5e, 0x3a, 0x51, 0x2d, 0xc8, 0x44, 0x5c, 0xf9, 0xe5,
	0x16, 0xe6, 0x9d, 0x3c, 0x5b, 0xeb, 0xca, 0x12, 0x3e, 0x40, 0x16, 0xfa, 0xb3, 0xa4, 0x18, 0xad,
	0x66, 0x9b, 0x8c, 0xfb, 0x07, 0x22, 0xba, 0xe5, 0x1a, 0x6e, 0x9f, 0xa4, 0x0c, 0x5e, 0x64, 0xbc,
	0x87, 0xd4, 0x0f, 0x3d, 0xe2, 0x65, 0x23, 0x98, 0xe5, 0x2b, 0xe4, 0x97, 0x85, 0x18, 0xb2, 0x80,
	0x94, 0x8c, 0xd1, 0xc6, 0x6f, 0x4c, 0x77, 0x03, 0x11, 0x24, 0xe6, 0xfe, 0xc6, 0xbe, 0xd1, 0x6e,
	0xd0, 0x5e, 0xf9, 0x79, 0x02, 0x99, 0x21, 0xdb, 0x9c, 0xdc, 0xc6, 0xc0, 0x2c, 0xa0, 0x7e, 0x38,
	0x6d, 0x08, 0x37, 0x30, 0xfd, 0x83, 0xc7, 0x39, 0xef, 0x1f, 0x6e, 0x89, 0xfc, 0xba, 0xdb, 0x1a,
	0x26, 0x67, 0x14, 0xcc, 0x79, 0xaf, 0xc6, 0x72, 0xce, 0xfb, 0x9c, 0x9f, 0x99, 0xff, 0x03, 0xdb,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xd0, 0x6e, 0x05, 0x13, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PlacesStoreClient is the client API for PlacesStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlacesStoreClient interface {
	GetPlaces(ctx context.Context, in *GetPlacesRequest, opts ...grpc.CallOption) (*GetPlacesResponse, error)
	AddPlace(ctx context.Context, in *AddPlaceRequest, opts ...grpc.CallOption) (*AddPlaceResponse, error)
}

type placesStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewPlacesStoreClient(cc grpc.ClientConnInterface) PlacesStoreClient {
	return &placesStoreClient{cc}
}

func (c *placesStoreClient) GetPlaces(ctx context.Context, in *GetPlacesRequest, opts ...grpc.CallOption) (*GetPlacesResponse, error) {
	out := new(GetPlacesResponse)
	err := c.cc.Invoke(ctx, "/PlacesStore/GetPlaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *placesStoreClient) AddPlace(ctx context.Context, in *AddPlaceRequest, opts ...grpc.CallOption) (*AddPlaceResponse, error) {
	out := new(AddPlaceResponse)
	err := c.cc.Invoke(ctx, "/PlacesStore/AddPlace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlacesStoreServer is the server API for PlacesStore service.
type PlacesStoreServer interface {
	GetPlaces(context.Context, *GetPlacesRequest) (*GetPlacesResponse, error)
	AddPlace(context.Context, *AddPlaceRequest) (*AddPlaceResponse, error)
}

// UnimplementedPlacesStoreServer can be embedded to have forward compatible implementations.
type UnimplementedPlacesStoreServer struct {
}

func (*UnimplementedPlacesStoreServer) GetPlaces(ctx context.Context, req *GetPlacesRequest) (*GetPlacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaces not implemented")
}
func (*UnimplementedPlacesStoreServer) AddPlace(ctx context.Context, req *AddPlaceRequest) (*AddPlaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPlace not implemented")
}

func RegisterPlacesStoreServer(s *grpc.Server, srv PlacesStoreServer) {
	s.RegisterService(&_PlacesStore_serviceDesc, srv)
}

func _PlacesStore_GetPlaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlacesStoreServer).GetPlaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlacesStore/GetPlaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlacesStoreServer).GetPlaces(ctx, req.(*GetPlacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlacesStore_AddPlace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPlaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlacesStoreServer).AddPlace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlacesStore/AddPlace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlacesStoreServer).AddPlace(ctx, req.(*AddPlaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlacesStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PlacesStore",
	HandlerType: (*PlacesStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlaces",
			Handler:    _PlacesStore_GetPlaces_Handler,
		},
		{
			MethodName: "AddPlace",
			Handler:    _PlacesStore_AddPlace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "places.proto",
}
