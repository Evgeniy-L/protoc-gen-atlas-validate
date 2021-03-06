// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/infobloxopen/protoc-gen-atlas-validate/options/atlas_validate.proto

/*
Package options is a generated protocol buffer package.

It is generated from these files:
	github.com/infobloxopen/protoc-gen-atlas-validate/options/atlas_validate.proto

It has these top-level messages:
	AtlasValidateFileOption
	AtlasValidateMethodOption
	AtlasValidateServiceOption
	AtlasValidateFieldOption
*/
package options

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AtlasValidateFieldOption_Operation int32

const (
	// Field allow only on create operation
	AtlasValidateFieldOption_create AtlasValidateFieldOption_Operation = 0
	// Field allow only on update operation
	AtlasValidateFieldOption_update AtlasValidateFieldOption_Operation = 1
)

var AtlasValidateFieldOption_Operation_name = map[int32]string{
	0: "create",
	1: "update",
}
var AtlasValidateFieldOption_Operation_value = map[string]int32{
	"create": 0,
	"update": 1,
}

func (x AtlasValidateFieldOption_Operation) String() string {
	return proto.EnumName(AtlasValidateFieldOption_Operation_name, int32(x))
}
func (AtlasValidateFieldOption_Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorAtlasValidate, []int{3, 0}
}

type AtlasValidateFileOption struct {
	AllowUnknownFields bool `protobuf:"varint,1,opt,name=allow_unknown_fields,json=allowUnknownFields,proto3" json:"allow_unknown_fields,omitempty"`
}

func (m *AtlasValidateFileOption) Reset()         { *m = AtlasValidateFileOption{} }
func (m *AtlasValidateFileOption) String() string { return proto.CompactTextString(m) }
func (*AtlasValidateFileOption) ProtoMessage()    {}
func (*AtlasValidateFileOption) Descriptor() ([]byte, []int) {
	return fileDescriptorAtlasValidate, []int{0}
}

func (m *AtlasValidateFileOption) GetAllowUnknownFields() bool {
	if m != nil {
		return m.AllowUnknownFields
	}
	return false
}

type AtlasValidateMethodOption struct {
	AllowUnknownFields bool `protobuf:"varint,1,opt,name=allow_unknown_fields,json=allowUnknownFields,proto3" json:"allow_unknown_fields,omitempty"`
}

func (m *AtlasValidateMethodOption) Reset()         { *m = AtlasValidateMethodOption{} }
func (m *AtlasValidateMethodOption) String() string { return proto.CompactTextString(m) }
func (*AtlasValidateMethodOption) ProtoMessage()    {}
func (*AtlasValidateMethodOption) Descriptor() ([]byte, []int) {
	return fileDescriptorAtlasValidate, []int{1}
}

func (m *AtlasValidateMethodOption) GetAllowUnknownFields() bool {
	if m != nil {
		return m.AllowUnknownFields
	}
	return false
}

type AtlasValidateServiceOption struct {
	AllowUnknownFields bool `protobuf:"varint,1,opt,name=allow_unknown_fields,json=allowUnknownFields,proto3" json:"allow_unknown_fields,omitempty"`
}

func (m *AtlasValidateServiceOption) Reset()         { *m = AtlasValidateServiceOption{} }
func (m *AtlasValidateServiceOption) String() string { return proto.CompactTextString(m) }
func (*AtlasValidateServiceOption) ProtoMessage()    {}
func (*AtlasValidateServiceOption) Descriptor() ([]byte, []int) {
	return fileDescriptorAtlasValidate, []int{2}
}

func (m *AtlasValidateServiceOption) GetAllowUnknownFields() bool {
	if m != nil {
		return m.AllowUnknownFields
	}
	return false
}

type AtlasValidateFieldOption struct {
	Deny     []AtlasValidateFieldOption_Operation `protobuf:"varint,1,rep,packed,name=deny,enum=atlas_validate.AtlasValidateFieldOption_Operation" json:"deny,omitempty"`
	Required []AtlasValidateFieldOption_Operation `protobuf:"varint,2,rep,packed,name=required,enum=atlas_validate.AtlasValidateFieldOption_Operation" json:"required,omitempty"`
}

