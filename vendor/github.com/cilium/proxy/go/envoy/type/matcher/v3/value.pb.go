// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.29.2
// source: envoy/type/matcher/v3/value.proto

package matcherv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Specifies the way to match a ProtobufWkt::Value. Primitive values and ListValue are supported.
// StructValue is not supported and is always not matched.
// [#next-free-field: 8]
type ValueMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies how to match a value.
	//
	// Types that are assignable to MatchPattern:
	//
	//	*ValueMatcher_NullMatch_
	//	*ValueMatcher_DoubleMatch
	//	*ValueMatcher_StringMatch
	//	*ValueMatcher_BoolMatch
	//	*ValueMatcher_PresentMatch
	//	*ValueMatcher_ListMatch
	//	*ValueMatcher_OrMatch
	MatchPattern isValueMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
}

func (x *ValueMatcher) Reset() {
	*x = ValueMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_matcher_v3_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueMatcher) ProtoMessage() {}

func (x *ValueMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_matcher_v3_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueMatcher.ProtoReflect.Descriptor instead.
func (*ValueMatcher) Descriptor() ([]byte, []int) {
	return file_envoy_type_matcher_v3_value_proto_rawDescGZIP(), []int{0}
}

func (m *ValueMatcher) GetMatchPattern() isValueMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (x *ValueMatcher) GetNullMatch() *ValueMatcher_NullMatch {
	if x, ok := x.GetMatchPattern().(*ValueMatcher_NullMatch_); ok {
		return x.NullMatch
	}
	return nil
}

func (x *ValueMatcher) GetDoubleMatch() *DoubleMatcher {
	if x, ok := x.GetMatchPattern().(*ValueMatcher_DoubleMatch); ok {
		return x.DoubleMatch
	}
	return nil
}

func (x *ValueMatcher) GetStringMatch() *StringMatcher {
	if x, ok := x.GetMatchPattern().(*ValueMatcher_StringMatch); ok {
		return x.StringMatch
	}
	return nil
}

func (x *ValueMatcher) GetBoolMatch() bool {
	if x, ok := x.GetMatchPattern().(*ValueMatcher_BoolMatch); ok {
		return x.BoolMatch
	}
	return false
}

func (x *ValueMatcher) GetPresentMatch() bool {
	if x, ok := x.GetMatchPattern().(*ValueMatcher_PresentMatch); ok {
		return x.PresentMatch
	}
	return false
}

func (x *ValueMatcher) GetListMatch() *ListMatcher {
	if x, ok := x.GetMatchPattern().(*ValueMatcher_ListMatch); ok {
		return x.ListMatch
	}
	return nil
}

func (x *ValueMatcher) GetOrMatch() *OrMatcher {
	if x, ok := x.GetMatchPattern().(*ValueMatcher_OrMatch); ok {
		return x.OrMatch
	}
	return nil
}

type isValueMatcher_MatchPattern interface {
	isValueMatcher_MatchPattern()
}

type ValueMatcher_NullMatch_ struct {
	// If specified, a match occurs if and only if the target value is a NullValue.
	NullMatch *ValueMatcher_NullMatch `protobuf:"bytes,1,opt,name=null_match,json=nullMatch,proto3,oneof"`
}

type ValueMatcher_DoubleMatch struct {
	// If specified, a match occurs if and only if the target value is a double value and is
	// matched to this field.
	DoubleMatch *DoubleMatcher `protobuf:"bytes,2,opt,name=double_match,json=doubleMatch,proto3,oneof"`
}

type ValueMatcher_StringMatch struct {
	// If specified, a match occurs if and only if the target value is a string value and is
	// matched to this field.
	StringMatch *StringMatcher `protobuf:"bytes,3,opt,name=string_match,json=stringMatch,proto3,oneof"`
}

type ValueMatcher_BoolMatch struct {
	// If specified, a match occurs if and only if the target value is a bool value and is equal
	// to this field.
	BoolMatch bool `protobuf:"varint,4,opt,name=bool_match,json=boolMatch,proto3,oneof"`
}

type ValueMatcher_PresentMatch struct {
	// If specified, value match will be performed based on whether the path is referring to a
	// valid primitive value in the metadata. If the path is referring to a non-primitive value,
	// the result is always not matched.
	PresentMatch bool `protobuf:"varint,5,opt,name=present_match,json=presentMatch,proto3,oneof"`
}

type ValueMatcher_ListMatch struct {
	// If specified, a match occurs if and only if the target value is a list value and
	// is matched to this field.
	ListMatch *ListMatcher `protobuf:"bytes,6,opt,name=list_match,json=listMatch,proto3,oneof"`
}

type ValueMatcher_OrMatch struct {
	// If specified, a match occurs if and only if any of the alternatives in the match accept the value.
	OrMatch *OrMatcher `protobuf:"bytes,7,opt,name=or_match,json=orMatch,proto3,oneof"`
}

func (*ValueMatcher_NullMatch_) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_DoubleMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_StringMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_BoolMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_PresentMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_ListMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_OrMatch) isValueMatcher_MatchPattern() {}

