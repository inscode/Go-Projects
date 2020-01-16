// Code generated by protoc-gen-go. DO NOT EDIT.
// source: person.proto

//生成go文件的包名称

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

//枚举类型
type PhoneType int32

const (
	PhoneType_MOBILE PhoneType = 0
	PhoneType_HOME   PhoneType = 1
	PhoneType_WORK   PhoneType = 2
)

var PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x PhoneType) String() string {
	return proto.EnumName(PhoneType_name, int32(x))
}

func (PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0}
}

//message 是关键字
type Person struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	//repeated表示
	Emails               []string       `protobuf:"bytes,3,rep,name=emails,proto3" json:"emails,omitempty"`
	Phones               []*PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Person) GetEmails() []string {
	if m != nil {
		return m.Emails
	}
	return nil
}

func (m *Person) GetPhones() []*PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

//为Person 的使用定义类型
type PhoneNumber struct {
	Numer                string    `protobuf:"bytes,1,opt,name=numer,proto3" json:"numer,omitempty"`
	Type                 PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=pb.PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PhoneNumber) Reset()         { *m = PhoneNumber{} }
func (m *PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*PhoneNumber) ProtoMessage()    {}
func (*PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{1}
}

func (m *PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhoneNumber.Unmarshal(m, b)
}
func (m *PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhoneNumber.Merge(m, src)
}
func (m *PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_PhoneNumber.Size(m)
}
func (m *PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_PhoneNumber proto.InternalMessageInfo

func (m *PhoneNumber) GetNumer() string {
	if m != nil {
		return m.Numer
	}
	return ""
}

func (m *PhoneNumber) GetType() PhoneType {
	if m != nil {
		return m.Type
	}
	return PhoneType_MOBILE
}

func init() {
	proto.RegisterEnum("pb.PhoneType", PhoneType_name, PhoneType_value)
	proto.RegisterType((*Person)(nil), "pb.Person")
	proto.RegisterType((*PhoneNumber)(nil), "pb.PhoneNumber")
}

func init() { proto.RegisterFile("person.proto", fileDescriptor_4c9e10cf24b1156d) }

var fileDescriptor_4c9e10cf24b1156d = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x86, 0xdd, 0x4d, 0x1a, 0xdc, 0xa9, 0x1f, 0x61, 0x10, 0xc9, 0x31, 0xf6, 0x62, 0x50, 0xd8,
	0x43, 0xfd, 0x07, 0x42, 0x45, 0xd1, 0xba, 0x25, 0x08, 0x9e, 0x37, 0x30, 0xa8, 0x60, 0x3e, 0xc8,
	0xb6, 0x87, 0xfe, 0x7b, 0xd9, 0x74, 0x29, 0xde, 0x9e, 0xf7, 0x79, 0x61, 0x5e, 0x06, 0xce, 0x12,
	0xe5, 0x21, 0x86, 0x36, 0xe5, 0xb8, 0x8d, 0x58, 0x27, 0xb7, 0x88, 0x20, 0x36, 0xc5, 0x21, 0x02,
	0x0f, 0xbd, 0x27, 0x55, 0xe9, 0xca, 0x34, 0xb6, 0x30, 0x4a, 0x60, 0xfd, 0x17, 0xa9, 0x5a, 0x57,
	0x66, 0x66, 0x47, 0xc4, 0x6b, 0x10, 0xe4, 0xfb, 0x9f, 0xdf, 0x41, 0x31, 0xcd, 0x4c, 0x63, 0xa7,
	0x84, 0xb7, 0x20, 0xd2, 0x77, 0x0c, 0x34, 0x28, 0xae, 0x99, 0x99, 0x2f, 0x2f, 0xdb, 0xe4, 0xda,
	0xcd, 0x68, 0xde, 0x77, 0xde, 0x51, 0xb6, 0x53, 0xbd, 0x78, 0x82, 0xf9, 0x3f, 0x8d, 0x57, 0x30,
	0x0b, 0x3b, 0x4f, 0x79, 0x9a, 0x3d, 0x04, 0xbc, 0x01, 0xbe, 0xdd, 0xa7, 0xc3, 0xf0, 0xc5, 0xf2,
	0xfc, 0x78, 0xeb, 0x63, 0x9f, 0xc8, 0x96, 0xea, 0xee, 0x1e, 0x9a, 0xa3, 0x42, 0x00, 0xb1, 0xee,
	0x1e, 0x5f, 0xde, 0x56, 0xf2, 0x04, 0x4f, 0x81, 0x3f, 0x77, 0xeb, 0x95, 0xac, 0x46, 0xfa, 0xec,
	0xec, 0xab, 0xac, 0x9d, 0x28, 0x0f, 0x3f, 0xfc, 0x05, 0x00, 0x00, 0xff, 0xff, 0x65, 0xd1, 0x4b,
	0x6a, 0x00, 0x01, 0x00, 0x00,
}
