// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sync_service.proto

package mtproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 同步类型
type SyncType int32

const (
	SyncType_SYNC_TYPE_UNKNOWN    SyncType = 0
	SyncType_SYNC_TYPE_USER       SyncType = 1
	SyncType_SYNC_TYPE_USER_NOTME SyncType = 2
	SyncType_SYNC_TYPE_USER_ME    SyncType = 3
	SyncType_SYNC_TYPE_RPC_RESULT SyncType = 4
	SyncType_SYNC_TYPE_NONE       SyncType = 5
)

var SyncType_name = map[int32]string{
	0: "SYNC_TYPE_UNKNOWN",
	1: "SYNC_TYPE_USER",
	2: "SYNC_TYPE_USER_NOTME",
	3: "SYNC_TYPE_USER_ME",
	4: "SYNC_TYPE_RPC_RESULT",
	5: "SYNC_TYPE_NONE",
}
var SyncType_value = map[string]int32{
	"SYNC_TYPE_UNKNOWN":    0,
	"SYNC_TYPE_USER":       1,
	"SYNC_TYPE_USER_NOTME": 2,
	"SYNC_TYPE_USER_ME":    3,
	"SYNC_TYPE_RPC_RESULT": 4,
	"SYNC_TYPE_NONE":       5,
}

func (x SyncType) String() string {
	return proto.EnumName(SyncType_name, int32(x))
}
func (SyncType) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

type VoidRsp struct {
}

func (m *VoidRsp) Reset()                    { *m = VoidRsp{} }
func (m *VoidRsp) String() string            { return proto.CompactTextString(m) }
func (*VoidRsp) ProtoMessage()               {}
func (*VoidRsp) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

// PushMessage state
type ClientUpdatesState struct {
	Pts      int32 `protobuf:"varint,1,opt,name=pts" json:"pts,omitempty"`
	PtsCount int32 `protobuf:"varint,2,opt,name=pts_count,json=ptsCount" json:"pts_count,omitempty"`
	Qts      int32 `protobuf:"varint,3,opt,name=qts" json:"qts,omitempty"`
	QtsCount int32 `protobuf:"varint,4,opt,name=qts_count,json=qtsCount" json:"qts_count,omitempty"`
	Seq      int32 `protobuf:"varint,5,opt,name=seq" json:"seq,omitempty"`
	SeqStart int32 `protobuf:"varint,6,opt,name=seq_start,json=seqStart" json:"seq_start,omitempty"`
	Date     int32 `protobuf:"varint,7,opt,name=date" json:"date,omitempty"`
}

func (m *ClientUpdatesState) Reset()                    { *m = ClientUpdatesState{} }
func (m *ClientUpdatesState) String() string            { return proto.CompactTextString(m) }
func (*ClientUpdatesState) ProtoMessage()               {}
func (*ClientUpdatesState) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *ClientUpdatesState) GetPts() int32 {
	if m != nil {
		return m.Pts
	}
	return 0
}

func (m *ClientUpdatesState) GetPtsCount() int32 {
	if m != nil {
		return m.PtsCount
	}
	return 0
}

func (m *ClientUpdatesState) GetQts() int32 {
	if m != nil {
		return m.Qts
	}
	return 0
}

func (m *ClientUpdatesState) GetQtsCount() int32 {
	if m != nil {
		return m.QtsCount
	}
	return 0
}

func (m *ClientUpdatesState) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *ClientUpdatesState) GetSeqStart() int32 {
	if m != nil {
		return m.SeqStart
	}
	return 0
}

func (m *ClientUpdatesState) GetDate() int32 {
	if m != nil {
		return m.Date
	}
	return 0
}

// /////////////////////////////////////////////////////////////////////
// SERVER_AUTH_REQ
type ConnectToSessionServerReq struct {
}

func (m *ConnectToSessionServerReq) Reset()                    { *m = ConnectToSessionServerReq{} }
func (m *ConnectToSessionServerReq) String() string            { return proto.CompactTextString(m) }
func (*ConnectToSessionServerReq) ProtoMessage()               {}
func (*ConnectToSessionServerReq) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

type SessionServerConnectedRsp struct {
	ServerId   int32  `protobuf:"varint,1,opt,name=server_id,json=serverId" json:"server_id,omitempty"`
	ServerName string `protobuf:"bytes,2,opt,name=server_name,json=serverName" json:"server_name,omitempty"`
}