// Specifies the way to match a list value.
type ListMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MatchPattern:
	//
	//	*ListMatcher_OneOf
	MatchPattern isListMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
}

func (x *ListMatcher) Reset() {
	*x = ListMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_matcher_v3_value_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMatcher) ProtoMessage() {}

func (x *ListMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_matcher_v3_value_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMatcher.ProtoReflect.Descriptor instead.
func (*ListMatcher) Descriptor() ([]byte, []int) {
	return file_envoy_type_matcher_v3_value_proto_rawDescGZIP(), []int{1}
}

func (m *ListMatcher) GetMatchPattern() isListMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (x *ListMatcher) GetOneOf() *ValueMatcher {
	if x, ok := x.GetMatchPattern().(*ListMatcher_OneOf); ok {
		return x.OneOf
	}
	return nil
}

type isListMatcher_MatchPattern interface {
	isListMatcher_MatchPattern()
}

type ListMatcher_OneOf struct {
	// If specified, at least one of the values in the list must match the value specified.
	OneOf *ValueMatcher `protobuf:"bytes,1,opt,name=one_of,json=oneOf,proto3,oneof"`
}

func (*ListMatcher_OneOf) isListMatcher_MatchPattern() {}

// Specifies a list of alternatives for the match.
type OrMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValueMatchers []*ValueMatcher `protobuf:"bytes,1,rep,name=value_matchers,json=valueMatchers,proto3" json:"value_matchers,omitempty"`
}

func (x *OrMatcher) Reset() {
	*x = OrMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_matcher_v3_value_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrMatcher) ProtoMessage() {}

func (x *OrMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_matcher_v3_value_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrMatcher.ProtoReflect.Descriptor instead.
func (*OrMatcher) Descriptor() ([]byte, []int) {
	return file_envoy_type_matcher_v3_value_proto_rawDescGZIP(), []int{2}
}

func (x *OrMatcher) GetValueMatchers() []*ValueMatcher {
	if x != nil {
		return x.ValueMatchers
	}
	return nil
}

// NullMatch is an empty message to specify a null value.
type ValueMatcher_NullMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ValueMatcher_NullMatch) Reset() {
	*x = ValueMatcher_NullMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_matcher_v3_value_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueMatcher_NullMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueMatcher_NullMatch) ProtoMessage() {}

func (x *ValueMatcher_NullMatch) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_matcher_v3_value_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueMatcher_NullMatch.ProtoReflect.Descriptor instead.
func (*ValueMatcher_NullMatch) Descriptor() ([]byte, []int) {
	return file_envoy_type_matcher_v3_value_proto_rawDescGZIP(), []int{0, 0}
}

var File_envoy_type_matcher_v3_value_proto protoreflect.FileDescriptor

var file_envoy_type_matcher_v3_value_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x22, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76,
	0x33, 0x2f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x04,
	0x0a, 0x0c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x4e,
	0x0a, 0x0a, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x75, 0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x49,
	0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x49, 0x0a, 0x0c, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x25, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74,
	0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0c,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x43, 0x0a, 0x0a,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x48, 0x00, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x3d, 0x0a, 0x08, 0x6f, 0x72, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x72, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x1a, 0x3d, 0x0a, 0x09, 0x4e, 0x75, 0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x3a, 0x30, 0x9a,
	0xc5, 0x88, 0x1e, 0x2b, 0x0a, 0x29, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x3a,
	0x26, 0x9a, 0xc5, 0x88, 0x1e, 0x21, 0x0a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x14, 0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x88, 0x01,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x3c, 0x0a,
	0x06, 0x6f, 0x6e, 0x65, 0x5f, 0x6f, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x48, 0x00, 0x52, 0x05, 0x6f, 0x6e, 0x65, 0x4f, 0x66, 0x3a, 0x25, 0x9a, 0xc5, 0x88,
	0x1e, 0x20, 0x0a, 0x1e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x42, 0x14, 0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x70, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x6e, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x61, 0x0a, 0x09, 0x4f, 0x72, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x0e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x02, 0x52, 0x0d, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x42, 0x83, 0x01, 0x0a, 0x23,
	0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x76, 0x33, 0x42, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x3b,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_type_matcher_v3_value_proto_rawDescOnce sync.Once
	file_envoy_type_matcher_v3_value_proto_rawDescData = file_envoy_type_matcher_v3_value_proto_rawDesc
)

