// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

/*
Package tests is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	TestSuite
	Test
	GetTest
	CreateTest
	SetTest
	UpdateTest
	UpdatePathsTest
	DeleteTest
	SetOption
	FieldPath
*/
package tests

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_firestore_v1beta14 "google.golang.org/genproto/googleapis/firestore/v1beta1"
import google_firestore_v1beta1 "google.golang.org/genproto/googleapis/firestore/v1beta1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A collection of tests.
type TestSuite struct {
	Tests []*Test `protobuf:"bytes,1,rep,name=tests" json:"tests,omitempty"`
}

func (m *TestSuite) Reset()                    { *m = TestSuite{} }
func (m *TestSuite) String() string            { return proto.CompactTextString(m) }
func (*TestSuite) ProtoMessage()               {}
func (*TestSuite) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TestSuite) GetTests() []*Test {
	if m != nil {
		return m.Tests
	}
	return nil
}

// A Test describes a single client method call and its expected result.
type Test struct {
	Description string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	// Types that are valid to be assigned to Test:
	//	*Test_Get
	//	*Test_Create
	//	*Test_Set
	//	*Test_Update
	//	*Test_UpdatePaths
	//	*Test_Delete
	Test isTest_Test `protobuf_oneof:"test"`
}

func (m *Test) Reset()                    { *m = Test{} }
func (m *Test) String() string            { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()               {}
func (*Test) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isTest_Test interface {
	isTest_Test()
}

type Test_Get struct {
	Get *GetTest `protobuf:"bytes,2,opt,name=get,oneof"`
}
type Test_Create struct {
	Create *CreateTest `protobuf:"bytes,3,opt,name=create,oneof"`
}
type Test_Set struct {
	Set *SetTest `protobuf:"bytes,4,opt,name=set,oneof"`
}
type Test_Update struct {
	Update *UpdateTest `protobuf:"bytes,5,opt,name=update,oneof"`
}
type Test_UpdatePaths struct {
	UpdatePaths *UpdatePathsTest `protobuf:"bytes,6,opt,name=update_paths,json=updatePaths,oneof"`
}
type Test_Delete struct {
	Delete *DeleteTest `protobuf:"bytes,7,opt,name=delete,oneof"`
}

func (*Test_Get) isTest_Test()         {}
func (*Test_Create) isTest_Test()      {}
func (*Test_Set) isTest_Test()         {}
func (*Test_Update) isTest_Test()      {}
func (*Test_UpdatePaths) isTest_Test() {}
func (*Test_Delete) isTest_Test()      {}

func (m *Test) GetTest() isTest_Test {
	if m != nil {
		return m.Test
	}
	return nil
}

func (m *Test) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Test) GetGet() *GetTest {
	if x, ok := m.GetTest().(*Test_Get); ok {
		return x.Get
	}
	return nil
}

func (m *Test) GetCreate() *CreateTest {
	if x, ok := m.GetTest().(*Test_Create); ok {
		return x.Create
	}
	return nil
}

func (m *Test) GetSet() *SetTest {
	if x, ok := m.GetTest().(*Test_Set); ok {
		return x.Set
	}
	return nil
}

func (m *Test) GetUpdate() *UpdateTest {
	if x, ok := m.GetTest().(*Test_Update); ok {
		return x.Update
	}
	return nil
}

func (m *Test) GetUpdatePaths() *UpdatePathsTest {
	if x, ok := m.GetTest().(*Test_UpdatePaths); ok {
		return x.UpdatePaths
	}
	return nil
}

func (m *Test) GetDelete() *DeleteTest {
	if x, ok := m.GetTest().(*Test_Delete); ok {
		return x.Delete
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Test) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Test_OneofMarshaler, _Test_OneofUnmarshaler, _Test_OneofSizer, []interface{}{
		(*Test_Get)(nil),
		(*Test_Create)(nil),
		(*Test_Set)(nil),
		(*Test_Update)(nil),
		(*Test_UpdatePaths)(nil),
		(*Test_Delete)(nil),
	}
}