func (m *AtlasValidateFieldOption) Reset()         { *m = AtlasValidateFieldOption{} }
func (m *AtlasValidateFieldOption) String() string { return proto.CompactTextString(m) }
func (*AtlasValidateFieldOption) ProtoMessage()    {}
func (*AtlasValidateFieldOption) Descriptor() ([]byte, []int) {
	return fileDescriptorAtlasValidate, []int{3}
}

func (m *AtlasValidateFieldOption) GetDeny() []AtlasValidateFieldOption_Operation {
	if m != nil {
		return m.Deny
	}
	return nil
}

func (m *AtlasValidateFieldOption) GetRequired() []AtlasValidateFieldOption_Operation {
	if m != nil {
		return m.Required
	}
	return nil
}

var E_File = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*AtlasValidateFileOption)(nil),
	Field:         52219,
	Name:          "atlas_validate.file",
	Tag:           "bytes,52219,opt,name=file",
	Filename:      "github.com/infobloxopen/protoc-gen-atlas-validate/options/atlas_validate.proto",
}

var E_Method = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*AtlasValidateMethodOption)(nil),
	Field:         52219,
	Name:          "atlas_validate.method",
	Tag:           "bytes,52219,opt,name=method",
	Filename:      "github.com/infobloxopen/protoc-gen-atlas-validate/options/atlas_validate.proto",
}

var E_Service = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*AtlasValidateServiceOption)(nil),
	Field:         52219,
	Name:          "atlas_validate.service",
	Tag:           "bytes,52219,opt,name=service",
	Filename:      "github.com/infobloxopen/protoc-gen-atlas-validate/options/atlas_validate.proto",
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*AtlasValidateFieldOption)(nil),
	Field:         52219,
	Name:          "atlas_validate.field",
	Tag:           "bytes,52219,opt,name=field",
	Filename:      "github.com/infobloxopen/protoc-gen-atlas-validate/options/atlas_validate.proto",
}

func init() {
	proto.RegisterType((*AtlasValidateFileOption)(nil), "atlas_validate.AtlasValidateFileOption")
	proto.RegisterType((*AtlasValidateMethodOption)(nil), "atlas_validate.AtlasValidateMethodOption")
	proto.RegisterType((*AtlasValidateServiceOption)(nil), "atlas_validate.AtlasValidateServiceOption")
	proto.RegisterType((*AtlasValidateFieldOption)(nil), "atlas_validate.AtlasValidateFieldOption")
	proto.RegisterEnum("atlas_validate.AtlasValidateFieldOption_Operation", AtlasValidateFieldOption_Operation_name, AtlasValidateFieldOption_Operation_value)
	proto.RegisterExtension(E_File)
	proto.RegisterExtension(E_Method)
	proto.RegisterExtension(E_Service)
	proto.RegisterExtension(E_Field)
}

func init() {
	proto.RegisterFile("github.com/infobloxopen/protoc-gen-atlas-validate/options/atlas_validate.proto", fileDescriptorAtlasValidate)
}