func file_envoy_type_matcher_v3_value_proto_rawDescGZIP() []byte {
	file_envoy_type_matcher_v3_value_proto_rawDescOnce.Do(func() {
		file_envoy_type_matcher_v3_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_type_matcher_v3_value_proto_rawDescData)
	})
	return file_envoy_type_matcher_v3_value_proto_rawDescData
}

var file_envoy_type_matcher_v3_value_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_type_matcher_v3_value_proto_goTypes = []interface{}{
	(*ValueMatcher)(nil),           // 0: envoy.type.matcher.v3.ValueMatcher
	(*ListMatcher)(nil),            // 1: envoy.type.matcher.v3.ListMatcher
	(*OrMatcher)(nil),              // 2: envoy.type.matcher.v3.OrMatcher
	(*ValueMatcher_NullMatch)(nil), // 3: envoy.type.matcher.v3.ValueMatcher.NullMatch
	(*DoubleMatcher)(nil),          // 4: envoy.type.matcher.v3.DoubleMatcher
	(*StringMatcher)(nil),          // 5: envoy.type.matcher.v3.StringMatcher
}
var file_envoy_type_matcher_v3_value_proto_depIdxs = []int32{
	3, // 0: envoy.type.matcher.v3.ValueMatcher.null_match:type_name -> envoy.type.matcher.v3.ValueMatcher.NullMatch
	4, // 1: envoy.type.matcher.v3.ValueMatcher.double_match:type_name -> envoy.type.matcher.v3.DoubleMatcher
	5, // 2: envoy.type.matcher.v3.ValueMatcher.string_match:type_name -> envoy.type.matcher.v3.StringMatcher
	1, // 3: envoy.type.matcher.v3.ValueMatcher.list_match:type_name -> envoy.type.matcher.v3.ListMatcher
	2, // 4: envoy.type.matcher.v3.ValueMatcher.or_match:type_name -> envoy.type.matcher.v3.OrMatcher
	0, // 5: envoy.type.matcher.v3.ListMatcher.one_of:type_name -> envoy.type.matcher.v3.ValueMatcher
	0, // 6: envoy.type.matcher.v3.OrMatcher.value_matchers:type_name -> envoy.type.matcher.v3.ValueMatcher
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_envoy_type_matcher_v3_value_proto_init() }
func file_envoy_type_matcher_v3_value_proto_init() {
	if File_envoy_type_matcher_v3_value_proto != nil {
		return
	}
	file_envoy_type_matcher_v3_number_proto_init()
	file_envoy_type_matcher_v3_string_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_type_matcher_v3_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueMatcher); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_type_matcher_v3_value_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMatcher); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_type_matcher_v3_value_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrMatcher); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_type_matcher_v3_value_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueMatcher_NullMatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_envoy_type_matcher_v3_value_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ValueMatcher_NullMatch_)(nil),
		(*ValueMatcher_DoubleMatch)(nil),
		(*ValueMatcher_StringMatch)(nil),
		(*ValueMatcher_BoolMatch)(nil),
		(*ValueMatcher_PresentMatch)(nil),
		(*ValueMatcher_ListMatch)(nil),
		(*ValueMatcher_OrMatch)(nil),
	}
	file_envoy_type_matcher_v3_value_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ListMatcher_OneOf)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_type_matcher_v3_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_type_matcher_v3_value_proto_goTypes,
		DependencyIndexes: file_envoy_type_matcher_v3_value_proto_depIdxs,
		MessageInfos:      file_envoy_type_matcher_v3_value_proto_msgTypes,
	}.Build()
	File_envoy_type_matcher_v3_value_proto = out.File
	file_envoy_type_matcher_v3_value_proto_rawDesc = nil
	file_envoy_type_matcher_v3_value_proto_goTypes = nil
	file_envoy_type_matcher_v3_value_proto_depIdxs = nil
}
