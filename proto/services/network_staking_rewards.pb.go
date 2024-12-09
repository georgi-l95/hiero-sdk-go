// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: network_staking_rewards.proto

package services

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

// *
// Representation of a Hedera Token Service staking reward entity in the network Merkle tree.
// This consists of all the information needed to calculate the staking rewards for all nodes in the network. It is
// calculated at the beginning of each staking period for all nodes and is needed to have same values
// for reconnect.
//
// As with all network entities, staking info is per node and has a unique entity number represented as shard.realm.X.
type NetworkStakingRewards struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Whether staking rewards are activated on the network. This is set to true when the balance of 0.0.800
	// reaches minimum required balance.
	StakingRewardsActivated bool `protobuf:"varint,1,opt,name=staking_rewards_activated,json=stakingRewardsActivated,proto3" json:"staking_rewards_activated,omitempty"`
	// *
	// Total of (balance + stakedToMe) for all accounts staked to all nodes in the network with declineReward=false, at the
	// beginning of the new staking period.
	TotalStakedRewardStart int64 `protobuf:"varint,2,opt,name=total_staked_reward_start,json=totalStakedRewardStart,proto3" json:"total_staked_reward_start,omitempty"`
	// *
	// Total of (balance + stakedToMe) for all accounts staked to all nodes in the network, at the beginning of the new
	// staking period.
	TotalStakedStart int64 `protobuf:"varint,3,opt,name=total_staked_start,json=totalStakedStart,proto3" json:"total_staked_start,omitempty"`
	// *
	// The total staking rewards in tinybars that COULD be collected by all accounts staking to all nodes after the end
	// of this staking period; assuming that no account "renounces" its rewards by, for example, setting declineReward=true.
	PendingRewards int64 `protobuf:"varint,4,opt,name=pending_rewards,json=pendingRewards,proto3" json:"pending_rewards,omitempty"`
}

func (x *NetworkStakingRewards) Reset() {
	*x = NetworkStakingRewards{}
	mi := &file_network_staking_rewards_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkStakingRewards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkStakingRewards) ProtoMessage() {}

func (x *NetworkStakingRewards) ProtoReflect() protoreflect.Message {
	mi := &file_network_staking_rewards_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkStakingRewards.ProtoReflect.Descriptor instead.
func (*NetworkStakingRewards) Descriptor() ([]byte, []int) {
	return file_network_staking_rewards_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkStakingRewards) GetStakingRewardsActivated() bool {
	if x != nil {
		return x.StakingRewardsActivated
	}
	return false
}

func (x *NetworkStakingRewards) GetTotalStakedRewardStart() int64 {
	if x != nil {
		return x.TotalStakedRewardStart
	}
	return 0
}

func (x *NetworkStakingRewards) GetTotalStakedStart() int64 {
	if x != nil {
		return x.TotalStakedStart
	}
	return 0
}

func (x *NetworkStakingRewards) GetPendingRewards() int64 {
	if x != nil {
		return x.PendingRewards
	}
	return 0
}

var File_network_staking_rewards_proto protoreflect.FileDescriptor

var file_network_staking_rewards_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x01, 0x0a, 0x15, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73,
	0x12, 0x3a, 0x0a, 0x19, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x17, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x19,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x16, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x64,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x42, 0x26,
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_network_staking_rewards_proto_rawDescOnce sync.Once
	file_network_staking_rewards_proto_rawDescData = file_network_staking_rewards_proto_rawDesc
)

func file_network_staking_rewards_proto_rawDescGZIP() []byte {
	file_network_staking_rewards_proto_rawDescOnce.Do(func() {
		file_network_staking_rewards_proto_rawDescData = protoimpl.X.CompressGZIP(file_network_staking_rewards_proto_rawDescData)
	})
	return file_network_staking_rewards_proto_rawDescData
}

var file_network_staking_rewards_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_network_staking_rewards_proto_goTypes = []any{
	(*NetworkStakingRewards)(nil), // 0: proto.NetworkStakingRewards
}
var file_network_staking_rewards_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_network_staking_rewards_proto_init() }
func file_network_staking_rewards_proto_init() {
	if File_network_staking_rewards_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_network_staking_rewards_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_network_staking_rewards_proto_goTypes,
		DependencyIndexes: file_network_staking_rewards_proto_depIdxs,
		MessageInfos:      file_network_staking_rewards_proto_msgTypes,
	}.Build()
	File_network_staking_rewards_proto = out.File
	file_network_staking_rewards_proto_rawDesc = nil
	file_network_staking_rewards_proto_goTypes = nil
	file_network_staking_rewards_proto_depIdxs = nil
}
