// Code generated by protoc-gen-go.
// source: project.proto
// DO NOT EDIT!

/*
Package com_cst14_im_protobuf is a generated protocol buffer package.

It is generated from these files:
	project.proto

It has these top-level messages:
	Msg
*/
package com_cst14_im_protobuf

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

type MsgType int32

const (
	MsgType_LOGIN MsgType = 0
)

var MsgType_name = map[int32]string{
	0: "LOGIN",
}
var MsgType_value = map[string]int32{
	"LOGIN": 0,
}

func (x MsgType) Enum() *MsgType {
	p := new(MsgType)
	*p = x
	return p
}
func (x MsgType) String() string {
	return proto.EnumName(MsgType_name, int32(x))
}
func (x *MsgType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MsgType_value, data, "MsgType")
	if err != nil {
		return err
	}
	*x = MsgType(value)
	return nil
}
func (MsgType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StatusCode int32

const (
	StatusCode_SUCCESS StatusCode = 0
	StatusCode_FAILED  StatusCode = 1
)

var StatusCode_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILED",
}
var StatusCode_value = map[string]int32{
	"SUCCESS": 0,
	"FAILED":  1,
}

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}
func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}
func (x *StatusCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(StatusCode_value, data, "StatusCode")
	if err != nil {
		return err
	}
	*x = StatusCode(value)
	return nil
}
func (StatusCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Msg struct {
	MsgType          *MsgType    `protobuf:"varint,1,req,name=msgType,enum=com.cst14.im.protobuf.MsgType" json:"msgType,omitempty"`
	MsgUniqueTag     *string     `protobuf:"bytes,2,opt,name=msgUniqueTag" json:"msgUniqueTag,omitempty"`
	ResponseState    *StatusCode `protobuf:"varint,3,opt,name=responseState,enum=com.cst14.im.protobuf.StatusCode" json:"responseState,omitempty"`
	Account          *string     `protobuf:"bytes,4,opt,name=account" json:"account,omitempty"`
	Pwd              *string     `protobuf:"bytes,5,opt,name=pwd" json:"pwd,omitempty"`
	Token            *string     `protobuf:"bytes,6,opt,name=token" json:"token,omitempty"`
	ErrMsg           *string     `protobuf:"bytes,7,opt,name=errMsg" json:"errMsg,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Msg) Reset()                    { *m = Msg{} }
func (m *Msg) String() string            { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()               {}
func (*Msg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Msg) GetMsgType() MsgType {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return MsgType_LOGIN
}

func (m *Msg) GetMsgUniqueTag() string {
	if m != nil && m.MsgUniqueTag != nil {
		return *m.MsgUniqueTag
	}
	return ""
}

func (m *Msg) GetResponseState() StatusCode {
	if m != nil && m.ResponseState != nil {
		return *m.ResponseState
	}
	return StatusCode_SUCCESS
}

func (m *Msg) GetAccount() string {
	if m != nil && m.Account != nil {
		return *m.Account
	}
	return ""
}

func (m *Msg) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

func (m *Msg) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *Msg) GetErrMsg() string {
	if m != nil && m.ErrMsg != nil {
		return *m.ErrMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*Msg)(nil), "com.cst14.im.protobuf.Msg")
	proto.RegisterEnum("com.cst14.im.protobuf.MsgType", MsgType_name, MsgType_value)
	proto.RegisterEnum("com.cst14.im.protobuf.StatusCode", StatusCode_name, StatusCode_value)
}

var fileDescriptor0 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x8f, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x9b, 0xc4, 0x64, 0xe9, 0xd8, 0x96, 0x30, 0x54, 0xd9, 0x93, 0xd4, 0x82, 0x50, 0x7a,
	0x58, 0xb0, 0x78, 0xf0, 0x6a, 0x63, 0x2d, 0x85, 0x56, 0x25, 0x69, 0x1f, 0x20, 0x6e, 0xd7, 0x50,
	0x35, 0xd9, 0x98, 0xdd, 0x20, 0xbe, 0x80, 0xcf, 0x6d, 0xfe, 0xb4, 0x48, 0xc0, 0x9e, 0x66, 0x7e,
	0x1f, 0xdf, 0x30, 0x33, 0xd0, 0x4d, 0x33, 0xf9, 0x26, 0xb8, 0x66, 0x45, 0xd5, 0x12, 0xcf, 0xb8,
	0x8c, 0x19, 0x57, 0xfa, 0xfa, 0x86, 0xed, 0xe2, 0x9a, 0xbd, 0xe4, 0xaf, 0xc3, 0x1f, 0x13, 0xac,
	0x95, 0x8a, 0xf0, 0x16, 0x48, 0xac, 0xa2, 0xf5, 0x77, 0x2a, 0xa8, 0x31, 0x30, 0x47, 0xbd, 0xc9,
	0x05, 0xfb, 0x77, 0x80, 0xad, 0x6a, 0xcb, 0x3f, 0xe8, 0x38, 0x84, 0x4e, 0xd1, 0x6e, 0x92, 0xdd,
	0x67, 0x2e, 0xd6, 0x61, 0x44, 0xcd, 0x81, 0x31, 0x6a, 0xfb, 0x0d, 0x86, 0x73, 0xe8, 0x66, 0x42,
	0xa5, 0x32, 0x51, 0x22, 0xd0, 0xa1, 0x16, 0xd4, 0x2a, 0xa4, 0xde, 0xe4, 0xf2, 0xc8, 0x8e, 0xd2,
	0xc9, 0x95, 0x27, 0xb7, 0xc2, 0x6f, 0xce, 0x21, 0x05, 0x12, 0x72, 0x2e, 0xf3, 0x44, 0xd3, 0x93,
	0x6a, 0xcf, 0x21, 0xa2, 0x0b, 0x56, 0xfa, 0xb5, 0xa5, 0x76, 0x45, 0xcb, 0x16, 0xfb, 0x60, 0x6b,
	0xf9, 0x2e, 0x12, 0xea, 0x54, 0xac, 0x0e, 0x78, 0x0e, 0x8e, 0xc8, 0xb2, 0xe2, 0x0b, 0x4a, 0x2a,
	0xbc, 0x4f, 0xe3, 0x3e, 0x90, 0xfd, 0x6b, 0xd8, 0x06, 0x7b, 0xf9, 0x34, 0x5f, 0x3c, 0xba, 0xad,
	0xf1, 0x15, 0xc0, 0xdf, 0x31, 0x78, 0x0a, 0x24, 0xd8, 0x78, 0xde, 0x2c, 0x08, 0xdc, 0x16, 0x02,
	0x38, 0x0f, 0x77, 0x8b, 0xe5, 0xec, 0xde, 0x35, 0xa6, 0x9d, 0x29, 0x3c, 0x97, 0xc7, 0x7b, 0x1f,
	0xa1, 0x52, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x79, 0x8c, 0xe9, 0xa6, 0x7a, 0x01, 0x00, 0x00,
}