func _Test_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Test)
	// test
	switch x := m.Test.(type) {
	case *Test_Get:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Get); err != nil {
			return err
		}
	case *Test_Create:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *Test_Set:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Set); err != nil {
			return err
		}
	case *Test_Update:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *Test_UpdatePaths:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpdatePaths); err != nil {
			return err
		}
	case *Test_Delete:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Delete); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Test.Test has unexpected type %T", x)
	}
	return nil
}

func _Test_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Test)
	switch tag {
	case 2: // test.get
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GetTest)
		err := b.DecodeMessage(msg)
		m.Test = &Test_Get{msg}
		return true, err
	case 3: // test.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CreateTest)
		err := b.DecodeMessage(msg)
		m.Test = &Test_Create{msg}
		return true, err
	case 4: // test.set
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SetTest)
		err := b.DecodeMessage(msg)
		m.Test = &Test_Set{msg}
		return true, err
	case 5: // test.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpdateTest)
		err := b.DecodeMessage(msg)
		m.Test = &Test_Update{msg}
		return true, err
	case 6: // test.update_paths
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpdatePathsTest)
		err := b.DecodeMessage(msg)
		m.Test = &Test_UpdatePaths{msg}
		return true, err
	case 7: // test.delete
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DeleteTest)
		err := b.DecodeMessage(msg)
		m.Test = &Test_Delete{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Test_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Test)
	// test
	switch x := m.Test.(type) {
	case *Test_Get:
		s := proto.Size(x.Get)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Test_Create:
		s := proto.Size(x.Create)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Test_Set:
		s := proto.Size(x.Set)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Test_Update:
		s := proto.Size(x.Update)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Test_UpdatePaths:
		s := proto.Size(x.UpdatePaths)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Test_Delete:
		s := proto.Size(x.Delete)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Call to the DocumentRef.Get method.
type GetTest struct {
	// The path of the doc, e.g. "projects/projectID/databases/(default)/documents/C/d"
	DocRefPath string `protobuf:"bytes,1,opt,name=doc_ref_path,json=docRefPath" json:"doc_ref_path,omitempty"`
	// The request that the call should send to the Firestore service.
	Request *google_firestore_v1beta14.GetDocumentRequest `protobuf:"bytes,2,opt,name=request" json:"request,omitempty"`
}

func (m *GetTest) Reset()                    { *m = GetTest{} }
func (m *GetTest) String() string            { return proto.CompactTextString(m) }
func (*GetTest) ProtoMessage()               {}
func (*GetTest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetTest) GetDocRefPath() string {
	if m != nil {
		return m.DocRefPath
	}
	return ""
}

func (m *GetTest) GetRequest() *google_firestore_v1beta14.GetDocumentRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

// Call to DocumentRef.Create.
type CreateTest struct {
	// The path of the doc, e.g. "projects/projectID/databases/(default)/documents/C/d"
	DocRefPath string `protobuf:"bytes,1,opt,name=doc_ref_path,json=docRefPath" json:"doc_ref_path,omitempty"`
	// The data passed to Create, as JSON. The strings "Delete" and "ServerTimestamp"
	// denote the two special sentinel values. Values that could be interpreted as integers
	// (i.e. digit strings) should be treated as integers.
	JsonData string `protobuf:"bytes,2,opt,name=json_data,json=jsonData" json:"json_data,omitempty"`
	// The request that the call should generate.
	Request *google_firestore_v1beta14.CommitRequest `protobuf:"bytes,3,opt,name=request" json:"request,omitempty"`
	// If true, the call should result in an error without generating a request.
	// If this is true, request should not be set.
	IsError bool `protobuf:"varint,4,opt,name=is_error,json=isError" json:"is_error,omitempty"`
}

func (m *CreateTest) Reset()                    { *m = CreateTest{} }
func (m *CreateTest) String() string            { return proto.CompactTextString(m) }
func (*CreateTest) ProtoMessage()               {}
func (*CreateTest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateTest) GetDocRefPath() string {
	if m != nil {
		return m.DocRefPath
	}
	return ""
}

func (m *CreateTest) GetJsonData() string {
	if m != nil {
		return m.JsonData
	}
	return ""
}

func (m *CreateTest) GetRequest() *google_firestore_v1beta14.CommitRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *CreateTest) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

