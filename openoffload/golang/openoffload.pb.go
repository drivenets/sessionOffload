// Code generated by protoc-gen-go. DO NOT EDIT.
// source: openoffload.proto

package openoffload_v1alpha2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type IP_VERSION int32

const (
	IP_VERSION__IPV4 IP_VERSION = 0
	IP_VERSION__IPV6 IP_VERSION = 1
)

var IP_VERSION_name = map[int32]string{
	0: "_IPV4",
	1: "_IPV6",
}

var IP_VERSION_value = map[string]int32{
	"_IPV4": 0,
	"_IPV6": 1,
}

func (x IP_VERSION) String() string {
	return proto.EnumName(IP_VERSION_name, int32(x))
}

func (IP_VERSION) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{0}
}

type PROTOCOL_ID int32

const (
	PROTOCOL_ID__HOPOPT PROTOCOL_ID = 0
	PROTOCOL_ID__TCP    PROTOCOL_ID = 6
	PROTOCOL_ID__UDP    PROTOCOL_ID = 17
)

var PROTOCOL_ID_name = map[int32]string{
	0:  "_HOPOPT",
	6:  "_TCP",
	17: "_UDP",
}

var PROTOCOL_ID_value = map[string]int32{
	"_HOPOPT": 0,
	"_TCP":    6,
	"_UDP":    17,
}

func (x PROTOCOL_ID) String() string {
	return proto.EnumName(PROTOCOL_ID_name, int32(x))
}

func (PROTOCOL_ID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{1}
}

type SESSION_STATE int32

const (
	SESSION_STATE__ESTABLISHED SESSION_STATE = 0
	SESSION_STATE__CLOSING_1   SESSION_STATE = 1
	SESSION_STATE__CLOSING_2   SESSION_STATE = 2
	SESSION_STATE__CLOSED      SESSION_STATE = 3
)

var SESSION_STATE_name = map[int32]string{
	0: "_ESTABLISHED",
	1: "_CLOSING_1",
	2: "_CLOSING_2",
	3: "_CLOSED",
}

var SESSION_STATE_value = map[string]int32{
	"_ESTABLISHED": 0,
	"_CLOSING_1":   1,
	"_CLOSING_2":   2,
	"_CLOSED":      3,
}

func (x SESSION_STATE) String() string {
	return proto.EnumName(SESSION_STATE_name, int32(x))
}

func (SESSION_STATE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{2}
}

type SESSION_CLOSE_CODE int32

const (
	SESSION_CLOSE_CODE__NOT_CLOSED SESSION_CLOSE_CODE = 0
	SESSION_CLOSE_CODE__FINACK     SESSION_CLOSE_CODE = 1
	SESSION_CLOSE_CODE__RST        SESSION_CLOSE_CODE = 2
	SESSION_CLOSE_CODE__TIMEOUT    SESSION_CLOSE_CODE = 3
)

var SESSION_CLOSE_CODE_name = map[int32]string{
	0: "_NOT_CLOSED",
	1: "_FINACK",
	2: "_RST",
	3: "_TIMEOUT",
}

var SESSION_CLOSE_CODE_value = map[string]int32{
	"_NOT_CLOSED": 0,
	"_FINACK":     1,
	"_RST":        2,
	"_TIMEOUT":    3,
}

func (x SESSION_CLOSE_CODE) String() string {
	return proto.EnumName(SESSION_CLOSE_CODE_name, int32(x))
}

func (SESSION_CLOSE_CODE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{3}
}

type REQUEST_STATUS int32

const (
	REQUEST_STATUS__ACCEPTED                        REQUEST_STATUS = 0
	REQUEST_STATUS__REJECTED                        REQUEST_STATUS = 1
	REQUEST_STATUS__REJECTED_SESSION_NONEXISTENT    REQUEST_STATUS = 2
	REQUEST_STATUS__REJECTED_SESSION_TABLE_FULL     REQUEST_STATUS = 3
	REQUEST_STATUS__REJECTED_SESSION_ALREADY_EXISTS REQUEST_STATUS = 4
)

var REQUEST_STATUS_name = map[int32]string{
	0: "_ACCEPTED",
	1: "_REJECTED",
	2: "_REJECTED_SESSION_NONEXISTENT",
	3: "_REJECTED_SESSION_TABLE_FULL",
	4: "_REJECTED_SESSION_ALREADY_EXISTS",
}

var REQUEST_STATUS_value = map[string]int32{
	"_ACCEPTED":                        0,
	"_REJECTED":                        1,
	"_REJECTED_SESSION_NONEXISTENT":    2,
	"_REJECTED_SESSION_TABLE_FULL":     3,
	"_REJECTED_SESSION_ALREADY_EXISTS": 4,
}

func (x REQUEST_STATUS) String() string {
	return proto.EnumName(REQUEST_STATUS_name, int32(x))
}

func (REQUEST_STATUS) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{4}
}

type ACTION_TYPE int32

const (
	ACTION_TYPE__DROP    ACTION_TYPE = 0
	ACTION_TYPE__FORWARD ACTION_TYPE = 1
	ACTION_TYPE__MIRROR  ACTION_TYPE = 2
	ACTION_TYPE__SNOOP   ACTION_TYPE = 3
)

var ACTION_TYPE_name = map[int32]string{
	0: "_DROP",
	1: "_FORWARD",
	2: "_MIRROR",
	3: "_SNOOP",
}

var ACTION_TYPE_value = map[string]int32{
	"_DROP":    0,
	"_FORWARD": 1,
	"_MIRROR":  2,
	"_SNOOP":   3,
}

func (x ACTION_TYPE) String() string {
	return proto.EnumName(ACTION_TYPE_name, int32(x))
}

func (ACTION_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{5}
}

//////////////////////////////////////////////////////////////////////////////////
//
// Activation Service
//
type INTERFACE_TYPE int32