func (m *SessionServerConnectedRsp) Reset()                    { *m = SessionServerConnectedRsp{} }
func (m *SessionServerConnectedRsp) String() string            { return proto.CompactTextString(m) }
func (*SessionServerConnectedRsp) ProtoMessage()               {}
func (*SessionServerConnectedRsp) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *SessionServerConnectedRsp) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *SessionServerConnectedRsp) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

// PushUpdatesData --> VoidRsp
type PushUpdatesData struct {
	AuthKeyId   int64               `protobuf:"varint,1,opt,name=auth_key_id,json=authKeyId" json:"auth_key_id,omitempty"`
	SessionId   int64               `protobuf:"varint,2,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	State       *ClientUpdatesState `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
	UpdatesData []byte              `protobuf:"bytes,5,opt,name=updates_data,json=updatesData,proto3" json:"updates_data,omitempty"`
}

func (m *PushUpdatesData) Reset()                    { *m = PushUpdatesData{} }
func (m *PushUpdatesData) String() string            { return proto.CompactTextString(m) }
func (*PushUpdatesData) ProtoMessage()               {}
func (*PushUpdatesData) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{4} }

func (m *PushUpdatesData) GetAuthKeyId() int64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

func (m *PushUpdatesData) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *PushUpdatesData) GetState() *ClientUpdatesState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *PushUpdatesData) GetUpdatesData() []byte {
	if m != nil {
		return m.UpdatesData
	}
	return nil
}

// Updates
// messages_affectedHistory
// messages_affectedMessages
type RpcResultData struct {
	// int32 rpc_result_type = 1;
	Updates          *Updates                    `protobuf:"bytes,2,opt,name=updates" json:"updates,omitempty"`
	AffectedHistory  *TLMessagesAffectedHistory  `protobuf:"bytes,3,opt,name=affected_history,json=affectedHistory" json:"affected_history,omitempty"`
	AffectedMessages *TLMessagesAffectedMessages `protobuf:"bytes,4,opt,name=affected_messages,json=affectedMessages" json:"affected_messages,omitempty"`
}

func (m *RpcResultData) Reset()                    { *m = RpcResultData{} }
func (m *RpcResultData) String() string            { return proto.CompactTextString(m) }
func (*RpcResultData) ProtoMessage()               {}
func (*RpcResultData) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{5} }

func (m *RpcResultData) GetUpdates() *Updates {
	if m != nil {
		return m.Updates
	}
	return nil
}

func (m *RpcResultData) GetAffectedHistory() *TLMessagesAffectedHistory {
	if m != nil {
		return m.AffectedHistory
	}
	return nil
}