// A call to DocumentRef.Set.
type SetTest struct {
	DocRefPath string                                   `protobuf:"bytes,1,opt,name=doc_ref_path,json=docRefPath" json:"doc_ref_path,omitempty"`
	Option     *SetOption                               `protobuf:"bytes,2,opt,name=option" json:"option,omitempty"`
	JsonData   string                                   `protobuf:"bytes,3,opt,name=json_data,json=jsonData" json:"json_data,omitempty"`
	Request    *google_firestore_v1beta14.CommitRequest `protobuf:"bytes,4,opt,name=request" json:"request,omitempty"`
	IsError    bool                                     `protobuf:"varint,5,opt,name=is_error,json=isError" json:"is_error,omitempty"`
}

func (m *SetTest) Reset()                    { *m = SetTest{} }
func (m *SetTest) String() string            { return proto.CompactTextString(m) }
func (*SetTest) ProtoMessage()               {}
func (*SetTest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SetTest) GetDocRefPath() string {
	if m != nil {
		return m.DocRefPath
	}
	return ""
}

func (m *SetTest) GetOption() *SetOption {
	if m != nil {
		return m.Option
	}
	return nil
}

func (m *SetTest) GetJsonData() string {
	if m != nil {
		return m.JsonData
	}
	return ""
}

func (m *SetTest) GetRequest() *google_firestore_v1beta14.CommitRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *SetTest) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

// A call to the form of DocumentRef.Update that represents the data as a map
// or dictionary.
type UpdateTest struct {
	DocRefPath   string                                   `protobuf:"bytes,1,opt,name=doc_ref_path,json=docRefPath" json:"doc_ref_path,omitempty"`
	Precondition *google_firestore_v1beta1.Precondition   `protobuf:"bytes,2,opt,name=precondition" json:"precondition,omitempty"`
	JsonData     string                                   `protobuf:"bytes,3,opt,name=json_data,json=jsonData" json:"json_data,omitempty"`
	Request      *google_firestore_v1beta14.CommitRequest `protobuf:"bytes,4,opt,name=request" json:"request,omitempty"`
	IsError      bool                                     `protobuf:"varint,5,opt,name=is_error,json=isError" json:"is_error,omitempty"`
}

func (m *UpdateTest) Reset()                    { *m = UpdateTest{} }
func (m *UpdateTest) String() string            { return proto.CompactTextString(m) }
func (*UpdateTest) ProtoMessage()               {}
func (*UpdateTest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateTest) GetDocRefPath() string {
	if m != nil {
		return m.DocRefPath
	}
	return ""
}

func (m *UpdateTest) GetPrecondition() *google_firestore_v1beta1.Precondition {
	if m != nil {
		return m.Precondition
	}
	return nil
}

func (m *UpdateTest) GetJsonData() string {
	if m != nil {
		return m.JsonData
	}
	return ""
}

func (m *UpdateTest) GetRequest() *google_firestore_v1beta14.CommitRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *UpdateTest) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

// A call to the form of DocumentRef.Update that represents the data as a list
// of field paths and their values.
type UpdatePathsTest struct {
	DocRefPath   string                                 `protobuf:"bytes,1,opt,name=doc_ref_path,json=docRefPath" json:"doc_ref_path,omitempty"`
	Precondition *google_firestore_v1beta1.Precondition `protobuf:"bytes,2,opt,name=precondition" json:"precondition,omitempty"`
	// parallel sequences: field_paths[i] corresponds to json_values[i]
	FieldPaths []*FieldPath                             `protobuf:"bytes,3,rep,name=field_paths,json=fieldPaths" json:"field_paths,omitempty"`
	JsonValues []string                                 `protobuf:"bytes,4,rep,name=json_values,json=jsonValues" json:"json_values,omitempty"`
	Request    *google_firestore_v1beta14.CommitRequest `protobuf:"bytes,5,opt,name=request" json:"request,omitempty"`
	IsError    bool                                     `protobuf:"varint,6,opt,name=is_error,json=isError" json:"is_error,omitempty"`
}

