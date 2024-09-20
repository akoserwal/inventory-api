// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: kessel/inventory/v1beta1/relationships/k8spolicy_ispropagatedto_k8scluster_detail.proto

package relationships

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// the aggregate status of the cluster
type K8SPolicyIsPropagatedToK8SClusterDetail_Status int32

const (
	K8SPolicyIsPropagatedToK8SClusterDetail_STATUS_UNSPECIFIED K8SPolicyIsPropagatedToK8SClusterDetail_Status = 0
	K8SPolicyIsPropagatedToK8SClusterDetail_STATUS_OTHER       K8SPolicyIsPropagatedToK8SClusterDetail_Status = 1
	K8SPolicyIsPropagatedToK8SClusterDetail_VIOLATION          K8SPolicyIsPropagatedToK8SClusterDetail_Status = 2
	K8SPolicyIsPropagatedToK8SClusterDetail_NON_VIOLATION      K8SPolicyIsPropagatedToK8SClusterDetail_Status = 3
)

// Enum value maps for K8SPolicyIsPropagatedToK8SClusterDetail_Status.
var (
	K8SPolicyIsPropagatedToK8SClusterDetail_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_OTHER",
		2: "VIOLATION",
		3: "NON_VIOLATION",
	}
	K8SPolicyIsPropagatedToK8SClusterDetail_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_OTHER":       1,
		"VIOLATION":          2,
		"NON_VIOLATION":      3,
	}
)

func (x K8SPolicyIsPropagatedToK8SClusterDetail_Status) Enum() *K8SPolicyIsPropagatedToK8SClusterDetail_Status {
	p := new(K8SPolicyIsPropagatedToK8SClusterDetail_Status)
	*p = x
	return p
}

func (x K8SPolicyIsPropagatedToK8SClusterDetail_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (K8SPolicyIsPropagatedToK8SClusterDetail_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_enumTypes[0].Descriptor()
}

func (K8SPolicyIsPropagatedToK8SClusterDetail_Status) Type() protoreflect.EnumType {
	return &file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_enumTypes[0]
}

func (x K8SPolicyIsPropagatedToK8SClusterDetail_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use K8SPolicyIsPropagatedToK8SClusterDetail_Status.Descriptor instead.
func (K8SPolicyIsPropagatedToK8SClusterDetail_Status) EnumDescriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_rawDescGZIP(), []int{0, 0}
}

type K8SPolicyIsPropagatedToK8SClusterDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource ID assigned to the resource by Kessel Asset Inventory.
	K8SPolicyId int64 `protobuf:"varint,225679544,opt,name=k8s_policy_id,proto3" json:"k8s_policy_id,omitempty"`
	// The resource ID assigned to the resource by Kessel Asset Inventory.
	K8SClusterId int64                                          `protobuf:"varint,240280960,opt,name=k8s_cluster_id,proto3" json:"k8s_cluster_id,omitempty"`
	Status       K8SPolicyIsPropagatedToK8SClusterDetail_Status `protobuf:"varint,355610639,opt,name=status,proto3,enum=kessel.inventory.v1beta1.relationships.K8SPolicyIsPropagatedToK8SClusterDetail_Status" json:"status,omitempty"`
}

func (x *K8SPolicyIsPropagatedToK8SClusterDetail) Reset() {
	*x = K8SPolicyIsPropagatedToK8SClusterDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *K8SPolicyIsPropagatedToK8SClusterDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8SPolicyIsPropagatedToK8SClusterDetail) ProtoMessage() {}

func (x *K8SPolicyIsPropagatedToK8SClusterDetail) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8SPolicyIsPropagatedToK8SClusterDetail.ProtoReflect.Descriptor instead.
func (*K8SPolicyIsPropagatedToK8SClusterDetail) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_rawDescGZIP(), []int{0}
}

func (x *K8SPolicyIsPropagatedToK8SClusterDetail) GetK8SPolicyId() int64 {
	if x != nil {
		return x.K8SPolicyId
	}
	return 0
}

func (x *K8SPolicyIsPropagatedToK8SClusterDetail) GetK8SClusterId() int64 {
	if x != nil {
		return x.K8SClusterId
	}
	return 0
}

func (x *K8SPolicyIsPropagatedToK8SClusterDetail) GetStatus() K8SPolicyIsPropagatedToK8SClusterDetail_Status {
	if x != nil {
		return x.Status
	}
	return K8SPolicyIsPropagatedToK8SClusterDetail_STATUS_UNSPECIFIED
}

var File_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_rawDesc = []byte{
	0x0a, 0x57, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x69, 0x73, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x64, 0x74,
	0x6f, 0x5f, 0x6b, 0x38, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x6b, 0x65, 0x73, 0x73, 0x65,
	0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xde, 0x02, 0x0a, 0x27, 0x4b, 0x38, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x73, 0x50,
	0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x64, 0x54, 0x6f, 0x4b, 0x38, 0x53, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x2c, 0x0a, 0x0d, 0x6b,
	0x38, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0xb8, 0xb1, 0xce,
	0x6b, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x6b, 0x38, 0x73, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x0e, 0x6b, 0x38, 0x73,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x80, 0xcb, 0xc9, 0x72,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x6b, 0x38, 0x73, 0x5f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x7f, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x8f, 0xe0, 0xc8, 0xa9, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x56, 0x2e,
	0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x2e, 0x4b, 0x38, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x49, 0x73, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x64, 0x54, 0x6f, 0x4b, 0x38,
	0x53, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0b, 0xba, 0x48, 0x08, 0x82, 0x01, 0x05, 0x10, 0x01, 0x22,
	0x01, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x54, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0d,
	0x0a, 0x09, 0x56, 0x49, 0x4f, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x11, 0x0a,
	0x0d, 0x4e, 0x4f, 0x4e, 0x5f, 0x56, 0x49, 0x4f, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03,
	0x42, 0x8e, 0x01, 0x0a, 0x36, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x50, 0x01, 0x5a, 0x52, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2d, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x73, 0x73, 0x65,
	0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_rawDescData = file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_rawDesc
)

func file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_rawDescData)
	})
	return file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_rawDescData
}

var file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_goTypes = []any{
	(K8SPolicyIsPropagatedToK8SClusterDetail_Status)(0), // 0: kessel.inventory.v1beta1.relationships.K8SPolicyIsPropagatedToK8SClusterDetail.Status
	(*K8SPolicyIsPropagatedToK8SClusterDetail)(nil),     // 1: kessel.inventory.v1beta1.relationships.K8SPolicyIsPropagatedToK8SClusterDetail
}
var file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_depIdxs = []int32{
	0, // 0: kessel.inventory.v1beta1.relationships.K8SPolicyIsPropagatedToK8SClusterDetail.status:type_name -> kessel.inventory.v1beta1.relationships.K8SPolicyIsPropagatedToK8SClusterDetail.Status
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_init()
}
func file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_init() {
	if File_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*K8SPolicyIsPropagatedToK8SClusterDetail); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_depIdxs,
		EnumInfos:         file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_enumTypes,
		MessageInfos:      file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_msgTypes,
	}.Build()
	File_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto = out.File
	file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_rawDesc = nil
	file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_goTypes = nil
	file_kessel_inventory_v1beta1_relationships_k8spolicy_ispropagatedto_k8scluster_detail_proto_depIdxs = nil
}