func (m *RpcResultData) GetAffectedMessages() *TLMessagesAffectedMessages {
	if m != nil {
		return m.AffectedMessages
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////
// RPC
type UpdatesRequest struct {
	PushType    SyncType       `protobuf:"varint,1,opt,name=push_type,json=pushType,enum=mtproto.SyncType" json:"push_type,omitempty"`
	Layer       int32          `protobuf:"varint,2,opt,name=layer" json:"layer,omitempty"`
	ServerId    int32          `protobuf:"varint,3,opt,name=server_id,json=serverId" json:"server_id,omitempty"`
	AuthKeyId   int64          `protobuf:"varint,4,opt,name=auth_key_id,json=authKeyId" json:"auth_key_id,omitempty"`
	SessionId   int64          `protobuf:"varint,5,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	PushUserId  int32          `protobuf:"varint,6,opt,name=push_user_id,json=pushUserId" json:"push_user_id,omitempty"`
	ClientMsgId int64          `protobuf:"varint,7,opt,name=client_msg_id,json=clientMsgId" json:"client_msg_id,omitempty"`
	Updates     *Updates       `protobuf:"bytes,8,opt,name=updates" json:"updates,omitempty"`
	RpcResult   *RpcResultData `protobuf:"bytes,9,opt,name=rpc_result,json=rpcResult" json:"rpc_result,omitempty"`
}

func (m *UpdatesRequest) Reset()                    { *m = UpdatesRequest{} }
func (m *UpdatesRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdatesRequest) ProtoMessage()               {}
func (*UpdatesRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{6} }

func (m *UpdatesRequest) GetPushType() SyncType {
	if m != nil {
		return m.PushType
	}
	return SyncType_SYNC_TYPE_UNKNOWN
}

func (m *UpdatesRequest) GetLayer() int32 {
	if m != nil {
		return m.Layer
	}
	return 0
}

func (m *UpdatesRequest) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *UpdatesRequest) GetAuthKeyId() int64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

func (m *UpdatesRequest) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *UpdatesRequest) GetPushUserId() int32 {
	if m != nil {
		return m.PushUserId
	}
	return 0
}

func (m *UpdatesRequest) GetClientMsgId() int64 {
	if m != nil {
		return m.ClientMsgId
	}
	return 0
}

func (m *UpdatesRequest) GetUpdates() *Updates {
	if m != nil {
		return m.Updates
	}
	return nil
}

func (m *UpdatesRequest) GetRpcResult() *RpcResultData {
	if m != nil {
		return m.RpcResult
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////
type NewUpdatesRequest struct {
	AuthKeyId int64 `protobuf:"varint,1,opt,name=auth_key_id,json=authKeyId" json:"auth_key_id,omitempty"`
	UserId    int32 `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *NewUpdatesRequest) Reset()                    { *m = NewUpdatesRequest{} }
func (m *NewUpdatesRequest) String() string            { return proto.CompactTextString(m) }
func (*NewUpdatesRequest) ProtoMessage()               {}
func (*NewUpdatesRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{7} }

func (m *NewUpdatesRequest) GetAuthKeyId() int64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

func (m *NewUpdatesRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

// GetServerUpdatesState
// func (m *UpdateModel) GetUpdateListByGtPts(userId, pts int32) []*mtproto.Update {
type UserGtPtsUpdatesRequest struct {
	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Pts    int32 `protobuf:"varint,2,opt,name=pts" json:"pts,omitempty"`
}

func (m *UserGtPtsUpdatesRequest) Reset()                    { *m = UserGtPtsUpdatesRequest{} }
func (m *UserGtPtsUpdatesRequest) String() string            { return proto.CompactTextString(m) }
func (*UserGtPtsUpdatesRequest) ProtoMessage()               {}
func (*UserGtPtsUpdatesRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{8} }

func (m *UserGtPtsUpdatesRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserGtPtsUpdatesRequest) GetPts() int32 {
	if m != nil {
		return m.Pts
	}
	return 0
}

type ChannelGtPtsUpdatesRequest struct {
	ChannelId int32 `protobuf:"varint,1,opt,name=channel_id,json=channelId" json:"channel_id,omitempty"`
	Pts       int32 `protobuf:"varint,2,opt,name=pts" json:"pts,omitempty"`
}

func (m *ChannelGtPtsUpdatesRequest) Reset()                    { *m = ChannelGtPtsUpdatesRequest{} }
func (m *ChannelGtPtsUpdatesRequest) String() string            { return proto.CompactTextString(m) }
func (*ChannelGtPtsUpdatesRequest) ProtoMessage()               {}
func (*ChannelGtPtsUpdatesRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{9} }

func (m *ChannelGtPtsUpdatesRequest) GetChannelId() int32 {
	if m != nil {
		return m.ChannelId
	}
	return 0
}

func (m *ChannelGtPtsUpdatesRequest) GetPts() int32 {
	if m != nil {
		return m.Pts
	}
	return 0
}

type ChannelPtsRequest struct {
	ChannelId int32 `protobuf:"varint,1,opt,name=channel_id,json=channelId" json:"channel_id,omitempty"`
}

func (m *ChannelPtsRequest) Reset()                    { *m = ChannelPtsRequest{} }
func (m *ChannelPtsRequest) String() string            { return proto.CompactTextString(m) }
func (*ChannelPtsRequest) ProtoMessage()               {}
func (*ChannelPtsRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{10} }

func (m *ChannelPtsRequest) GetChannelId() int32 {
	if m != nil {
		return m.ChannelId
	}
	return 0
}

type SeqId struct {
	Pts int32 `protobuf:"varint,1,opt,name=pts" json:"pts,omitempty"`
}

func (m *SeqId) Reset()                    { *m = SeqId{} }
func (m *SeqId) String() string            { return proto.CompactTextString(m) }
func (*SeqId) ProtoMessage()               {}
func (*SeqId) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{11} }

func (m *SeqId) GetPts() int32 {
	if m != nil {
		return m.Pts
	}
	return 0
}

type UpdatesStateRequest struct {
	AuthKeyId int64 `protobuf:"varint,1,opt,name=auth_key_id,json=authKeyId" json:"auth_key_id,omitempty"`
	UserId    int32 `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Pts       int32 `protobuf:"varint,3,opt,name=pts" json:"pts,omitempty"`
	Qts       int32 `protobuf:"varint,4,opt,name=qts" json:"qts,omitempty"`
}

func (m *UpdatesStateRequest) Reset()                    { *m = UpdatesStateRequest{} }
func (m *UpdatesStateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdatesStateRequest) ProtoMessage()               {}
func (*UpdatesStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{12} }

func (m *UpdatesStateRequest) GetAuthKeyId() int64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

func (m *UpdatesStateRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UpdatesStateRequest) GetPts() int32 {
	if m != nil {
		return m.Pts
	}
	return 0
}

func (m *UpdatesStateRequest) GetQts() int32 {
	if m != nil {
		return m.Qts
	}
	return 0
}

func init() {
	proto.RegisterType((*VoidRsp)(nil), "mtproto.VoidRsp")
	proto.RegisterType((*ClientUpdatesState)(nil), "mtproto.ClientUpdatesState")
	proto.RegisterType((*ConnectToSessionServerReq)(nil), "mtproto.ConnectToSessionServerReq")
	proto.RegisterType((*SessionServerConnectedRsp)(nil), "mtproto.SessionServerConnectedRsp")
	proto.RegisterType((*PushUpdatesData)(nil), "mtproto.PushUpdatesData")
	proto.RegisterType((*RpcResultData)(nil), "mtproto.RpcResultData")
	proto.RegisterType((*UpdatesRequest)(nil), "mtproto.UpdatesRequest")
	proto.RegisterType((*NewUpdatesRequest)(nil), "mtproto.NewUpdatesRequest")
	proto.RegisterType((*UserGtPtsUpdatesRequest)(nil), "mtproto.UserGtPtsUpdatesRequest")
	proto.RegisterType((*ChannelGtPtsUpdatesRequest)(nil), "mtproto.ChannelGtPtsUpdatesRequest")
	proto.RegisterType((*ChannelPtsRequest)(nil), "mtproto.ChannelPtsRequest")
	proto.RegisterType((*SeqId)(nil), "mtproto.SeqId")
	proto.RegisterType((*UpdatesStateRequest)(nil), "mtproto.UpdatesStateRequest")
	proto.RegisterEnum("mtproto.SyncType", SyncType_name, SyncType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RPCSync service

type RPCSyncClient interface {
	SyncUpdatesData(ctx context.Context, in *UpdatesRequest, opts ...grpc.CallOption) (*ClientUpdatesState, error)
	PushUpdatesData(ctx context.Context, in *UpdatesRequest, opts ...grpc.CallOption) (*VoidRsp, error)
	// PushUpdatesDataList(UpdatesListRequest) returns (VoidRsp);
	GetNewUpdatesData(ctx context.Context, in *NewUpdatesRequest, opts ...grpc.CallOption) (*Updates, error)
	GetCurrentChannelPts(ctx context.Context, in *ChannelPtsRequest, opts ...grpc.CallOption) (*SeqId, error)
	GetUserGtPtsUpdatesData(ctx context.Context, in *UserGtPtsUpdatesRequest, opts ...grpc.CallOption) (*Updates, error)
	GetChannelGtPtsUpdatesData(ctx context.Context, in *ChannelGtPtsUpdatesRequest, opts ...grpc.CallOption) (*Updates, error)
	GetServerUpdatesState(ctx context.Context, in *UpdatesStateRequest, opts ...grpc.CallOption) (*Updates_State, error)
	UpdateUpdatesState(ctx context.Context, in *UpdatesStateRequest, opts ...grpc.CallOption) (*VoidRsp, error)
}

type rPCSyncClient struct {
	cc *grpc.ClientConn
}

func NewRPCSyncClient(cc *grpc.ClientConn) RPCSyncClient {
	return &rPCSyncClient{cc}
}

func (c *rPCSyncClient) SyncUpdatesData(ctx context.Context, in *UpdatesRequest, opts ...grpc.CallOption) (*ClientUpdatesState, error) {
	out := new(ClientUpdatesState)
	err := grpc.Invoke(ctx, "/mtproto.RPCSync/SyncUpdatesData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) PushUpdatesData(ctx context.Context, in *UpdatesRequest, opts ...grpc.CallOption) (*VoidRsp, error) {
	out := new(VoidRsp)
	err := grpc.Invoke(ctx, "/mtproto.RPCSync/PushUpdatesData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) GetNewUpdatesData(ctx context.Context, in *NewUpdatesRequest, opts ...grpc.CallOption) (*Updates, error) {
	out := new(Updates)
	err := grpc.Invoke(ctx, "/mtproto.RPCSync/GetNewUpdatesData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) GetCurrentChannelPts(ctx context.Context, in *ChannelPtsRequest, opts ...grpc.CallOption) (*SeqId, error) {
	out := new(SeqId)
	err := grpc.Invoke(ctx, "/mtproto.RPCSync/GetCurrentChannelPts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) GetUserGtPtsUpdatesData(ctx context.Context, in *UserGtPtsUpdatesRequest, opts ...grpc.CallOption) (*Updates, error) {
	out := new(Updates)
	err := grpc.Invoke(ctx, "/mtproto.RPCSync/GetUserGtPtsUpdatesData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) GetChannelGtPtsUpdatesData(ctx context.Context, in *ChannelGtPtsUpdatesRequest, opts ...grpc.CallOption) (*Updates, error) {
	out := new(Updates)
	err := grpc.Invoke(ctx, "/mtproto.RPCSync/GetChannelGtPtsUpdatesData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) GetServerUpdatesState(ctx context.Context, in *UpdatesStateRequest, opts ...grpc.CallOption) (*Updates_State, error) {
	out := new(Updates_State)
	err := grpc.Invoke(ctx, "/mtproto.RPCSync/GetServerUpdatesState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) UpdateUpdatesState(ctx context.Context, in *UpdatesStateRequest, opts ...grpc.CallOption) (*VoidRsp, error) {
	out := new(VoidRsp)
	err := grpc.Invoke(ctx, "/mtproto.RPCSync/UpdateUpdatesState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RPCSync service

type RPCSyncServer interface {
	SyncUpdatesData(context.Context, *UpdatesRequest) (*ClientUpdatesState, error)
	PushUpdatesData(context.Context, *UpdatesRequest) (*VoidRsp, error)
	// PushUpdatesDataList(UpdatesListRequest) returns (VoidRsp);
	GetNewUpdatesData(context.Context, *NewUpdatesRequest) (*Updates, error)
	GetCurrentChannelPts(context.Context, *ChannelPtsRequest) (*SeqId, error)
	GetUserGtPtsUpdatesData(context.Context, *UserGtPtsUpdatesRequest) (*Updates, error)
	GetChannelGtPtsUpdatesData(context.Context, *ChannelGtPtsUpdatesRequest) (*Updates, error)
	GetServerUpdatesState(context.Context, *UpdatesStateRequest) (*Updates_State, error)
	UpdateUpdatesState(context.Context, *UpdatesStateRequest) (*VoidRsp, error)
}

func RegisterRPCSyncServer(s *grpc.Server, srv RPCSyncServer) {
	s.RegisterService(&_RPCSync_serviceDesc, srv)
}

func _RPCSync_SyncUpdatesData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).SyncUpdatesData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCSync/SyncUpdatesData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).SyncUpdatesData(ctx, req.(*UpdatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_PushUpdatesData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).PushUpdatesData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCSync/PushUpdatesData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).PushUpdatesData(ctx, req.(*UpdatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_GetNewUpdatesData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewUpdatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).GetNewUpdatesData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCSync/GetNewUpdatesData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).GetNewUpdatesData(ctx, req.(*NewUpdatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_GetCurrentChannelPts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelPtsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).GetCurrentChannelPts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCSync/GetCurrentChannelPts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).GetCurrentChannelPts(ctx, req.(*ChannelPtsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_GetUserGtPtsUpdatesData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGtPtsUpdatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).GetUserGtPtsUpdatesData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCSync/GetUserGtPtsUpdatesData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).GetUserGtPtsUpdatesData(ctx, req.(*UserGtPtsUpdatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_GetChannelGtPtsUpdatesData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelGtPtsUpdatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).GetChannelGtPtsUpdatesData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCSync/GetChannelGtPtsUpdatesData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).GetChannelGtPtsUpdatesData(ctx, req.(*ChannelGtPtsUpdatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_GetServerUpdatesState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatesStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).GetServerUpdatesState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCSync/GetServerUpdatesState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).GetServerUpdatesState(ctx, req.(*UpdatesStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_UpdateUpdatesState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatesStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).UpdateUpdatesState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCSync/UpdateUpdatesState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).UpdateUpdatesState(ctx, req.(*UpdatesStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCSync_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mtproto.RPCSync",
	HandlerType: (*RPCSyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncUpdatesData",
			Handler:    _RPCSync_SyncUpdatesData_Handler,
		},
		{
			MethodName: "PushUpdatesData",
			Handler:    _RPCSync_PushUpdatesData_Handler,
		},
		{
			MethodName: "GetNewUpdatesData",
			Handler:    _RPCSync_GetNewUpdatesData_Handler,
		},
		{
			MethodName: "GetCurrentChannelPts",
			Handler:    _RPCSync_GetCurrentChannelPts_Handler,
		},
		{
			MethodName: "GetUserGtPtsUpdatesData",
			Handler:    _RPCSync_GetUserGtPtsUpdatesData_Handler,
		},
		{
			MethodName: "GetChannelGtPtsUpdatesData",
			Handler:    _RPCSync_GetChannelGtPtsUpdatesData_Handler,
		},
		{
			MethodName: "GetServerUpdatesState",
			Handler:    _RPCSync_GetServerUpdatesState_Handler,
		},
		{
			MethodName: "UpdateUpdatesState",
			Handler:    _RPCSync_UpdateUpdatesState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sync_service.proto",
}

func init() { proto.RegisterFile("sync_service.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 924 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0xf9, 0x69, 0x92, 0x93, 0x6e, 0xeb, 0xcc, 0x76, 0x69, 0x9a, 0xb2, 0x50, 0x0c, 0x2b,
	0xa1, 0xbd, 0x88, 0x44, 0x10, 0x77, 0xdc, 0xd0, 0x34, 0xca, 0x56, 0xdb, 0xb8, 0xd1, 0x38, 0x05,
	0x95, 0x1b, 0xcb, 0xeb, 0x9c, 0x6d, 0x22, 0x12, 0xff, 0xcd, 0x78, 0x51, 0x78, 0x0e, 0x9e, 0x81,
	0x37, 0xe0, 0x3d, 0xb8, 0xe7, 0x86, 0x47, 0x41, 0x33, 0xe3, 0x9f, 0xc4, 0x0e, 0xdd, 0x45, 0x5c,
	0x79, 0xe6, 0x9b, 0x6f, 0xbe, 0x39, 0x73, 0xce, 0x37, 0xc7, 0x40, 0xd8, 0xc6, 0x73, 0x6d, 0x86,
	0xd1, 0xbb, 0xa5, 0x8b, 0xfd, 0x20, 0xf2, 0xb9, 0x4f, 0x1a, 0x6b, 0x2e, 0x07, 0xbd, 0x13, 0xe6,
	0x2e, 0x70, 0xed, 0xf4, 0xf9, 0xaa, 0x2f, 0x68, 0x6a, 0xd9, 0x68, 0x41, 0xe3, 0x07, 0x7f, 0x39,
	0xa7, 0x2c, 0x30, 0xfe, 0xd0, 0x80, 0x0c, 0x57, 0x4b, 0xf4, 0xf8, 0x5d, 0x30, 0x77, 0x38, 0x32,
	0x8b, 0x3b, 0x1c, 0x89, 0x0e, 0xd5, 0x80, 0xb3, 0xae, 0x76, 0xa1, 0x7d, 0x55, 0xa7, 0x62, 0x48,
	0xce, 0xa1, 0x15, 0x70, 0x66, 0xbb, 0x7e, 0xec, 0xf1, 0x6e, 0x45, 0xe2, 0xcd, 0x80, 0xb3, 0xa1,
	0x98, 0x0b, 0x7a, 0xc8, 0x59, 0xb7, 0xaa, 0xe8, 0xa1, 0xa2, 0x87, 0x19, 0xbd, 0xa6, 0xe8, 0xe1,
	0x16, 0x9d, 0x61, 0xd8, 0xad, 0x2b, 0x3a, 0xc3, 0x50, 0xd0, 0x19, 0x86, 0x36, 0xe3, 0x4e, 0xc4,
	0xbb, 0x07, 0x8a, 0xce, 0x30, 0xb4, 0xc4, 0x9c, 0x10, 0xa8, 0x89, 0xd0, 0xba, 0x0d, 0x89, 0xcb,
	0xb1, 0x71, 0x0e, 0x67, 0x43, 0xdf, 0xf3, 0xd0, 0xe5, 0x33, 0xdf, 0x42, 0xc6, 0x96, 0xbe, 0x67,
	0x61, 0xf4, 0x0e, 0x23, 0x8a, 0xa1, 0x71, 0x0f, 0x67, 0x3b, 0x58, 0xc2, 0x44, 0x71, 0x63, 0x75,
	0x94, 0x40, 0xed, 0xe5, 0x3c, 0xb9, 0x60, 0x53, 0x01, 0xd7, 0x73, 0xf2, 0x19, 0xb4, 0x93, 0x45,
	0xcf, 0x59, 0xa3, 0xbc, 0x67, 0x8b, 0x82, 0x82, 0x4c, 0x67, 0x8d, 0xc6, 0xef, 0x1a, 0x1c, 0x4f,
	0x63, 0xb6, 0x48, 0xb2, 0x75, 0xe5, 0x70, 0x87, 0x7c, 0x0a, 0x6d, 0x27, 0xe6, 0x0b, 0xfb, 0x67,
	0xdc, 0xa4, 0x9a, 0x55, 0xda, 0x12, 0xd0, 0x6b, 0xdc, 0x5c, 0xcf, 0xc9, 0x73, 0x00, 0xa6, 0xc2,
	0x11, 0xcb, 0x15, 0xb5, 0x9c, 0x20, 0xd7, 0x73, 0xf2, 0x35, 0xd4, 0x99, 0x48, 0xba, 0x4c, 0x5f,
	0x7b, 0x70, 0xde, 0x4f, 0x8a, 0xd7, 0x2f, 0xd7, 0x85, 0x2a, 0x26, 0xf9, 0x1c, 0x0e, 0x63, 0x05,
	0xdb, 0x73, 0x87, 0x3b, 0x32, 0x93, 0x87, 0xb4, 0x1d, 0xe7, 0x41, 0x19, 0x7f, 0x6b, 0xf0, 0x84,
	0x06, 0x2e, 0x45, 0x16, 0xaf, 0xb8, 0x0c, 0xf3, 0x25, 0x34, 0x12, 0x82, 0x8c, 0xa1, 0x3d, 0xd0,
	0xb3, 0x93, 0x92, 0x33, 0x68, 0x4a, 0x20, 0xb7, 0xa0, 0x3b, 0x6f, 0xdf, 0xca, 0x9c, 0xd9, 0x8b,
	0x25, 0xe3, 0x7e, 0xb4, 0x49, 0xc2, 0xfb, 0x32, 0xdb, 0x34, 0xbb, 0xb1, 0xd7, 0xc8, 0x98, 0xf3,
	0x80, 0xcc, 0x4e, 0xc9, 0xaf, 0x14, 0x97, 0x1e, 0x17, 0x00, 0x42, 0xa1, 0x93, 0x09, 0xa6, 0xbb,
	0xa4, 0x2f, 0xda, 0x83, 0x17, 0x8f, 0x2a, 0x4e, 0x12, 0x80, 0xea, 0x45, 0xc4, 0xf8, 0xab, 0x02,
	0x47, 0x69, 0xe4, 0x18, 0xc6, 0xc8, 0x38, 0xe9, 0x43, 0x2b, 0x88, 0xd9, 0xc2, 0xe6, 0x9b, 0x00,
	0x65, 0x21, 0x8e, 0x06, 0x9d, 0x4c, 0xde, 0xda, 0x78, 0xee, 0x6c, 0x13, 0x20, 0x6d, 0x0a, 0x8e,
	0x18, 0x91, 0x13, 0xa8, 0xaf, 0x9c, 0x0d, 0x46, 0x89, 0xa3, 0xd5, 0x64, 0xd7, 0x22, 0xd5, 0x82,
	0x45, 0x0a, 0xd5, 0xae, 0x3d, 0x5e, 0xed, 0x7a, 0xb1, 0xda, 0x17, 0x70, 0x28, 0x23, 0x8c, 0x99,
	0x92, 0x57, 0x66, 0x07, 0x81, 0xdd, 0x31, 0x79, 0x80, 0x01, 0x4f, 0x5c, 0x59, 0x79, 0x7b, 0xcd,
	0x1e, 0x04, 0xa5, 0x21, 0x35, 0xda, 0x0a, 0x9c, 0xb0, 0x87, 0xeb, 0xf9, 0x76, 0x2d, 0x9b, 0xef,
	0xab, 0xe5, 0xb7, 0x00, 0x51, 0xe0, 0xda, 0x91, 0x74, 0x42, 0xb7, 0x25, 0xe9, 0x1f, 0x67, 0xf4,
	0x1d, 0x8f, 0xd0, 0x56, 0x94, 0x4e, 0x8d, 0x1b, 0xe8, 0x98, 0xf8, 0x4b, 0x21, 0xbf, 0xef, 0xb3,
	0xfa, 0x29, 0x34, 0xd2, 0x8b, 0xa9, 0x8c, 0x1e, 0xc4, 0xf2, 0x52, 0xc6, 0x15, 0x9c, 0x8a, 0xeb,
	0x8d, 0xf9, 0x94, 0xb3, 0x82, 0xe6, 0xd6, 0x1e, 0x6d, 0x7b, 0x4f, 0xda, 0x84, 0x2a, 0x59, 0x13,
	0x32, 0x26, 0xd0, 0x1b, 0x2e, 0x1c, 0xcf, 0xc3, 0xd5, 0x3e, 0xa1, 0xe7, 0x00, 0xae, 0x5a, 0xcd,
	0xb5, 0x5a, 0x09, 0xb2, 0x57, 0x6e, 0x00, 0x9d, 0x44, 0x6e, 0xca, 0x3f, 0x50, 0xc5, 0x38, 0x83,
	0xba, 0x85, 0x61, 0x2e, 0x97, 0xb7, 0x48, 0x23, 0x82, 0xa7, 0x3b, 0x8f, 0xf5, 0x7f, 0xe6, 0x2c,
	0x3d, 0xa1, 0x9a, 0x37, 0xe1, 0xa4, 0xcf, 0xd6, 0xb2, 0x3e, 0xfb, 0xf2, 0x37, 0x0d, 0x9a, 0xa9,
	0xaf, 0xc9, 0x33, 0xe8, 0x58, 0xf7, 0xe6, 0xd0, 0x9e, 0xdd, 0x4f, 0x47, 0xf6, 0x9d, 0xf9, 0xda,
	0xbc, 0xfd, 0xd1, 0xd4, 0x3f, 0x22, 0x04, 0x8e, 0xb6, 0x60, 0x6b, 0x44, 0x75, 0x8d, 0x74, 0xe1,
	0x64, 0x17, 0xb3, 0xcd, 0xdb, 0xd9, 0x64, 0xa4, 0x57, 0x0a, 0x22, 0x62, 0x65, 0x32, 0xd2, 0xab,
	0xbb, 0x1b, 0xe8, 0x74, 0x68, 0xd3, 0x91, 0x75, 0x77, 0x33, 0xd3, 0x6b, 0xbb, 0xf2, 0xe6, 0xad,
	0x39, 0xd2, 0xeb, 0x83, 0x3f, 0x6b, 0xd0, 0xa0, 0xd3, 0xa1, 0x88, 0x8c, 0x8c, 0xe1, 0x58, 0x7c,
	0xb7, 0x3b, 0xe6, 0x69, 0xc9, 0xad, 0x2a, 0x57, 0xbd, 0xc7, 0x9a, 0x1f, 0xf9, 0xae, 0xdc, 0x7a,
	0xff, 0x55, 0x28, 0x7f, 0x0f, 0xc9, 0x9f, 0x8e, 0x7c, 0x0f, 0x9d, 0x31, 0xf2, 0xdc, 0xd2, 0x72,
	0x7f, 0x2f, 0xa3, 0x95, 0xbc, 0xde, 0x2b, 0x3d, 0x29, 0x72, 0x09, 0x27, 0x63, 0xe4, 0xc3, 0x38,
	0x8a, 0xd0, 0xe3, 0xb9, 0x73, 0xb6, 0x54, 0x4a, 0x76, 0xea, 0x1d, 0xe5, 0xed, 0x47, 0xda, 0x66,
	0x02, 0xa7, 0x63, 0xe4, 0xc5, 0xb7, 0x20, 0x83, 0xb9, 0xc8, 0x0f, 0xdc, 0xff, 0x54, 0xf6, 0x84,
	0x64, 0x41, 0x4f, 0x84, 0x54, 0x7e, 0x14, 0x52, 0xf1, 0x8b, 0x62, 0x60, 0x1f, 0x26, 0x3a, 0x81,
	0x67, 0x63, 0xe4, 0xea, 0xdf, 0xb9, 0x53, 0x81, 0x4f, 0x8a, 0xd4, 0x6d, 0xa3, 0xf7, 0xf2, 0xa6,
	0x92, 0xfe, 0x9c, 0xd4, 0xae, 0x2b, 0x20, 0x8a, 0xfe, 0x1f, 0xb4, 0x4a, 0xf5, 0xbb, 0x7c, 0x01,
	0x4f, 0x5d, 0x7f, 0xdd, 0xf7, 0xf0, 0x4d, 0xbc, 0x72, 0x96, 0xeb, 0xfe, 0xaf, 0x72, 0xf9, 0x12,
	0x7e, 0x9a, 0x8a, 0xaf, 0x70, 0xd8, 0xab, 0xca, 0x54, 0x7b, 0x73, 0x20, 0xe1, 0x6f, 0xfe, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x88, 0x1e, 0x66, 0x3d, 0x17, 0x09, 0x00, 0x00,
}