const (
	INTERFACE_TYPE__NONE     INTERFACE_TYPE = 0
	INTERFACE_TYPE__SOFTWARE INTERFACE_TYPE = 1
	INTERFACE_TYPE__SMARTNIC INTERFACE_TYPE = 2
	INTERFACE_TYPE__NOS      INTERFACE_TYPE = 3
)

var INTERFACE_TYPE_name = map[int32]string{
	0: "_NONE",
	1: "_SOFTWARE",
	2: "_SMARTNIC",
	3: "_NOS",
}

var INTERFACE_TYPE_value = map[string]int32{
	"_NONE":     0,
	"_SOFTWARE": 1,
	"_SMARTNIC": 2,
	"_NOS":      3,
}

func (x INTERFACE_TYPE) String() string {
	return proto.EnumName(INTERFACE_TYPE_name, int32(x))
}

func (INTERFACE_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{6}
}

type ACTIVATION_STATUS_TYPE int32

const (
	ACTIVATION_STATUS_TYPE__DEVICE_ACTIVATED   ACTIVATION_STATUS_TYPE = 0
	ACTIVATION_STATUS_TYPE__DEVICE_DEACTIVATED ACTIVATION_STATUS_TYPE = 1
)

var ACTIVATION_STATUS_TYPE_name = map[int32]string{
	0: "_DEVICE_ACTIVATED",
	1: "_DEVICE_DEACTIVATED",
}

var ACTIVATION_STATUS_TYPE_value = map[string]int32{
	"_DEVICE_ACTIVATED":   0,
	"_DEVICE_DEACTIVATED": 1,
}

func (x ACTIVATION_STATUS_TYPE) String() string {
	return proto.EnumName(ACTIVATION_STATUS_TYPE_name, int32(x))
}

func (ACTIVATION_STATUS_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{7}
}

type REGISTRATION_STATUS_TYPE int32

const (
	REGISTRATION_STATUS_TYPE__DEVICE_REGISTERED   REGISTRATION_STATUS_TYPE = 0
	REGISTRATION_STATUS_TYPE__DEVICE_DEREGISTERED REGISTRATION_STATUS_TYPE = 1
)

var REGISTRATION_STATUS_TYPE_name = map[int32]string{
	0: "_DEVICE_REGISTERED",
	1: "_DEVICE_DEREGISTERED",
}

var REGISTRATION_STATUS_TYPE_value = map[string]int32{
	"_DEVICE_REGISTERED":   0,
	"_DEVICE_DEREGISTERED": 1,
}

func (x REGISTRATION_STATUS_TYPE) String() string {
	return proto.EnumName(REGISTRATION_STATUS_TYPE_name, int32(x))
}

func (REGISTRATION_STATUS_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{8}
}

// should the Application assign the sessionID on AddSession and avoid conflicts
// or have the applications have a mechanism to avoid duplicate sessionIDs across
// applications since there will be many applications instances to 1 switch
type SessionId struct {
	SessionId            int64    `protobuf:"varint,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionId) Reset()         { *m = SessionId{} }
func (m *SessionId) String() string { return proto.CompactTextString(m) }
func (*SessionId) ProtoMessage()    {}
func (*SessionId) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{0}
}

func (m *SessionId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionId.Unmarshal(m, b)
}
func (m *SessionId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionId.Marshal(b, m, deterministic)
}
func (m *SessionId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionId.Merge(m, src)
}
func (m *SessionId) XXX_Size() int {
	return xxx_messageInfo_SessionId.Size(m)
}
func (m *SessionId) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionId.DiscardUnknown(m)
}

var xxx_messageInfo_SessionId proto.InternalMessageInfo

func (m *SessionId) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

// MIRROR and SNOOP require an actionNextHop
// DROP and FORWARD do not have an actionNextHop
type ActionType struct {
	ActionType           ACTION_TYPE `protobuf:"varint,1,opt,name=actionType,proto3,enum=openoffload.v1alpha2.ACTION_TYPE" json:"actionType,omitempty"`
	ActionNextHop        string      `protobuf:"bytes,2,opt,name=actionNextHop,proto3" json:"actionNextHop,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ActionType) Reset()         { *m = ActionType{} }
func (m *ActionType) String() string { return proto.CompactTextString(m) }
func (*ActionType) ProtoMessage()    {}
func (*ActionType) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{1}
}

func (m *ActionType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionType.Unmarshal(m, b)
}
func (m *ActionType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionType.Marshal(b, m, deterministic)
}
func (m *ActionType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionType.Merge(m, src)
}
func (m *ActionType) XXX_Size() int {
	return xxx_messageInfo_ActionType.Size(m)
}
func (m *ActionType) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionType.DiscardUnknown(m)
}

var xxx_messageInfo_ActionType proto.InternalMessageInfo

func (m *ActionType) GetActionType() ACTION_TYPE {
	if m != nil {
		return m.ActionType
	}
	return ACTION_TYPE__DROP
}

func (m *ActionType) GetActionNextHop() string {
	if m != nil {
		return m.ActionNextHop
	}
	return ""
}