var fileDescriptorAtlasValidate = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0xae, 0xd2, 0x40,
	0x14, 0xc6, 0xad, 0x20, 0xe2, 0x98, 0x10, 0x32, 0x31, 0xb1, 0x12, 0xff, 0x34, 0xb8, 0xb0, 0x9a,
	0xd0, 0x1a, 0xdc, 0xd5, 0x15, 0x9a, 0xb0, 0x31, 0x40, 0x52, 0xa3, 0x0b, 0x5d, 0x34, 0xfd, 0x73,
	0x5a, 0x26, 0x0e, 0x33, 0x75, 0x3a, 0x05, 0x7d, 0x12, 0x1f, 0xcd, 0x87, 0x71, 0x63, 0x3a, 0xa5,
	0x94, 0xc2, 0xbd, 0xdc, 0x9b, 0xae, 0x3a, 0x39, 0xa7, 0xdf, 0xef, 0x3b, 0x3d, 0x5f, 0x07, 0x2d,
	0x13, 0x22, 0xd7, 0x79, 0x60, 0x85, 0x7c, 0x63, 0x13, 0x16, 0xf3, 0x80, 0xf2, 0x5f, 0x3c, 0x05,
	0x66, 0xa7, 0x82, 0x4b, 0x1e, 0x4e, 0x12, 0x60, 0x13, 0x5f, 0x52, 0x3f, 0x9b, 0x6c, 0x7d, 0x4a,
	0x22, 0x5f, 0x82, 0xcd, 0x53, 0x49, 0x38, 0xcb, 0x6c, 0x55, 0xf6, 0xaa, 0xb2, 0xa5, 0x04, 0x78,
	0xd0, 0xac, 0x8e, 0x8c, 0x84, 0xf3, 0x84, 0x42, 0x89, 0x0b, 0xf2, 0xd8, 0x8e, 0x20, 0x0b, 0x05,
	0x49, 0x25, 0x17, 0xa5, 0x62, 0xfc, 0x09, 0x3d, 0x9e, 0x15, 0x9a, 0xaf, 0x7b, 0xc9, 0x9c, 0x50,
	0x58, 0x29, 0x0b, 0xfc, 0x16, 0x3d, 0xf2, 0x29, 0xe5, 0x3b, 0x2f, 0x67, 0x3f, 0x18, 0xdf, 0x31,
	0x2f, 0x26, 0x40, 0xa3, 0x4c, 0xd7, 0x0c, 0xcd, 0xec, 0xbb, 0x58, 0xf5, 0xbe, 0x94, 0xad, 0xb9,
	0xea, 0x8c, 0x17, 0xe8, 0x49, 0x03, 0xb6, 0x00, 0xb9, 0xe6, 0x51, 0x6b, 0xdc, 0x12, 0x8d, 0x1a,
	0xb8, 0xcf, 0x20, 0xb6, 0x24, 0x6c, 0x3f, 0xde, 0x5f, 0x0d, 0xe9, 0x27, 0x1f, 0x0b, 0xb4, 0x1a,
	0x6f, 0x8e, 0xba, 0x11, 0xb0, 0xdf, 0xba, 0x66, 0x74, 0xcc, 0xc1, 0x74, 0x6a, 0x9d, 0xec, 0xf7,
	0x3a, 0x9d, 0xb5, 0x4a, 0x41, 0xf8, 0xc5, 0xc9, 0x55, 0x7a, 0xbc, 0x44, 0x7d, 0x01, 0x3f, 0x73,
	0x22, 0x20, 0xd2, 0xef, 0xb6, 0x66, 0x1d, 0x18, 0xe3, 0x97, 0xe8, 0xc1, 0xa1, 0x8c, 0x11, 0xea,
	0x85, 0x02, 0x7c, 0x09, 0xc3, 0x3b, 0xc5, 0x39, 0x4f, 0x0b, 0xc2, 0x50, 0x73, 0xbe, 0xa3, 0x6e,
	0x4c, 0x28, 0xe0, 0xa7, 0x56, 0x19, 0xb8, 0x55, 0x05, 0x6e, 0xd5, 0x79, 0x66, 0xfa, 0xbf, 0x3f,
	0x1d, 0x43, 0x33, 0x1f, 0x4e, 0x5f, 0xdd, 0x30, 0x50, 0xa5, 0x70, 0x15, 0xd4, 0x09, 0x51, 0x6f,
	0xa3, 0x82, 0xc4, 0xcf, 0xcf, 0xf0, 0xc7, 0x09, 0xd7, 0x06, 0xaf, 0x2f, 0x1a, 0x1c, 0x6b, 0xdc,
	0x3d, 0xda, 0x49, 0xd0, 0xfd, 0xac, 0x8c, 0x17, 0xbf, 0x38, 0x73, 0x69, 0x04, 0x5f, 0xdb, 0xbc,
	0xb9, 0x68, 0xd3, 0x10, 0xb9, 0x15, 0xdd, 0xf1, 0xd0, 0x3d, 0xf5, 0xa3, 0xe0, 0x67, 0x57, 0xec,
	0xea, 0x10, 0x45, 0x6d, 0x62, 0xde, 0x36, 0x3d, 0xb7, 0xe4, 0x7e, 0xf8, 0xf8, 0x6d, 0xd6, 0xfa,
	0x56, 0xbf, 0xdf, 0x3f, 0x83, 0x9e, 0x7a, 0xf5, 0xdd, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c,
	0xc1, 0xea, 0xe2, 0x21, 0x04, 0x00, 0x00,
}
