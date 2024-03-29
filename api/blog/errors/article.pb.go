// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: article.proto

package errors

import (
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

type ArticleErrorReason int32

const (
	ArticleErrorReason_UnKnownError       ArticleErrorReason = 0
	ArticleErrorReason_MissingPublishedAt ArticleErrorReason = 404
	ArticleErrorReason_MissingContent     ArticleErrorReason = 422
	ArticleErrorReason_MissingCreatedBy   ArticleErrorReason = 412
)

// Enum value maps for ArticleErrorReason.
var (
	ArticleErrorReason_name = map[int32]string{
		0:   "UnKnownError",
		404: "MissingPublishedAt",
		422: "MissingContent",
		412: "MissingCreatedBy",
	}
	ArticleErrorReason_value = map[string]int32{
		"UnKnownError":       0,
		"MissingPublishedAt": 404,
		"MissingContent":     422,
		"MissingCreatedBy":   412,
	}
)

func (x ArticleErrorReason) Enum() *ArticleErrorReason {
	p := new(ArticleErrorReason)
	*p = x
	return p
}

func (x ArticleErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArticleErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_article_proto_enumTypes[0].Descriptor()
}

func (ArticleErrorReason) Type() protoreflect.EnumType {
	return &file_article_proto_enumTypes[0]
}

func (x ArticleErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArticleErrorReason.Descriptor instead.
func (ArticleErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{0}
}

var File_article_proto protoreflect.FileDescriptor

var file_article_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2a, 0x6b, 0x0a, 0x12, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x6e, 0x4b, 0x6e, 0x6f, 0x77,
	0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x12, 0x4d, 0x69, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x10, 0x94,
	0x03, 0x12, 0x13, 0x0a, 0x0e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x10, 0xa6, 0x03, 0x12, 0x15, 0x0a, 0x10, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x10, 0x9c, 0x03, 0x42, 0x1a, 0x50,
	0x01, 0x5a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_article_proto_rawDescOnce sync.Once
	file_article_proto_rawDescData = file_article_proto_rawDesc
)

func file_article_proto_rawDescGZIP() []byte {
	file_article_proto_rawDescOnce.Do(func() {
		file_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_article_proto_rawDescData)
	})
	return file_article_proto_rawDescData
}

var file_article_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_article_proto_goTypes = []interface{}{
	(ArticleErrorReason)(0), // 0: api.blog.errors.ArticleErrorReason
}
var file_article_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_article_proto_init() }
func file_article_proto_init() {
	if File_article_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_article_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_article_proto_goTypes,
		DependencyIndexes: file_article_proto_depIdxs,
		EnumInfos:         file_article_proto_enumTypes,
	}.Build()
	File_article_proto = out.File
	file_article_proto_rawDesc = nil
	file_article_proto_goTypes = nil
	file_article_proto_depIdxs = nil
}