// sessionId is returned by server side upon successful addSession
type SessionRequest struct {
	InLif                uint32      `protobuf:"varint,1,opt,name=inLif,proto3" json:"inLif,omitempty"`
	OutLif               uint32      `protobuf:"varint,2,opt,name=outLif,proto3" json:"outLif,omitempty"`
	IpVersion            IP_VERSION  `protobuf:"varint,3,opt,name=ipVersion,proto3,enum=openoffload.v1alpha2.IP_VERSION" json:"ipVersion,omitempty"`
	SourceIp             []byte      `protobuf:"bytes,4,opt,name=sourceIp,proto3" json:"sourceIp,omitempty"`
	SourcePort           uint32      `protobuf:"varint,5,opt,name=sourcePort,proto3" json:"sourcePort,omitempty"`
	DestinationIp        []byte      `protobuf:"bytes,6,opt,name=destinationIp,proto3" json:"destinationIp,omitempty"`
	DestinationPort      uint32      `protobuf:"varint,7,opt,name=destinationPort,proto3" json:"destinationPort,omitempty"`
	ProtocolId           PROTOCOL_ID `protobuf:"varint,8,opt,name=protocolId,proto3,enum=openoffload.v1alpha2.PROTOCOL_ID" json:"protocolId,omitempty"`
	Action               *ActionType `protobuf:"bytes,9,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SessionRequest) Reset()         { *m = SessionRequest{} }
func (m *SessionRequest) String() string { return proto.CompactTextString(m) }
func (*SessionRequest) ProtoMessage()    {}
func (*SessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{2}
}

func (m *SessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionRequest.Unmarshal(m, b)
}
func (m *SessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionRequest.Marshal(b, m, deterministic)
}
func (m *SessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionRequest.Merge(m, src)
}
func (m *SessionRequest) XXX_Size() int {
	return xxx_messageInfo_SessionRequest.Size(m)
}
func (m *SessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SessionRequest proto.InternalMessageInfo

func (m *SessionRequest) GetInLif() uint32 {
	if m != nil {
		return m.InLif
	}
	return 0
}

func (m *SessionRequest) GetOutLif() uint32 {
	if m != nil {
		return m.OutLif
	}
	return 0
}

func (m *SessionRequest) GetIpVersion() IP_VERSION {
	if m != nil {
		return m.IpVersion
	}
	return IP_VERSION__IPV4
}

func (m *SessionRequest) GetSourceIp() []byte {
	if m != nil {
		return m.SourceIp
	}
	return nil
}

func (m *SessionRequest) GetSourcePort() uint32 {
	if m != nil {
		return m.SourcePort
	}
	return 0
}

func (m *SessionRequest) GetDestinationIp() []byte {
	if m != nil {
		return m.DestinationIp
	}
	return nil
}

func (m *SessionRequest) GetDestinationPort() uint32 {
	if m != nil {
		return m.DestinationPort
	}
	return 0
}

func (m *SessionRequest) GetProtocolId() PROTOCOL_ID {
	if m != nil {
		return m.ProtocolId
	}
	return PROTOCOL_ID__HOPOPT
}

func (m *SessionRequest) GetAction() *ActionType {
	if m != nil {
		return m.Action
	}
	return nil
}

type AddSessionResponse struct {
	SessionId            int64                `protobuf:"varint,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	RequestStatus        REQUEST_STATUS       `protobuf:"varint,2,opt,name=requestStatus,proto3,enum=openoffload.v1alpha2.REQUEST_STATUS" json:"requestStatus,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=startTime,proto3" json:"startTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AddSessionResponse) Reset()         { *m = AddSessionResponse{} }
func (m *AddSessionResponse) String() string { return proto.CompactTextString(m) }
func (*AddSessionResponse) ProtoMessage()    {}
func (*AddSessionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{3}
}

func (m *AddSessionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSessionResponse.Unmarshal(m, b)
}
func (m *AddSessionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSessionResponse.Marshal(b, m, deterministic)
}
func (m *AddSessionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSessionResponse.Merge(m, src)
}
func (m *AddSessionResponse) XXX_Size() int {
	return xxx_messageInfo_AddSessionResponse.Size(m)
}
func (m *AddSessionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSessionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddSessionResponse proto.InternalMessageInfo

func (m *AddSessionResponse) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *AddSessionResponse) GetRequestStatus() REQUEST_STATUS {
	if m != nil {
		return m.RequestStatus
	}
	return REQUEST_STATUS__ACCEPTED
}

func (m *AddSessionResponse) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

