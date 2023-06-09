// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

package auth

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

type UserStruct struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserStruct) Reset()         { *m = UserStruct{} }
func (m *UserStruct) String() string { return proto.CompactTextString(m) }
func (*UserStruct) ProtoMessage()    {}
func (*UserStruct) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{0}
}

func (m *UserStruct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserStruct.Unmarshal(m, b)
}
func (m *UserStruct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserStruct.Marshal(b, m, deterministic)
}
func (m *UserStruct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserStruct.Merge(m, src)
}
func (m *UserStruct) XXX_Size() int {
	return xxx_messageInfo_UserStruct.Size(m)
}
func (m *UserStruct) XXX_DiscardUnknown() {
	xxx_messageInfo_UserStruct.DiscardUnknown(m)
}

var xxx_messageInfo_UserStruct proto.InternalMessageInfo

func (m *UserStruct) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserStruct) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserStruct) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UserStruct) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserStruct) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthenticationResponse struct {
	User                 *UserStruct `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	AuthorizationToken   string      `protobuf:"bytes,2,opt,name=authorization_token,json=authorizationToken,proto3" json:"authorization_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AuthenticationResponse) Reset()         { *m = AuthenticationResponse{} }
func (m *AuthenticationResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticationResponse) ProtoMessage()    {}
func (*AuthenticationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{1}
}

func (m *AuthenticationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationResponse.Unmarshal(m, b)
}
func (m *AuthenticationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationResponse.Marshal(b, m, deterministic)
}
func (m *AuthenticationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationResponse.Merge(m, src)
}
func (m *AuthenticationResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticationResponse.Size(m)
}
func (m *AuthenticationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationResponse proto.InternalMessageInfo

func (m *AuthenticationResponse) GetUser() *UserStruct {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *AuthenticationResponse) GetAuthorizationToken() string {
	if m != nil {
		return m.AuthorizationToken
	}
	return ""
}

type RegisterRequest struct {
	User                 *UserStruct `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{2}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetUser() *UserStruct {
	if m != nil {
		return m.User
	}
	return nil
}

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{3}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LogoutRequest struct {
	AuthorizationToken   string   `protobuf:"bytes,1,opt,name=authorization_token,json=authorizationToken,proto3" json:"authorization_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{4}
}

func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (m *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(m, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

func (m *LogoutRequest) GetAuthorizationToken() string {
	if m != nil {
		return m.AuthorizationToken
	}
	return ""
}

type LogoutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{5}
}

func (m *LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResponse.Unmarshal(m, b)
}
func (m *LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResponse.Marshal(b, m, deterministic)
}
func (m *LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResponse.Merge(m, src)
}
func (m *LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_LogoutResponse.Size(m)
}
func (m *LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResponse proto.InternalMessageInfo

type AuthenticateRequest struct {
	AuthorizationToken   string   `protobuf:"bytes,1,opt,name=authorization_token,json=authorizationToken,proto3" json:"authorization_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateRequest) Reset()         { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()    {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{6}
}

func (m *AuthenticateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateRequest.Unmarshal(m, b)
}
func (m *AuthenticateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateRequest.Marshal(b, m, deterministic)
}
func (m *AuthenticateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateRequest.Merge(m, src)
}
func (m *AuthenticateRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticateRequest.Size(m)
}
func (m *AuthenticateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateRequest proto.InternalMessageInfo

func (m *AuthenticateRequest) GetAuthorizationToken() string {
	if m != nil {
		return m.AuthorizationToken
	}
	return ""
}

type AuthenticateResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateResponse) Reset()         { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()    {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{7}
}

func (m *AuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateResponse.Unmarshal(m, b)
}
func (m *AuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateResponse.Marshal(b, m, deterministic)
}
func (m *AuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateResponse.Merge(m, src)
}
func (m *AuthenticateResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticateResponse.Size(m)
}
func (m *AuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateResponse proto.InternalMessageInfo

func (m *AuthenticateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetUserRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{8}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetUserResponse struct {
	User                 *UserStruct `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{9}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetUser() *UserStruct {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*UserStruct)(nil), "auth.UserStruct")
	proto.RegisterType((*AuthenticationResponse)(nil), "auth.AuthenticationResponse")
	proto.RegisterType((*RegisterRequest)(nil), "auth.RegisterRequest")
	proto.RegisterType((*LoginRequest)(nil), "auth.LoginRequest")
	proto.RegisterType((*LogoutRequest)(nil), "auth.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "auth.LogoutResponse")
	proto.RegisterType((*AuthenticateRequest)(nil), "auth.AuthenticateRequest")
	proto.RegisterType((*AuthenticateResponse)(nil), "auth.AuthenticateResponse")
	proto.RegisterType((*GetUserRequest)(nil), "auth.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "auth.GetUserResponse")
}

func init() {
	proto.RegisterFile("proto/auth/auth.proto", fileDescriptor_82b5829f48cfb8e5)
}

var fileDescriptor_82b5829f48cfb8e5 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x0f, 0xd2, 0x40,
	0x10, 0xb5, 0x95, 0xcf, 0x01, 0x91, 0x0c, 0xa0, 0xb5, 0x6a, 0x62, 0x1a, 0x63, 0xf4, 0x02, 0x09,
	0xc6, 0x68, 0xe2, 0x05, 0x2e, 0x12, 0x13, 0xe2, 0xa1, 0xe8, 0x99, 0x54, 0x3a, 0xc2, 0x46, 0xe8,
	0xe2, 0xee, 0x56, 0x13, 0x7f, 0x82, 0x7f, 0xcc, 0xbf, 0x65, 0x76, 0xdb, 0x2d, 0xb4, 0x42, 0xfc,
	0xb8, 0x10, 0x66, 0xde, 0xbc, 0xf9, 0x78, 0x6f, 0x0b, 0xa3, 0xa3, 0xe0, 0x8a, 0x4f, 0xa2, 0x54,
	0xed, 0xcc, 0xcf, 0xd8, 0xc4, 0x58, 0xd3, 0xff, 0x83, 0x1f, 0x0e, 0xc0, 0x07, 0x49, 0x62, 0xa5,
	0x44, 0xba, 0x51, 0xd8, 0x03, 0x97, 0xc5, 0x9e, 0xf3, 0xc8, 0x79, 0xda, 0x0e, 0x5d, 0x16, 0xe3,
	0x43, 0x80, 0x4f, 0x4c, 0x48, 0xb5, 0x4e, 0xa2, 0x03, 0x79, 0xae, 0xc9, 0xb7, 0x4d, 0xe6, 0x5d,
	0x74, 0x20, 0xbc, 0x0f, 0xed, 0x7d, 0x64, 0xd1, 0x9b, 0x06, 0x6d, 0xe9, 0x84, 0x01, 0x87, 0x50,
	0xa7, 0x43, 0xc4, 0xf6, 0x5e, 0xcd, 0x00, 0x59, 0x80, 0x3e, 0xb4, 0x8e, 0x91, 0x94, 0xdf, 0xb8,
	0x88, 0xbd, 0x7a, 0xc6, 0xb0, 0x71, 0xc0, 0xe1, 0xce, 0x3c, 0x55, 0x3b, 0x4a, 0x14, 0xdb, 0x44,
	0x8a, 0xf1, 0x24, 0x24, 0x79, 0xe4, 0x89, 0x24, 0x7c, 0x0c, 0xb5, 0x54, 0x92, 0x30, 0x9b, 0x75,
	0xa6, 0xfd, 0xb1, 0xb9, 0xe3, 0xb4, 0x77, 0x68, 0x50, 0x9c, 0xc0, 0x40, 0x03, 0x5c, 0xb0, 0xef,
	0x86, 0xbe, 0x56, 0xfc, 0x33, 0x25, 0xf9, 0xda, 0x58, 0x82, 0xde, 0x6b, 0x24, 0x78, 0x09, 0xb7,
	0x43, 0xda, 0x32, 0xa9, 0x48, 0x84, 0xf4, 0x25, 0x25, 0xa9, 0xfe, 0x6e, 0x52, 0x30, 0x83, 0xee,
	0x92, 0x6f, 0x59, 0x62, 0x59, 0xc5, 0xad, 0xce, 0xb5, 0x5b, 0xdd, 0xca, 0xad, 0x33, 0xb8, 0xb5,
	0xe4, 0x5b, 0x9e, 0x2a, 0xdb, 0xe2, 0xca, 0xf2, 0xce, 0xd5, 0xe5, 0xfb, 0xd0, 0xb3, 0x1d, 0x32,
	0x95, 0x82, 0x37, 0x30, 0x38, 0xd3, 0x8f, 0xfe, 0xbb, 0xf3, 0x13, 0x18, 0x96, 0xfb, 0xe4, 0x2e,
	0x54, 0x5e, 0x47, 0xf0, 0x0c, 0x7a, 0x0b, 0x52, 0x5a, 0x1c, 0x3b, 0xea, 0x2e, 0x34, 0xb5, 0x3e,
	0xeb, 0xa2, 0xac, 0xa1, 0xc3, 0xb7, 0xb1, 0x56, 0xba, 0x28, 0xfd, 0x17, 0x4f, 0xa7, 0x3f, 0x5d,
	0xe8, 0xe8, 0x65, 0x56, 0x24, 0xbe, 0xb2, 0x0d, 0xe1, 0x1c, 0x5a, 0xd6, 0x32, 0x1c, 0x65, 0x9c,
	0x8a, 0x85, 0xfe, 0x83, 0x2c, 0x7d, 0xf9, 0x29, 0x05, 0x37, 0xf0, 0x35, 0xd4, 0x8d, 0x79, 0x88,
	0x59, 0xe1, 0xb9, 0x93, 0x7f, 0x24, 0xbf, 0x80, 0x46, 0xa6, 0x3a, 0x0e, 0x0a, 0xf6, 0xc9, 0x45,
	0x7f, 0x58, 0x4e, 0x16, 0xb4, 0x05, 0x74, 0xcf, 0x25, 0xc5, 0x7b, 0xbf, 0x8d, 0xb1, 0x76, 0xf9,
	0xfe, 0x25, 0xa8, 0x68, 0xf4, 0x0a, 0x9a, 0xb9, 0x90, 0x98, 0xcf, 0x2a, 0x5b, 0xe0, 0x8f, 0x2a,
	0x59, 0xcb, 0xfc, 0xd8, 0x30, 0xdf, 0xfd, 0xf3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x80,
	0x89, 0x79, 0x10, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error) {
	out := new(AuthenticationResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error) {
	out := new(AuthenticationResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Register(context.Context, *RegisterRequest) (*AuthenticationResponse, error)
	Login(context.Context, *LoginRequest) (*AuthenticationResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) Register(ctx context.Context, req *RegisterRequest) (*AuthenticationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedAuthServiceServer) Login(ctx context.Context, req *LoginRequest) (*AuthenticationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthServiceServer) Logout(ctx context.Context, req *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (*UnimplementedAuthServiceServer) Authenticate(ctx context.Context, req *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (*UnimplementedAuthServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AuthService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AuthService_Logout_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _AuthService_Authenticate_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _AuthService_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/auth/auth.proto",
}
