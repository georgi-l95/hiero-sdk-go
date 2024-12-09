// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: contract_get_info.proto

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
// Get information about a smart contract instance. This includes the account that it uses, the file
// containing its initcode (if a file was used to initialize the contract), and the time when it will expire.
type ContractGetInfoQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// standard info sent from client to node, including the signed payment, and what kind of
	// response is requested (cost, state proof, both, or neither).
	Header *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// *
	// the contract for which information is requested
	ContractID *ContractID `protobuf:"bytes,2,opt,name=contractID,proto3" json:"contractID,omitempty"`
}

func (x *ContractGetInfoQuery) Reset() {
	*x = ContractGetInfoQuery{}
	mi := &file_contract_get_info_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContractGetInfoQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractGetInfoQuery) ProtoMessage() {}

func (x *ContractGetInfoQuery) ProtoReflect() protoreflect.Message {
	mi := &file_contract_get_info_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractGetInfoQuery.ProtoReflect.Descriptor instead.
func (*ContractGetInfoQuery) Descriptor() ([]byte, []int) {
	return file_contract_get_info_proto_rawDescGZIP(), []int{0}
}

func (x *ContractGetInfoQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ContractGetInfoQuery) GetContractID() *ContractID {
	if x != nil {
		return x.ContractID
	}
	return nil
}

// *
// Response when the client sends the node ContractGetInfoQuery
type ContractGetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// standard response from node to client, including the requested fields: cost, or state proof,
	// or both, or neither
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// *
	// the information about this contract instance (a state proof can be generated for this)
	ContractInfo *ContractGetInfoResponse_ContractInfo `protobuf:"bytes,2,opt,name=contractInfo,proto3" json:"contractInfo,omitempty"`
}

func (x *ContractGetInfoResponse) Reset() {
	*x = ContractGetInfoResponse{}
	mi := &file_contract_get_info_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContractGetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractGetInfoResponse) ProtoMessage() {}