func (m *UpdatePathsTest) Reset()                    { *m = UpdatePathsTest{} }
func (m *UpdatePathsTest) String() string            { return proto.CompactTextString(m) }
func (*UpdatePathsTest) ProtoMessage()               {}
func (*UpdatePathsTest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdatePathsTest) GetDocRefPath() string {
	if m != nil {
		return m.DocRefPath
	}
	return ""
}

func (m *UpdatePathsTest) GetPrecondition() *google_firestore_v1beta1.Precondition {
	if m != nil {
		return m.Precondition
	}
	return nil
}

func (m *UpdatePathsTest) GetFieldPaths() []*FieldPath {
	if m != nil {
		return m.FieldPaths
	}
	return nil
}

func (m *UpdatePathsTest) GetJsonValues() []string {
	if m != nil {
		return m.JsonValues
	}
	return nil
}

func (m *UpdatePathsTest) GetRequest() *google_firestore_v1beta14.CommitRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *UpdatePathsTest) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

// A call to DocmentRef.Delete
type DeleteTest struct {
	DocRefPath   string                                   `protobuf:"bytes,1,opt,name=doc_ref_path,json=docRefPath" json:"doc_ref_path,omitempty"`
	Precondition *google_firestore_v1beta1.Precondition   `protobuf:"bytes,2,opt,name=precondition" json:"precondition,omitempty"`
	Request      *google_firestore_v1beta14.CommitRequest `protobuf:"bytes,3,opt,name=request" json:"request,omitempty"`
	IsError      bool                                     `protobuf:"varint,4,opt,name=is_error,json=isError" json:"is_error,omitempty"`
}

func (m *DeleteTest) Reset()                    { *m = DeleteTest{} }
func (m *DeleteTest) String() string            { return proto.CompactTextString(m) }
func (*DeleteTest) ProtoMessage()               {}
func (*DeleteTest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteTest) GetDocRefPath() string {
	if m != nil {
		return m.DocRefPath
	}
	return ""
}

func (m *DeleteTest) GetPrecondition() *google_firestore_v1beta1.Precondition {
	if m != nil {
		return m.Precondition
	}
	return nil
}

func (m *DeleteTest) GetRequest() *google_firestore_v1beta14.CommitRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *DeleteTest) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

