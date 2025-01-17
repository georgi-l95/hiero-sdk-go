// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: schedule.proto

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
// Representation of a Hedera Schedule entry in the network Merkle tree.
//
// As with all network entities, a schedule has a unique entity number, which is usually given along
// with the network's shard and realm in the form of a shard.realm.number id.
type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// This schedule's unique ID within the global network state.
	ScheduleId *ScheduleID `protobuf:"bytes,1,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	// *
	// The schedule deleted flag
	// A schedule will either be executed or deleted, but never both.
	Deleted bool `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// *
	// The schedule executed flag
	// A schedule will either be executed or deleted, but never both.
	Executed bool `protobuf:"varint,3,opt,name=executed,proto3" json:"executed,omitempty"`
	// *
	// The schedule flag to wait for expiration
	// A schedule will be executed immediately when all necessary signatures are gathered, unless
	// this flag is set.  If this flag is set, the schedule will wait until the consensus time
	// reaches expiration_time_provided, when signatures will again be verified, and if all
	// required signatures are present at that time, the schedule will be executed.  Otherwise
	// the schedule will expire without execution.
	// Note that a schedule is always removed from state when it expires, regardless of whether it
	// was executed or not.
	WaitForExpiry bool `protobuf:"varint,4,opt,name=wait_for_expiry,json=waitForExpiry,proto3" json:"wait_for_expiry,omitempty"`
	// *
	// The memo associated with this schedule.
	Memo string `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`
	// *
	// The schedule account for this schedule.  This is the account that submitted the original
	// ScheduleCreate transaction.
	SchedulerAccountId *AccountID `protobuf:"bytes,6,opt,name=scheduler_account_id,json=schedulerAccountId,proto3" json:"scheduler_account_id,omitempty"`
	// *
	// The explicit payer account for the scheduled transaction.
	// This account is added to the accounts that must sign the schedule before it will execute.
	PayerAccountId *AccountID `protobuf:"bytes,7,opt,name=payer_account_id,json=payerAccountId,proto3" json:"payer_account_id,omitempty"`
	// *
	// The admin key for this schedule.
	// If this is not set, then the schedule cannot be deleted.
	AdminKey *Key `protobuf:"bytes,8,opt,name=admin_key,json=adminKey,proto3" json:"admin_key,omitempty"`
	// *
	// The transaction valid start value from the transaction that created this schedule.
	ScheduleValidStart *Timestamp `protobuf:"bytes,9,opt,name=schedule_valid_start,json=scheduleValidStart,proto3" json:"schedule_valid_start,omitempty"`
	// *
	// The requested expiration time of the schedule as provided by the user.
	// The actual calculated expiration time may be "earlier" than this, but will not be later.
	ProvidedExpirationSecond int64 `protobuf:"varint,10,opt,name=provided_expiration_second,json=providedExpirationSecond,proto3" json:"provided_expiration_second,omitempty"`
	// *
	// The calculated expiration time of the schedule.  This is calculated based on the requested
	// expiration time from the create transaction, and the maximum expiration permitted by the
	// network.
	// The schedule will be removed from global network state after the network reaches a consensus
	// time greater than or equal to this value.
	CalculatedExpirationSecond int64 `protobuf:"varint,11,opt,name=calculated_expiration_second,json=calculatedExpirationSecond,proto3" json:"calculated_expiration_second,omitempty"`
	// *
	// The consensus timestamp of the transaction that executed or deleted this schedule.
	ResolutionTime *Timestamp `protobuf:"bytes,12,opt,name=resolution_time,json=resolutionTime,proto3" json:"resolution_time,omitempty"`
	// *
	// The scheduled transaction to execute.
	ScheduledTransaction *SchedulableTransactionBody `protobuf:"bytes,13,opt,name=scheduled_transaction,json=scheduledTransaction,proto3" json:"scheduled_transaction,omitempty"`
	// *
	// The full transaction that created this schedule.  This is primarily used for duplicate
	// schedule create detection.  This is also the source of the parent transaction ID, from
	// which the child transaction ID is derived.
	OriginalCreateTransaction *TransactionBody `protobuf:"bytes,14,opt,name=original_create_transaction,json=originalCreateTransaction,proto3" json:"original_create_transaction,omitempty"`
	// *
	// All the primitive keys that have already signed this schedule.
	// The scheduled transaction will not be executed before this list is
	// sufficient to "activate" the required keys for the scheduled transaction.
	Signatories []*Key `protobuf:"bytes,15,rep,name=signatories,proto3" json:"signatories,omitempty"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	mi := &file_schedule_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{0}
}

func (x *Schedule) GetScheduleId() *ScheduleID {
	if x != nil {
		return x.ScheduleId
	}
	return nil
}

func (x *Schedule) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *Schedule) GetExecuted() bool {
	if x != nil {
		return x.Executed
	}
	return false
}

func (x *Schedule) GetWaitForExpiry() bool {
	if x != nil {
		return x.WaitForExpiry
	}
	return false
}

func (x *Schedule) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Schedule) GetSchedulerAccountId() *AccountID {
	if x != nil {
		return x.SchedulerAccountId
	}
	return nil
}

func (x *Schedule) GetPayerAccountId() *AccountID {
	if x != nil {
		return x.PayerAccountId
	}
	return nil
}

func (x *Schedule) GetAdminKey() *Key {
	if x != nil {
		return x.AdminKey
	}
	return nil
}

func (x *Schedule) GetScheduleValidStart() *Timestamp {
	if x != nil {
		return x.ScheduleValidStart
	}
	return nil
}

func (x *Schedule) GetProvidedExpirationSecond() int64 {
	if x != nil {
		return x.ProvidedExpirationSecond
	}
	return 0
}

func (x *Schedule) GetCalculatedExpirationSecond() int64 {
	if x != nil {
		return x.CalculatedExpirationSecond
	}
	return 0
}

func (x *Schedule) GetResolutionTime() *Timestamp {
	if x != nil {
		return x.ResolutionTime
	}
	return nil
}

func (x *Schedule) GetScheduledTransaction() *SchedulableTransactionBody {
	if x != nil {
		return x.ScheduledTransaction
	}
	return nil
}

func (x *Schedule) GetOriginalCreateTransaction() *TransactionBody {
	if x != nil {
		return x.OriginalCreateTransaction
	}
	return nil
}

func (x *Schedule) GetSignatories() []*Key {
	if x != nil {
		return x.Signatories
	}
	return nil
}

// *
// A message for storing a list of schedules in state.
// This is used to store lists of schedules that expire at a particular time or that have the same
// simplified hash code.
type ScheduleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// a list of schedules, in no particular order.
	Schedules []*Schedule `protobuf:"bytes,1,rep,name=schedules,proto3" json:"schedules,omitempty"`
}

func (x *ScheduleList) Reset() {
	*x = ScheduleList{}
	mi := &file_schedule_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScheduleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleList) ProtoMessage() {}

func (x *ScheduleList) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleList.ProtoReflect.Descriptor instead.
func (*ScheduleList) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{1}
}

func (x *ScheduleList) GetSchedules() []*Schedule {
	if x != nil {
		return x.Schedules
	}
	return nil
}

// *
// A message for storing a list of schedule identifiers in state.<br/>
// This is used to store lists of `ScheduleID` values.
// One example is all schedules that expire at a particular time.
type ScheduleIdList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// A list of schedule identifiers, in no particular order.
	// <p>
	// While the order is not _specified_, it MUST be deterministic.
	ScheduleIds []*ScheduleID `protobuf:"bytes,1,rep,name=schedule_ids,json=scheduleIds,proto3" json:"schedule_ids,omitempty"`
}

func (x *ScheduleIdList) Reset() {
	*x = ScheduleIdList{}
	mi := &file_schedule_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScheduleIdList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleIdList) ProtoMessage() {}

func (x *ScheduleIdList) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleIdList.ProtoReflect.Descriptor instead.
func (*ScheduleIdList) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{2}
}

func (x *ScheduleIdList) GetScheduleIds() []*ScheduleID {
	if x != nil {
		return x.ScheduleIds
	}
	return nil
}

// *
// The value of a map summarizing the counts of scheduled and processed transactions
// within a particular consensus second.
type ScheduledCounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The number of transactions scheduled to expire at a consensus second.
	NumberScheduled uint32 `protobuf:"varint,1,opt,name=number_scheduled,json=numberScheduled,proto3" json:"number_scheduled,omitempty"`
	// *
	// The number of scheduled transactions that have been processed at a consensus second.
	NumberProcessed uint32 `protobuf:"varint,2,opt,name=number_processed,json=numberProcessed,proto3" json:"number_processed,omitempty"`
}

func (x *ScheduledCounts) Reset() {
	*x = ScheduledCounts{}
	mi := &file_schedule_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScheduledCounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduledCounts) ProtoMessage() {}

func (x *ScheduledCounts) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduledCounts.ProtoReflect.Descriptor instead.
func (*ScheduledCounts) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{3}
}

func (x *ScheduledCounts) GetNumberScheduled() uint32 {
	if x != nil {
		return x.NumberScheduled
	}
	return 0
}

func (x *ScheduledCounts) GetNumberProcessed() uint32 {
	if x != nil {
		return x.NumberProcessed
	}
	return 0
}

// *
// A key mapping to a particular ScheduleID that will execute at a given order number
// within a given consensus second.
type ScheduledOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The consensus second in which the transaction is expired.
	ExpirySecond uint64 `protobuf:"varint,1,opt,name=expiry_second,json=expirySecond,proto3" json:"expiry_second,omitempty"`
	// The ordered position within the consensus second that the transaction will be executed.
	OrderNumber uint32 `protobuf:"varint,2,opt,name=order_number,json=orderNumber,proto3" json:"order_number,omitempty"`
}

func (x *ScheduledOrder) Reset() {
	*x = ScheduledOrder{}
	mi := &file_schedule_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScheduledOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduledOrder) ProtoMessage() {}

func (x *ScheduledOrder) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduledOrder.ProtoReflect.Descriptor instead.
func (*ScheduledOrder) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{4}
}

func (x *ScheduledOrder) GetExpirySecond() uint64 {
	if x != nil {
		return x.ExpirySecond
	}
	return 0
}

func (x *ScheduledOrder) GetOrderNumber() uint32 {
	if x != nil {
		return x.OrderNumber
	}
	return 0
}

var File_schedule_proto protoreflect.FileDescriptor

var file_schedule_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x6f, 0x64,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x06, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x0a, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x12, 0x26,
	0x0a, 0x0f, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x77, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72,
	0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x42, 0x0a, 0x14, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x12, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3a,
	0x0a, 0x10, 0x70, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x09, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x4b, 0x65, 0x79, 0x12, 0x42, 0x0a, 0x14, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x12, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x3c, 0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x18, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x64, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x40, 0x0a, 0x1c, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1a, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x39, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x56, 0x0a, 0x15, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x6f, 0x64, 0x79, 0x52, 0x14, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x56, 0x0a, 0x1b, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x19, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4b, 0x65, 0x79, 0x52, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x22, 0x3d, 0x0a, 0x0c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x2d, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x52, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x22,
	0x46, 0x0a, 0x0e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x34, 0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x0b, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x22, 0x67, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64,
	0x22, 0x58, 0x0a, 0x0e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x79, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f,
	0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61,
	0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schedule_proto_rawDescOnce sync.Once
	file_schedule_proto_rawDescData = file_schedule_proto_rawDesc
)

func file_schedule_proto_rawDescGZIP() []byte {
	file_schedule_proto_rawDescOnce.Do(func() {
		file_schedule_proto_rawDescData = protoimpl.X.CompressGZIP(file_schedule_proto_rawDescData)
	})
	return file_schedule_proto_rawDescData
}

var file_schedule_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_schedule_proto_goTypes = []any{
	(*Schedule)(nil),                   // 0: proto.Schedule
	(*ScheduleList)(nil),               // 1: proto.ScheduleList
	(*ScheduleIdList)(nil),             // 2: proto.ScheduleIdList
	(*ScheduledCounts)(nil),            // 3: proto.ScheduledCounts
	(*ScheduledOrder)(nil),             // 4: proto.ScheduledOrder
	(*ScheduleID)(nil),                 // 5: proto.ScheduleID
	(*AccountID)(nil),                  // 6: proto.AccountID
	(*Key)(nil),                        // 7: proto.Key
	(*Timestamp)(nil),                  // 8: proto.Timestamp
	(*SchedulableTransactionBody)(nil), // 9: proto.SchedulableTransactionBody
	(*TransactionBody)(nil),            // 10: proto.TransactionBody
}
var file_schedule_proto_depIdxs = []int32{
	5,  // 0: proto.Schedule.schedule_id:type_name -> proto.ScheduleID
	6,  // 1: proto.Schedule.scheduler_account_id:type_name -> proto.AccountID
	6,  // 2: proto.Schedule.payer_account_id:type_name -> proto.AccountID
	7,  // 3: proto.Schedule.admin_key:type_name -> proto.Key
	8,  // 4: proto.Schedule.schedule_valid_start:type_name -> proto.Timestamp
	8,  // 5: proto.Schedule.resolution_time:type_name -> proto.Timestamp
	9,  // 6: proto.Schedule.scheduled_transaction:type_name -> proto.SchedulableTransactionBody
	10, // 7: proto.Schedule.original_create_transaction:type_name -> proto.TransactionBody
	7,  // 8: proto.Schedule.signatories:type_name -> proto.Key
	0,  // 9: proto.ScheduleList.schedules:type_name -> proto.Schedule
	5,  // 10: proto.ScheduleIdList.schedule_ids:type_name -> proto.ScheduleID
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_schedule_proto_init() }
func file_schedule_proto_init() {
	if File_schedule_proto != nil {
		return
	}
	file_basic_types_proto_init()
	file_timestamp_proto_init()
	file_schedulable_transaction_body_proto_init()
	file_transaction_body_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schedule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schedule_proto_goTypes,
		DependencyIndexes: file_schedule_proto_depIdxs,
		MessageInfos:      file_schedule_proto_msgTypes,
	}.Build()
	File_schedule_proto = out.File
	file_schedule_proto_rawDesc = nil
	file_schedule_proto_goTypes = nil
	file_schedule_proto_depIdxs = nil
}