type SessionResponse struct {
	SessionId            int64                `protobuf:"varint,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	InPackets            int64                `protobuf:"varint,2,opt,name=inPackets,proto3" json:"inPackets,omitempty"`
	OutPackets           int64                `protobuf:"varint,3,opt,name=outPackets,proto3" json:"outPackets,omitempty"`
	InBytes              int64                `protobuf:"varint,4,opt,name=inBytes,proto3" json:"inBytes,omitempty"`
	OutBytes             int64                `protobuf:"varint,5,opt,name=outBytes,proto3" json:"outBytes,omitempty"`
	SessionState         SESSION_STATE        `protobuf:"varint,6,opt,name=sessionState,proto3,enum=openoffload.v1alpha2.SESSION_STATE" json:"sessionState,omitempty"`
	SessionCloseCode     SESSION_CLOSE_CODE   `protobuf:"varint,7,opt,name=sessionCloseCode,proto3,enum=openoffload.v1alpha2.SESSION_CLOSE_CODE" json:"sessionCloseCode,omitempty"`
	RequestStatus        REQUEST_STATUS       `protobuf:"varint,8,opt,name=requestStatus,proto3,enum=openoffload.v1alpha2.REQUEST_STATUS" json:"requestStatus,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,10,opt,name=endTime,proto3" json:"endTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SessionResponse) Reset()         { *m = SessionResponse{} }
func (m *SessionResponse) String() string { return proto.CompactTextString(m) }
func (*SessionResponse) ProtoMessage()    {}
func (*SessionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{4}
}

func (m *SessionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionResponse.Unmarshal(m, b)
}
func (m *SessionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionResponse.Marshal(b, m, deterministic)
}
func (m *SessionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionResponse.Merge(m, src)
}
func (m *SessionResponse) XXX_Size() int {
	return xxx_messageInfo_SessionResponse.Size(m)
}
func (m *SessionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SessionResponse proto.InternalMessageInfo

func (m *SessionResponse) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *SessionResponse) GetInPackets() int64 {
	if m != nil {
		return m.InPackets
	}
	return 0
}

func (m *SessionResponse) GetOutPackets() int64 {
	if m != nil {
		return m.OutPackets
	}
	return 0
}

func (m *SessionResponse) GetInBytes() int64 {
	if m != nil {
		return m.InBytes
	}
	return 0
}

func (m *SessionResponse) GetOutBytes() int64 {
	if m != nil {
		return m.OutBytes
	}
	return 0
}

func (m *SessionResponse) GetSessionState() SESSION_STATE {
	if m != nil {
		return m.SessionState
	}
	return SESSION_STATE__ESTABLISHED
}

func (m *SessionResponse) GetSessionCloseCode() SESSION_CLOSE_CODE {
	if m != nil {
		return m.SessionCloseCode
	}
	return SESSION_CLOSE_CODE__NOT_CLOSED
}

func (m *SessionResponse) GetRequestStatus() REQUEST_STATUS {
	if m != nil {
		return m.RequestStatus
	}
	return REQUEST_STATUS__ACCEPTED
}

func (m *SessionResponse) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *SessionResponse) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type StatisticsRequestArgs struct {
	//  pageSize = 0 will turn off paging
	//  does paging make sense for a stream ?
	//  the client should read/process each event on the stream anyway.
	PageSize             int32    `protobuf:"varint,1,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatisticsRequestArgs) Reset()         { *m = StatisticsRequestArgs{} }
func (m *StatisticsRequestArgs) String() string { return proto.CompactTextString(m) }
func (*StatisticsRequestArgs) ProtoMessage()    {}
func (*StatisticsRequestArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{5}
}

func (m *StatisticsRequestArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatisticsRequestArgs.Unmarshal(m, b)
}
func (m *StatisticsRequestArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatisticsRequestArgs.Marshal(b, m, deterministic)
}
func (m *StatisticsRequestArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatisticsRequestArgs.Merge(m, src)
}
func (m *StatisticsRequestArgs) XXX_Size() int {
	return xxx_messageInfo_StatisticsRequestArgs.Size(m)
}
func (m *StatisticsRequestArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_StatisticsRequestArgs.DiscardUnknown(m)
}

var xxx_messageInfo_StatisticsRequestArgs proto.InternalMessageInfo

func (m *StatisticsRequestArgs) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

//
// Empty message simulating void
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{6}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

//
//
type DeviceDescription struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 INTERFACE_TYPE `protobuf:"varint,2,opt,name=type,proto3,enum=openoffload.v1alpha2.INTERFACE_TYPE" json:"type,omitempty"`
	Description          string         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	SessionCapacity      int32          `protobuf:"varint,4,opt,name=sessionCapacity,proto3" json:"sessionCapacity,omitempty"`
	SessionRate          int32          `protobuf:"varint,5,opt,name=sessionRate,proto3" json:"sessionRate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeviceDescription) Reset()         { *m = DeviceDescription{} }
func (m *DeviceDescription) String() string { return proto.CompactTextString(m) }
func (*DeviceDescription) ProtoMessage()    {}
func (*DeviceDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{7}
}

func (m *DeviceDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceDescription.Unmarshal(m, b)
}
func (m *DeviceDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceDescription.Marshal(b, m, deterministic)
}
func (m *DeviceDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceDescription.Merge(m, src)
}
func (m *DeviceDescription) XXX_Size() int {
	return xxx_messageInfo_DeviceDescription.Size(m)
}
func (m *DeviceDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceDescription.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceDescription proto.InternalMessageInfo

func (m *DeviceDescription) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceDescription) GetType() INTERFACE_TYPE {
	if m != nil {
		return m.Type
	}
	return INTERFACE_TYPE__NONE
}

func (m *DeviceDescription) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DeviceDescription) GetSessionCapacity() int32 {
	if m != nil {
		return m.SessionCapacity
	}
	return 0
}

func (m *DeviceDescription) GetSessionRate() int32 {
	if m != nil {
		return m.SessionRate
	}
	return 0
}

//
// List of Offload devices available to be activated
//
type DeviceList struct {
	Devices              []*DeviceDescription `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DeviceList) Reset()         { *m = DeviceList{} }
func (m *DeviceList) String() string { return proto.CompactTextString(m) }
func (*DeviceList) ProtoMessage()    {}
func (*DeviceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{8}
}

func (m *DeviceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceList.Unmarshal(m, b)
}
func (m *DeviceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceList.Marshal(b, m, deterministic)
}
func (m *DeviceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceList.Merge(m, src)
}
func (m *DeviceList) XXX_Size() int {
	return xxx_messageInfo_DeviceList.Size(m)
}
func (m *DeviceList) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceList.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceList proto.InternalMessageInfo

func (m *DeviceList) GetDevices() []*DeviceDescription {
	if m != nil {
		return m.Devices
	}
	return nil
}

//
// Status of a registration request
//
type RegistrationStatus struct {
	Status               REGISTRATION_STATUS_TYPE `protobuf:"varint,1,opt,name=status,proto3,enum=openoffload.v1alpha2.REGISTRATION_STATUS_TYPE" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *RegistrationStatus) Reset()         { *m = RegistrationStatus{} }
func (m *RegistrationStatus) String() string { return proto.CompactTextString(m) }
func (*RegistrationStatus) ProtoMessage()    {}
func (*RegistrationStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{9}
}