func (x *ContractGetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_get_info_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractGetInfoResponse.ProtoReflect.Descriptor instead.
func (*ContractGetInfoResponse) Descriptor() ([]byte, []int) {
	return file_contract_get_info_proto_rawDescGZIP(), []int{1}
}

func (x *ContractGetInfoResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ContractGetInfoResponse) GetContractInfo() *ContractGetInfoResponse_ContractInfo {
	if x != nil {
		return x.ContractInfo
	}
	return nil
}

type ContractGetInfoResponse_ContractInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// ID of the contract instance, in the format used in transactions
	ContractID *ContractID `protobuf:"bytes,1,opt,name=contractID,proto3" json:"contractID,omitempty"`
	// *
	// ID of the cryptocurrency account owned by the contract instance, in the format used in
	// transactions
	AccountID *AccountID `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	// *
	// ID of both the contract instance and the cryptocurrency account owned by the contract
	// instance, in the format used by Solidity
	ContractAccountID string `protobuf:"bytes,3,opt,name=contractAccountID,proto3" json:"contractAccountID,omitempty"`
	// *
	// the state of the instance and its fields can be modified arbitrarily if this key signs a
	// transaction to modify it. If this is null, then such modifications are not possible, and
	// there is no administrator that can override the normal operation of this smart contract
	// instance. Note that if it is created with no admin keys, then there is no administrator
	// to authorize changing the admin keys, so there can never be any admin keys for that
	// instance.
	AdminKey *Key `protobuf:"bytes,4,opt,name=adminKey,proto3" json:"adminKey,omitempty"`
	// *
	// the current time at which this contract instance (and its account) is set to expire
	ExpirationTime *Timestamp `protobuf:"bytes,5,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	// *
	// the expiration time will extend every this many seconds. If there are insufficient funds,
	// then it extends as long as possible. If the account is empty when it expires, then it is
	// deleted.
	AutoRenewPeriod *Duration `protobuf:"bytes,6,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`
	// *
	// number of bytes of storage being used by this instance (which affects the cost to extend
	// the expiration time)
	Storage int64 `protobuf:"varint,7,opt,name=storage,proto3" json:"storage,omitempty"`
	// *
	// the memo associated with the contract (max 100 bytes)
	Memo string `protobuf:"bytes,8,opt,name=memo,proto3" json:"memo,omitempty"`
	// *
	// The current balance, in tinybars
	Balance uint64 `protobuf:"varint,9,opt,name=balance,proto3" json:"balance,omitempty"`
	// *
	// Whether the contract has been deleted
	Deleted bool `protobuf:"varint,10,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// *
	// [DEPRECATED] The metadata of the tokens associated to the contract. This field was
	// deprecated by <a href="https://hips.hedera.com/hip/hip-367">HIP-367</a>, which allowed
	// an account to be associated to an unlimited number of tokens. This scale makes it more
	// efficient for users to consult mirror nodes to review their token associations.
	//
	// Deprecated: Marked as deprecated in contract_get_info.proto.
	TokenRelationships []*TokenRelationship `protobuf:"bytes,11,rep,name=tokenRelationships,proto3" json:"tokenRelationships,omitempty"`
	// *
	// The ledger ID the response was returned from; please see <a href="https://github.com/hashgraph/hedera-improvement-proposal/blob/master/HIP/hip-198.md">HIP-198</a> for the network-specific IDs.
	LedgerId []byte `protobuf:"bytes,12,opt,name=ledger_id,json=ledgerId,proto3" json:"ledger_id,omitempty"`
	// *
	// ID of the an account to charge for auto-renewal of this contract. If not set, or set to an account with zero hbar
	// balance, the contract's own hbar balance will be used to cover auto-renewal fees.
	AutoRenewAccountId *AccountID `protobuf:"bytes,13,opt,name=auto_renew_account_id,json=autoRenewAccountId,proto3" json:"auto_renew_account_id,omitempty"`
	// *
	// The maximum number of tokens that a contract can be implicitly associated with.
	MaxAutomaticTokenAssociations int32 `protobuf:"varint,14,opt,name=max_automatic_token_associations,json=maxAutomaticTokenAssociations,proto3" json:"max_automatic_token_associations,omitempty"`
	// *
	// Staking metadata for this contract.
	StakingInfo *StakingInfo `protobuf:"bytes,15,opt,name=staking_info,json=stakingInfo,proto3" json:"staking_info,omitempty"`
}

func (x *ContractGetInfoResponse_ContractInfo) Reset() {
	*x = ContractGetInfoResponse_ContractInfo{}
	mi := &file_contract_get_info_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContractGetInfoResponse_ContractInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractGetInfoResponse_ContractInfo) ProtoMessage() {}

func (x *ContractGetInfoResponse_ContractInfo) ProtoReflect() protoreflect.Message {
	mi := &file_contract_get_info_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractGetInfoResponse_ContractInfo.ProtoReflect.Descriptor instead.
func (*ContractGetInfoResponse_ContractInfo) Descriptor() ([]byte, []int) {
	return file_contract_get_info_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ContractGetInfoResponse_ContractInfo) GetContractID() *ContractID {
	if x != nil {
		return x.ContractID
	}
	return nil
}

func (x *ContractGetInfoResponse_ContractInfo) GetAccountID() *AccountID {
	if x != nil {
		return x.AccountID
	}
	return nil
}

func (x *ContractGetInfoResponse_ContractInfo) GetContractAccountID() string {
	if x != nil {
		return x.ContractAccountID
	}
	return ""
}

func (x *ContractGetInfoResponse_ContractInfo) GetAdminKey() *Key {
	if x != nil {
		return x.AdminKey
	}
	return nil
}

func (x *ContractGetInfoResponse_ContractInfo) GetExpirationTime() *Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

func (x *ContractGetInfoResponse_ContractInfo) GetAutoRenewPeriod() *Duration {
	if x != nil {
		return x.AutoRenewPeriod
	}
	return nil
}

func (x *ContractGetInfoResponse_ContractInfo) GetStorage() int64 {
	if x != nil {
		return x.Storage
	}
	return 0
}

func (x *ContractGetInfoResponse_ContractInfo) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *ContractGetInfoResponse_ContractInfo) GetBalance() uint64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *ContractGetInfoResponse_ContractInfo) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

// Deprecated: Marked as deprecated in contract_get_info.proto.
func (x *ContractGetInfoResponse_ContractInfo) GetTokenRelationships() []*TokenRelationship {
	if x != nil {
		return x.TokenRelationships
	}
	return nil
}

func (x *ContractGetInfoResponse_ContractInfo) GetLedgerId() []byte {
	if x != nil {
		return x.LedgerId
	}
	return nil
}

func (x *ContractGetInfoResponse_ContractInfo) GetAutoRenewAccountId() *AccountID {
	if x != nil {
		return x.AutoRenewAccountId
	}
	return nil
}

func (x *ContractGetInfoResponse_ContractInfo) GetMaxAutomaticTokenAssociations() int32 {
	if x != nil {
		return x.MaxAutomaticTokenAssociations
	}
	return 0
}

func (x *ContractGetInfoResponse_ContractInfo) GetStakingInfo() *StakingInfo {
	if x != nil {
		return x.StakingInfo
	}
	return nil
}

var File_contract_get_info_proto protoreflect.FileDescriptor

var file_contract_get_info_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0e, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x75, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x22, 0xea, 0x06, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x4f, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0xce, 0x05, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b,
	0x65, 0x79, 0x52, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x0e,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65,
	0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x12, 0x4c, 0x0a, 0x12, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x42, 0x02, 0x18, 0x01, 0x52, 0x12, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12, 0x43,
	0x0a, 0x15, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52,
	0x12, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x20, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x69, 0x63, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x6f, 0x63,
	0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1d, 0x6d,
	0x61, 0x78, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x35, 0x0a, 0x0c,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_contract_get_info_proto_rawDescOnce sync.Once
	file_contract_get_info_proto_rawDescData = file_contract_get_info_proto_rawDesc
)

func file_contract_get_info_proto_rawDescGZIP() []byte {
	file_contract_get_info_proto_rawDescOnce.Do(func() {
		file_contract_get_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_get_info_proto_rawDescData)
	})
	return file_contract_get_info_proto_rawDescData
}

var file_contract_get_info_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_contract_get_info_proto_goTypes = []any{
	(*ContractGetInfoQuery)(nil),                 // 0: proto.ContractGetInfoQuery
	(*ContractGetInfoResponse)(nil),              // 1: proto.ContractGetInfoResponse
	(*ContractGetInfoResponse_ContractInfo)(nil), // 2: proto.ContractGetInfoResponse.ContractInfo
	(*QueryHeader)(nil),                          // 3: proto.QueryHeader
	(*ContractID)(nil),                           // 4: proto.ContractID
	(*ResponseHeader)(nil),                       // 5: proto.ResponseHeader
	(*AccountID)(nil),                            // 6: proto.AccountID
	(*Key)(nil),                                  // 7: proto.Key
	(*Timestamp)(nil),                            // 8: proto.Timestamp
	(*Duration)(nil),                             // 9: proto.Duration
	(*TokenRelationship)(nil),                    // 10: proto.TokenRelationship
	(*StakingInfo)(nil),                          // 11: proto.StakingInfo
}
var file_contract_get_info_proto_depIdxs = []int32{
	3,  // 0: proto.ContractGetInfoQuery.header:type_name -> proto.QueryHeader
	4,  // 1: proto.ContractGetInfoQuery.contractID:type_name -> proto.ContractID
	5,  // 2: proto.ContractGetInfoResponse.header:type_name -> proto.ResponseHeader
	2,  // 3: proto.ContractGetInfoResponse.contractInfo:type_name -> proto.ContractGetInfoResponse.ContractInfo
	4,  // 4: proto.ContractGetInfoResponse.ContractInfo.contractID:type_name -> proto.ContractID
	6,  // 5: proto.ContractGetInfoResponse.ContractInfo.accountID:type_name -> proto.AccountID
	7,  // 6: proto.ContractGetInfoResponse.ContractInfo.adminKey:type_name -> proto.Key
	8,  // 7: proto.ContractGetInfoResponse.ContractInfo.expirationTime:type_name -> proto.Timestamp
	9,  // 8: proto.ContractGetInfoResponse.ContractInfo.autoRenewPeriod:type_name -> proto.Duration
	10, // 9: proto.ContractGetInfoResponse.ContractInfo.tokenRelationships:type_name -> proto.TokenRelationship
	6,  // 10: proto.ContractGetInfoResponse.ContractInfo.auto_renew_account_id:type_name -> proto.AccountID
	11, // 11: proto.ContractGetInfoResponse.ContractInfo.staking_info:type_name -> proto.StakingInfo
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_contract_get_info_proto_init() }
func file_contract_get_info_proto_init() {
	if File_contract_get_info_proto != nil {
		return
	}
	file_timestamp_proto_init()
	file_duration_proto_init()
	file_basic_types_proto_init()
	file_query_header_proto_init()
	file_response_header_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contract_get_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contract_get_info_proto_goTypes,
		DependencyIndexes: file_contract_get_info_proto_depIdxs,
		MessageInfos:      file_contract_get_info_proto_msgTypes,
	}.Build()
	File_contract_get_info_proto = out.File
	file_contract_get_info_proto_rawDesc = nil
	file_contract_get_info_proto_goTypes = nil
	file_contract_get_info_proto_depIdxs = nil
}