// An option to the DocumentRef.Set call.
type SetOption struct {
	All    bool         `protobuf:"varint,1,opt,name=all" json:"all,omitempty"`
	Fields []*FieldPath `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
}

func (m *SetOption) Reset()                    { *m = SetOption{} }
func (m *SetOption) String() string            { return proto.CompactTextString(m) }
func (*SetOption) ProtoMessage()               {}
func (*SetOption) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SetOption) GetAll() bool {
	if m != nil {
		return m.All
	}
	return false
}

func (m *SetOption) GetFields() []*FieldPath {
	if m != nil {
		return m.Fields
	}
	return nil
}

// A field path.
type FieldPath struct {
	Field []string `protobuf:"bytes,1,rep,name=field" json:"field,omitempty"`
}

func (m *FieldPath) Reset()                    { *m = FieldPath{} }
func (m *FieldPath) String() string            { return proto.CompactTextString(m) }
func (*FieldPath) ProtoMessage()               {}
func (*FieldPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *FieldPath) GetField() []string {
	if m != nil {
		return m.Field
	}
	return nil
}

func init() {
	proto.RegisterType((*TestSuite)(nil), "tests.TestSuite")
	proto.RegisterType((*Test)(nil), "tests.Test")
	proto.RegisterType((*GetTest)(nil), "tests.GetTest")
	proto.RegisterType((*CreateTest)(nil), "tests.CreateTest")
	proto.RegisterType((*SetTest)(nil), "tests.SetTest")
	proto.RegisterType((*UpdateTest)(nil), "tests.UpdateTest")
	proto.RegisterType((*UpdatePathsTest)(nil), "tests.UpdatePathsTest")
	proto.RegisterType((*DeleteTest)(nil), "tests.DeleteTest")
	proto.RegisterType((*SetOption)(nil), "tests.SetOption")
	proto.RegisterType((*FieldPath)(nil), "tests.FieldPath")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x55, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0x1c, 0x27, 0x4e, 0x7c, 0x5d, 0x7d, 0x94, 0x11, 0x42, 0xa6, 0x2c, 0x70, 0x2d, 0x15,
	0x22, 0x81, 0x5c, 0xb5, 0x2c, 0x59, 0xd1, 0x84, 0x04, 0xb1, 0x21, 0x9a, 0x14, 0xb6, 0x91, 0x6b,
	0xdf, 0x04, 0x23, 0x27, 0x13, 0x3c, 0xe3, 0xbe, 0x13, 0xec, 0xd8, 0xf3, 0x10, 0x2c, 0x79, 0x11,
	0xf6, 0x68, 0x7e, 0x1c, 0xc7, 0xa0, 0x48, 0x59, 0x94, 0xb2, 0x9b, 0x39, 0xf7, 0xf7, 0x9c, 0x3b,
	0xd7, 0x06, 0x10, 0xc8, 0x45, 0xb4, 0x2e, 0x98, 0x60, 0xa4, 0x23, 0xcf, 0xfc, 0xa8, 0xbf, 0x60,
	0x6c, 0x91, 0xe3, 0xe9, 0x3c, 0x2b, 0x90, 0x0b, 0x56, 0xe0, 0xe9, 0xf5, 0xd9, 0x15, 0x8a, 0xf8,
	0xac, 0x46, 0x74, 0xc0, 0xd1, 0xc9, 0x4e, 0xcf, 0x84, 0x2d, 0x97, 0x6c, 0xa5, 0xdd, 0xc2, 0x08,
	0xdc, 0x4b, 0xe4, 0x62, 0x5a, 0x66, 0x02, 0xc9, 0x31, 0xe8, 0x32, 0xbe, 0x15, 0xd8, 0x7d, 0xef,
	0xdc, 0x8b, 0xd4, 0x2d, 0x92, 0x0e, 0x54, 0x5b, 0xc2, 0x6f, 0x2d, 0x68, 0xcb, 0x3b, 0x09, 0xc0,
	0x4b, 0x91, 0x27, 0x45, 0xb6, 0x16, 0x19, 0x5b, 0xf9, 0x56, 0x60, 0xf5, 0x5d, 0xba, 0x0d, 0x91,
	0x10, 0xec, 0x05, 0x0a, 0xbf, 0x15, 0x58, 0x7d, 0xef, 0xfc, 0x7f, 0x93, 0x6b, 0x8c, 0x42, 0x86,
	0xbf, 0xfe, 0x8f, 0x4a, 0x23, 0x79, 0x0a, 0x4e, 0x52, 0x60, 0x2c, 0xd0, 0xb7, 0x95, 0xdb, 0x5d,
	0xe3, 0x36, 0x50, 0xa0, 0xf1, 0x34, 0x2e, 0x32, 0x21, 0x47, 0xe1, 0xb7, 0x1b, 0x09, 0xa7, 0x75,
	0x42, 0xae, 0x13, 0x96, 0xeb, 0x54, 0x26, 0xec, 0x34, 0x12, 0xbe, 0x53, 0x60, 0x95, 0x50, 0xbb,
	0x90, 0x17, 0x70, 0xa0, 0x4f, 0xb3, 0x75, 0x2c, 0x3e, 0x70, 0xdf, 0x51, 0x21, 0xf7, 0x1b, 0x21,
	0x13, 0x69, 0x31, 0x71, 0x5e, 0x59, 0x43, 0xb2, 0x52, 0x8a, 0x39, 0x0a, 0xf4, 0xbb, 0x8d, 0x4a,
	0x43, 0x05, 0x56, 0x95, 0xb4, 0xcb, 0x85, 0x03, 0x6d, 0x69, 0x0d, 0x39, 0x74, 0x8d, 0x02, 0x24,
	0x80, 0x83, 0x94, 0x25, 0xb3, 0x02, 0xe7, 0xaa, 0xba, 0x51, 0x10, 0x52, 0x96, 0x50, 0x9c, 0xcb,
	0x12, 0x64, 0x04, 0xdd, 0x02, 0x3f, 0x95, 0xc8, 0x2b, 0x11, 0x9f, 0x45, 0x7a, 0xa8, 0x51, 0x3d,
	0x6c, 0x33, 0x54, 0xa9, 0xeb, 0x90, 0x25, 0xe5, 0x12, 0x57, 0x82, 0xea, 0x18, 0x5a, 0x05, 0x87,
	0x9f, 0x2d, 0x80, 0x5a, 0xd0, 0x3d, 0x0a, 0x3f, 0x04, 0xf7, 0x23, 0x67, 0xab, 0x59, 0x1a, 0x8b,
	0x58, 0x95, 0x76, 0x69, 0x4f, 0x02, 0xc3, 0x58, 0xc4, 0xe4, 0x65, 0xdd, 0x95, 0x9e, 0xd9, 0x93,
	0xdd, 0x5d, 0x0d, 0xd8, 0x72, 0x99, 0xfd, 0xd1, 0x10, 0x79, 0x00, 0xbd, 0x8c, 0xcf, 0xb0, 0x28,
	0x58, 0xa1, 0xa6, 0xd9, 0xa3, 0xdd, 0x8c, 0xbf, 0x92, 0xd7, 0xf0, 0xbb, 0x05, 0xdd, 0xe9, 0xde,
	0x0a, 0xf5, 0xc1, 0x61, 0xfa, 0xfd, 0x69, 0x81, 0x0e, 0xeb, 0x47, 0xf1, 0x56, 0xe1, 0xd4, 0xd8,
	0x9b, 0x94, 0xec, 0xdd, 0x94, 0xda, 0x37, 0x40, 0xa9, 0xd3, 0xa4, 0xf4, 0xd3, 0x02, 0xa8, 0x9f,
	0xdf, 0x1e, 0xac, 0xde, 0xc0, 0xc1, 0xba, 0xc0, 0x84, 0xad, 0xd2, 0x6c, 0x8b, 0xdb, 0xe3, 0xdd,
	0x3d, 0x4d, 0xb6, 0xbc, 0x69, 0x23, 0xf6, 0x5f, 0xf2, 0xfe, 0xda, 0x82, 0x3b, 0xbf, 0xed, 0xd0,
	0x2d, 0x93, 0x3f, 0x03, 0x6f, 0x9e, 0x61, 0x9e, 0x9a, 0xf5, 0xb6, 0xd5, 0x57, 0xad, 0x7a, 0x23,
	0x23, 0x69, 0x91, 0x25, 0x29, 0xcc, 0xab, 0x23, 0x27, 0x8f, 0xc0, 0x53, 0x7a, 0x5d, 0xc7, 0x79,
	0x89, 0xdc, 0x6f, 0x07, 0xb6, 0xec, 0x4f, 0x42, 0xef, 0x15, 0xb2, 0xad, 0x59, 0xe7, 0x06, 0x34,
	0x73, 0x9a, 0x9a, 0xfd, 0xb0, 0x00, 0xea, 0x0f, 0xc8, 0x2d, 0xcb, 0xf5, 0x77, 0x37, 0x7b, 0x0c,
	0xee, 0x66, 0x2d, 0xc9, 0x21, 0xd8, 0x71, 0x9e, 0x2b, 0x3e, 0x3d, 0x2a, 0x8f, 0x72, 0x95, 0xd5,
	0x18, 0xb8, 0xdf, 0xda, 0x31, 0x26, 0x63, 0x0f, 0x8f, 0xc1, 0xdd, 0x80, 0xe4, 0x1e, 0x74, 0x14,
	0xac, 0x7e, 0x59, 0x2e, 0xd5, 0x97, 0x8b, 0x93, 0x2f, 0xad, 0x70, 0xac, 0x5b, 0x1f, 0xe4, 0xac,
	0x4c, 0xa3, 0xd1, 0x86, 0xc0, 0xa5, 0xca, 0x3a, 0x91, 0x3f, 0xbf, 0x2b, 0x47, 0xfd, 0x03, 0x9f,
	0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x52, 0x21, 0xa1, 0x34, 0x69, 0x07, 0x00, 0x00,
}