func (m *RegistrationStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationStatus.Unmarshal(m, b)
}
func (m *RegistrationStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationStatus.Marshal(b, m, deterministic)
}
func (m *RegistrationStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationStatus.Merge(m, src)
}
func (m *RegistrationStatus) XXX_Size() int {
	return xxx_messageInfo_RegistrationStatus.Size(m)
}
func (m *RegistrationStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationStatus proto.InternalMessageInfo

func (m *RegistrationStatus) GetStatus() REGISTRATION_STATUS_TYPE {
	if m != nil {
		return m.Status
	}
	return REGISTRATION_STATUS_TYPE__DEVICE_REGISTERED
}

//
// Status of registration de-registration
// may want to expand to include reason for failures
//
type ActivationStatus struct {
	Status               ACTIVATION_STATUS_TYPE `protobuf:"varint,1,opt,name=status,proto3,enum=openoffload.v1alpha2.ACTIVATION_STATUS_TYPE" json:"status,omitempty"`
	Device               *DeviceDescription     `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ActivationStatus) Reset()         { *m = ActivationStatus{} }
func (m *ActivationStatus) String() string { return proto.CompactTextString(m) }
func (*ActivationStatus) ProtoMessage()    {}
func (*ActivationStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c682865d40e0bb3, []int{10}
}

func (m *ActivationStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivationStatus.Unmarshal(m, b)
}
func (m *ActivationStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivationStatus.Marshal(b, m, deterministic)
}
func (m *ActivationStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivationStatus.Merge(m, src)
}
func (m *ActivationStatus) XXX_Size() int {
	return xxx_messageInfo_ActivationStatus.Size(m)
}
func (m *ActivationStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivationStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ActivationStatus proto.InternalMessageInfo

func (m *ActivationStatus) GetStatus() ACTIVATION_STATUS_TYPE {
	if m != nil {
		return m.Status
	}
	return ACTIVATION_STATUS_TYPE__DEVICE_ACTIVATED
}

func (m *ActivationStatus) GetDevice() *DeviceDescription {
	if m != nil {
		return m.Device
	}
	return nil
}

func init() {
	proto.RegisterEnum("openoffload.v1alpha2.IP_VERSION", IP_VERSION_name, IP_VERSION_value)
	proto.RegisterEnum("openoffload.v1alpha2.PROTOCOL_ID", PROTOCOL_ID_name, PROTOCOL_ID_value)
	proto.RegisterEnum("openoffload.v1alpha2.SESSION_STATE", SESSION_STATE_name, SESSION_STATE_value)
	proto.RegisterEnum("openoffload.v1alpha2.SESSION_CLOSE_CODE", SESSION_CLOSE_CODE_name, SESSION_CLOSE_CODE_value)
	proto.RegisterEnum("openoffload.v1alpha2.REQUEST_STATUS", REQUEST_STATUS_name, REQUEST_STATUS_value)
	proto.RegisterEnum("openoffload.v1alpha2.ACTION_TYPE", ACTION_TYPE_name, ACTION_TYPE_value)
	proto.RegisterEnum("openoffload.v1alpha2.INTERFACE_TYPE", INTERFACE_TYPE_name, INTERFACE_TYPE_value)
	proto.RegisterEnum("openoffload.v1alpha2.ACTIVATION_STATUS_TYPE", ACTIVATION_STATUS_TYPE_name, ACTIVATION_STATUS_TYPE_value)
	proto.RegisterEnum("openoffload.v1alpha2.REGISTRATION_STATUS_TYPE", REGISTRATION_STATUS_TYPE_name, REGISTRATION_STATUS_TYPE_value)
	proto.RegisterType((*SessionId)(nil), "openoffload.v1alpha2.sessionId")
	proto.RegisterType((*ActionType)(nil), "openoffload.v1alpha2.actionType")
	proto.RegisterType((*SessionRequest)(nil), "openoffload.v1alpha2.sessionRequest")
	proto.RegisterType((*AddSessionResponse)(nil), "openoffload.v1alpha2.addSessionResponse")
	proto.RegisterType((*SessionResponse)(nil), "openoffload.v1alpha2.sessionResponse")
	proto.RegisterType((*StatisticsRequestArgs)(nil), "openoffload.v1alpha2.statisticsRequestArgs")
	proto.RegisterType((*Empty)(nil), "openoffload.v1alpha2.Empty")
	proto.RegisterType((*DeviceDescription)(nil), "openoffload.v1alpha2.deviceDescription")
	proto.RegisterType((*DeviceList)(nil), "openoffload.v1alpha2.deviceList")
	proto.RegisterType((*RegistrationStatus)(nil), "openoffload.v1alpha2.registrationStatus")
	proto.RegisterType((*ActivationStatus)(nil), "openoffload.v1alpha2.activationStatus")
}

func init() { proto.RegisterFile("openoffload.proto", fileDescriptor_0c682865d40e0bb3) }

var fileDescriptor_0c682865d40e0bb3 = []byte{
	// 1364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xef, 0x72, 0xdb, 0x44,
	0x10, 0xb7, 0xfc, 0xdf, 0xeb, 0xc4, 0x51, 0xae, 0x6d, 0xea, 0x49, 0x0b, 0x75, 0x45, 0x01, 0x63,
	0x18, 0x97, 0xba, 0x1d, 0xa6, 0x33, 0xcc, 0xd0, 0x51, 0x65, 0xa5, 0x51, 0x71, 0x25, 0x73, 0x52,
	0x52, 0xca, 0x40, 0x35, 0xaa, 0x75, 0x71, 0x35, 0x75, 0x2c, 0x61, 0x9d, 0x33, 0x84, 0x17, 0xe0,
	0x19, 0xfa, 0x91, 0xb7, 0xe0, 0x2d, 0x78, 0x14, 0x3e, 0xf2, 0x0d, 0x98, 0x3b, 0x49, 0xb1, 0x5c,
	0xff, 0x69, 0x32, 0x53, 0xfa, 0xed, 0x76, 0x6f, 0xf7, 0xb7, 0x7f, 0x6f, 0xf7, 0x60, 0xdb, 0x0f,
	0xc8, 0xd8, 0x3f, 0x3a, 0x1a, 0xf9, 0x8e, 0xdb, 0x0e, 0x26, 0x3e, 0xf5, 0xd1, 0xe5, 0x34, 0xeb,
	0xe4, 0x8e, 0x33, 0x0a, 0x5e, 0x3a, 0x9d, 0xdd, 0x1b, 0x43, 0xdf, 0x1f, 0x8e, 0xc8, 0x6d, 0x2e,
	0xf3, 0x62, 0x7a, 0x74, 0x9b, 0x7a, 0xc7, 0x24, 0xa4, 0xce, 0x71, 0x10, 0xa9, 0x49, 0x9f, 0x41,
	0x25, 0x24, 0x61, 0xe8, 0xf9, 0x63, 0xcd, 0x45, 0xd7, 0x53, 0x44, 0x5d, 0x68, 0x08, 0xcd, 0x1c,
	0x9e, 0x31, 0xa4, 0x29, 0x80, 0x33, 0xa0, 0x9e, 0x3f, 0xb6, 0x4e, 0x03, 0x82, 0xe4, 0x34, 0xc5,
	0x85, 0x6b, 0x9d, 0x9b, 0xed, 0x65, 0x4e, 0xb4, 0x65, 0xc5, 0xd2, 0x0c, 0xdd, 0xb6, 0x9e, 0xf5,
	0x55, 0x9c, 0x86, 0xb8, 0x05, 0x9b, 0x11, 0xa5, 0x93, 0x5f, 0xe8, 0xbe, 0x1f, 0xd4, 0xb3, 0x0d,
	0xa1, 0x59, 0xc1, 0xf3, 0x4c, 0xe9, 0xb7, 0x1c, 0xd4, 0x62, 0x27, 0x30, 0xf9, 0x79, 0x4a, 0x42,
	0x8a, 0x2e, 0x43, 0xc1, 0x1b, 0xf7, 0xbc, 0x23, 0x6e, 0x76, 0x13, 0x47, 0x04, 0xda, 0x81, 0xa2,
	0x3f, 0xa5, 0x8c, 0x9d, 0xe5, 0xec, 0x98, 0x42, 0xdf, 0x40, 0xc5, 0x0b, 0x0e, 0xc9, 0x84, 0x21,
	0xd4, 0x73, 0xdc, 0xd1, 0xc6, 0x72, 0x47, 0xb5, 0xbe, 0x7d, 0xa8, 0x62, 0x53, 0x33, 0x74, 0x3c,
	0x53, 0x41, 0xbb, 0x50, 0x0e, 0xfd, 0xe9, 0x64, 0x40, 0xb4, 0xa0, 0x9e, 0x6f, 0x08, 0xcd, 0x0d,
	0x7c, 0x46, 0xa3, 0x0f, 0x01, 0xa2, 0x73, 0xdf, 0x9f, 0xd0, 0x7a, 0x81, 0xdb, 0x4d, 0x71, 0x58,
	0x88, 0x2e, 0x09, 0xa9, 0x37, 0x76, 0x58, 0x48, 0x5a, 0x50, 0x2f, 0x72, 0x80, 0x79, 0x26, 0x6a,
	0xc2, 0x56, 0x8a, 0xc1, 0xa1, 0x4a, 0x1c, 0xea, 0x4d, 0x36, 0xcb, 0x3a, 0xaf, 0xdb, 0xc0, 0x1f,
	0x69, 0x6e, 0xbd, 0xbc, 0x2e, 0xeb, 0x7d, 0x6c, 0x58, 0x86, 0x62, 0xf4, 0x6c, 0xad, 0x8b, 0x53,
	0x4a, 0xe8, 0x3e, 0x14, 0xa3, 0x04, 0xd7, 0x2b, 0x0d, 0xa1, 0x59, 0x5d, 0x95, 0x8b, 0x59, 0x9d,
	0x70, 0x2c, 0x2f, 0xfd, 0x21, 0x00, 0x72, 0x5c, 0xd7, 0x4c, 0x8a, 0x11, 0x06, 0xfe, 0x38, 0x24,
	0xeb, 0xbb, 0x06, 0x3d, 0x86, 0xcd, 0x49, 0x54, 0x36, 0x93, 0x3a, 0x74, 0x1a, 0xf2, 0xe2, 0xd4,
	0x3a, 0xb7, 0x96, 0x5b, 0xc5, 0xea, 0x77, 0x07, 0xaa, 0x69, 0xd9, 0xa6, 0x25, 0x5b, 0x07, 0x26,
	0x9e, 0x57, 0x45, 0xf7, 0xa1, 0x12, 0x52, 0x67, 0x42, 0x2d, 0xef, 0x98, 0xf0, 0x4a, 0x56, 0x3b,
	0xbb, 0xed, 0xa8, 0xc3, 0xdb, 0x49, 0x87, 0xb7, 0xad, 0xa4, 0xc3, 0xf1, 0x4c, 0x58, 0xfa, 0x27,
	0x07, 0x5b, 0xe1, 0x85, 0xfc, 0xbe, 0x0e, 0x15, 0x6f, 0xdc, 0x77, 0x06, 0xaf, 0x08, 0x8d, 0x7c,
	0xce, 0xe1, 0x19, 0x83, 0xd5, 0xdd, 0x9f, 0xd2, 0xe4, 0x3a, 0xc7, 0xaf, 0x53, 0x1c, 0x54, 0x87,
	0x92, 0x37, 0x7e, 0x78, 0x4a, 0x49, 0xc8, 0x5b, 0x26, 0x87, 0x13, 0x92, 0x75, 0x93, 0x3f, 0xa5,
	0xd1, 0x55, 0x81, 0x5f, 0x9d, 0xd1, 0xe8, 0x11, 0x6c, 0xc4, 0x0e, 0xb0, 0x80, 0x09, 0x6f, 0x96,
	0x5a, 0xe7, 0xa3, 0xe5, 0xa9, 0x32, 0x55, 0x93, 0x75, 0x2a, 0x4f, 0x95, 0x8a, 0xe7, 0x14, 0x91,
	0x05, 0x62, 0x4c, 0x2b, 0x23, 0x3f, 0x24, 0x8a, 0xef, 0x12, 0xde, 0x51, 0xb5, 0x4e, 0x73, 0x3d,
	0x98, 0xd2, 0x33, 0x4c, 0xd5, 0x56, 0x8c, 0xae, 0x8a, 0x17, 0x10, 0x16, 0x4b, 0x59, 0x7e, 0x47,
	0xa5, 0xac, 0x5c, 0xa0, 0x94, 0xe8, 0x1e, 0x94, 0xc8, 0xd8, 0xe5, 0x7a, 0xf0, 0x56, 0xbd, 0x44,
	0x54, 0xba, 0x0b, 0x57, 0x42, 0xea, 0x50, 0x2f, 0xa4, 0xde, 0x20, 0x8c, 0xe7, 0x88, 0x3c, 0x19,
	0xf2, 0x7a, 0x04, 0xce, 0x90, 0x98, 0xde, 0xaf, 0xd1, 0x14, 0x2b, 0xe0, 0x33, 0x5a, 0x2a, 0x41,
	0x41, 0x3d, 0x0e, 0xe8, 0xa9, 0xf4, 0xa7, 0x00, 0xdb, 0x2e, 0x39, 0xf1, 0x06, 0xa4, 0x4b, 0xc2,
	0xc1, 0xc4, 0x0b, 0xd8, 0x7b, 0x40, 0x08, 0xf2, 0x63, 0xe7, 0x38, 0x52, 0xab, 0x60, 0x7e, 0x46,
	0xf7, 0x21, 0x4f, 0xd9, 0x40, 0x5c, 0xdb, 0xe5, 0x9a, 0x6e, 0xa9, 0x78, 0x4f, 0x56, 0xd4, 0x68,
	0x26, 0x72, 0x0d, 0xd4, 0x80, 0xaa, 0x3b, 0x03, 0xe7, 0x3d, 0x55, 0xc1, 0x69, 0x16, 0x1b, 0x13,
	0x49, 0x4d, 0x9c, 0xc0, 0x19, 0x78, 0xf4, 0x94, 0x37, 0x57, 0x01, 0xbf, 0xc9, 0x66, 0x58, 0x49,
	0xb7, 0xb3, 0x3e, 0x2a, 0x70, 0xa9, 0x34, 0x4b, 0x32, 0x00, 0xa2, 0x80, 0x7a, 0x5e, 0xc8, 0xc6,
	0x4a, 0x29, 0xa2, 0xc2, 0xba, 0xd0, 0xc8, 0x35, 0xab, 0x9d, 0x4f, 0x97, 0x3b, 0xbe, 0x90, 0x03,
	0x9c, 0xe8, 0x49, 0x3f, 0x02, 0x9a, 0x90, 0xa1, 0x17, 0xd2, 0x09, 0x9f, 0x56, 0x71, 0x99, 0xf7,
	0xa0, 0x18, 0x46, 0xbd, 0x12, 0x6d, 0x88, 0xf6, 0xaa, 0x5e, 0x79, 0xa4, 0x99, 0x16, 0x96, 0xad,
	0xa4, 0xa1, 0x0f, 0xcc, 0x28, 0x35, 0xb1, 0xb6, 0xf4, 0x5a, 0x00, 0x91, 0x4d, 0xa1, 0x93, 0x34,
	0x78, 0xf7, 0x0d, 0xf0, 0x2f, 0x56, 0xaf, 0x9f, 0xc3, 0x95, 0xd0, 0xe8, 0x01, 0x14, 0xa3, 0x18,
	0x78, 0xcd, 0x2e, 0x10, 0x7a, 0xac, 0xd6, 0x92, 0x00, 0x66, 0x8b, 0x03, 0x55, 0xa0, 0x60, 0x6b,
	0xfd, 0xc3, 0x7b, 0x62, 0x26, 0x39, 0x7e, 0x25, 0x0a, 0x2d, 0x19, 0xaa, 0xa9, 0x79, 0x8c, 0xaa,
	0x50, 0xb2, 0xf7, 0x8d, 0xbe, 0xd1, 0xb7, 0xc4, 0x0c, 0x2a, 0x43, 0xde, 0xb6, 0x94, 0xbe, 0x58,
	0xe4, 0xa7, 0x83, 0x6e, 0x5f, 0xdc, 0x96, 0xf2, 0x65, 0x41, 0x2c, 0x48, 0xf9, 0x72, 0x49, 0x14,
	0xa5, 0x42, 0x19, 0x89, 0xff, 0x0a, 0x2d, 0x1d, 0x36, 0xe7, 0x9e, 0x3c, 0x12, 0x61, 0xc3, 0x56,
	0x4d, 0x4b, 0x7e, 0xd8, 0xd3, 0xcc, 0x7d, 0xb5, 0x2b, 0x66, 0x50, 0x0d, 0x80, 0xbf, 0x60, 0x4d,
	0x7f, 0x64, 0xdf, 0x11, 0x85, 0x39, 0xba, 0x23, 0x66, 0xb9, 0x59, 0xfe, 0xc2, 0xbb, 0x62, 0xae,
	0xf5, 0x18, 0xd0, 0xe2, 0xab, 0x47, 0x5b, 0x50, 0xb5, 0x75, 0xc3, 0x4a, 0xc4, 0x32, 0x5c, 0x67,
	0x4f, 0xd3, 0x65, 0xe5, 0x5b, 0x51, 0xe0, 0x0e, 0x62, 0xd3, 0x12, 0xb3, 0x68, 0x03, 0xca, 0xb6,
	0xa5, 0x3d, 0x51, 0x8d, 0x03, 0x4b, 0xcc, 0xb5, 0x5e, 0x0b, 0x50, 0x9b, 0x7f, 0xef, 0x68, 0x13,
	0x2a, 0xb6, 0xac, 0x28, 0x6a, 0xdf, 0xe2, 0x30, 0x8c, 0xc4, 0xea, 0x63, 0x55, 0x61, 0xa4, 0x80,
	0x6e, 0xc2, 0x07, 0x67, 0xa4, 0x9d, 0xb8, 0xa1, 0x1b, 0xba, 0xfa, 0xbd, 0x66, 0x5a, 0xaa, 0xce,
	0x2c, 0x34, 0xe0, 0xfa, 0xa2, 0x08, 0x8b, 0x56, 0xb5, 0xf7, 0x0e, 0x7a, 0x3d, 0x31, 0x87, 0x6e,
	0x41, 0x63, 0x51, 0x42, 0xee, 0x61, 0x55, 0xee, 0x3e, 0xb3, 0x39, 0x92, 0x29, 0xe6, 0x5b, 0x0f,
	0xa0, 0x9a, 0xfa, 0x80, 0xf0, 0xa2, 0x74, 0xb1, 0xd1, 0x17, 0x33, 0x3c, 0x86, 0x3d, 0x03, 0x3f,
	0x95, 0x31, 0x73, 0x89, 0x05, 0xfa, 0x44, 0xc3, 0xd8, 0xc0, 0x62, 0x16, 0x01, 0x14, 0x6d, 0x53,
	0x37, 0x8c, 0xbe, 0x98, 0x6b, 0x29, 0x50, 0x9b, 0x7f, 0xb0, 0x1c, 0x83, 0x39, 0x1b, 0xc7, 0x65,
	0x1a, 0x7b, 0xd6, 0x53, 0x19, 0xab, 0xa2, 0x10, 0x91, 0x4f, 0x64, 0x6c, 0xe9, 0x9a, 0x22, 0x66,
	0x79, 0xbe, 0x74, 0xc3, 0x14, 0x73, 0xad, 0x7d, 0xd8, 0x59, 0xde, 0x87, 0xe8, 0x0a, 0x6c, 0xdb,
	0x5d, 0xf5, 0x50, 0x53, 0x54, 0x3b, 0x96, 0xe0, 0x09, 0xbb, 0x0a, 0x97, 0x12, 0x76, 0x57, 0x9d,
	0x5d, 0x08, 0xad, 0x1e, 0xd4, 0x57, 0x3d, 0x17, 0xb4, 0x03, 0x28, 0x51, 0x8a, 0x64, 0x54, 0xcc,
	0xc1, 0xea, 0x70, 0x79, 0x06, 0x96, 0xba, 0x11, 0x3a, 0xbf, 0x67, 0x61, 0x23, 0x5e, 0xe8, 0x96,
	0xf3, 0x62, 0x44, 0xd0, 0x73, 0x80, 0xd9, 0x8e, 0x47, 0x2b, 0x06, 0xd8, 0xfc, 0x7f, 0x6c, 0x77,
	0xc5, 0x52, 0x59, 0xfc, 0x2b, 0x48, 0x19, 0x74, 0x08, 0x30, 0x24, 0x34, 0xc1, 0xbf, 0xb1, 0x16,
	0x5f, 0x73, 0x77, 0x3f, 0x7e, 0x8b, 0x03, 0x67, 0xb8, 0xcf, 0xd8, 0x4f, 0x6b, 0x44, 0x28, 0x79,
	0xe7, 0xd0, 0x9d, 0xbf, 0x04, 0xd8, 0x31, 0x67, 0xeb, 0x35, 0xda, 0x21, 0x51, 0xb6, 0x5e, 0x42,
	0x6d, 0x48, 0xa8, 0x3c, 0x1a, 0xc5, 0xf7, 0x21, 0xfa, 0x7c, 0x05, 0xea, 0xb2, 0xe5, 0x73, 0x6e,
	0x17, 0xbe, 0x14, 0xd0, 0x2b, 0xd8, 0x1e, 0x12, 0xca, 0x97, 0xb1, 0xfb, 0x7f, 0x1b, 0xeb, 0xfc,
	0x9d, 0x07, 0x90, 0xcf, 0xc6, 0x2d, 0x1a, 0xc1, 0x95, 0x68, 0xb6, 0x93, 0x89, 0x11, 0xa9, 0x77,
	0xf9, 0xe8, 0x43, 0xe7, 0x9d, 0x95, 0xab, 0x3a, 0x64, 0x71, 0x63, 0x48, 0x19, 0x34, 0x86, 0xab,
	0x2e, 0x79, 0x8f, 0xf6, 0x9e, 0xc3, 0xb5, 0x21, 0xa1, 0x38, 0x36, 0x48, 0xdc, 0x39, 0x93, 0x21,
	0xba, 0xb6, 0x1c, 0x8a, 0x7f, 0x0c, 0x76, 0x1b, 0xeb, 0x1c, 0x62, 0xab, 0x55, 0xca, 0x20, 0x17,
	0xb6, 0xe2, 0xd5, 0x45, 0x62, 0xe8, 0xf3, 0xc7, 0xf1, 0xc9, 0xea, 0xcf, 0xf9, 0xc9, 0x7c, 0x14,
	0x47, 0xec, 0x87, 0xf2, 0x1e, 0xec, 0xfc, 0x04, 0x97, 0x58, 0xc7, 0xc7, 0x86, 0x92, 0x64, 0xad,
	0xcf, 0xd2, 0x79, 0xdd, 0x90, 0x32, 0x0f, 0xeb, 0x3f, 0xec, 0xb4, 0xbf, 0x4e, 0x49, 0xdb, 0x89,
	0xf4, 0x8b, 0x22, 0xff, 0xde, 0xdd, 0xfd, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x69, 0x84, 0xf1, 0x4b,
	0xfc, 0x0e, 0x00, 0x00,
}