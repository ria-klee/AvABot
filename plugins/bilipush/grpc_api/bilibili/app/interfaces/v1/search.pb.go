// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_api/bilibili/app/interfaces/v1/search.proto

package bilibili_app_interface_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

//
type DefaultWordsReply struct {
	//
	Trackid string `protobuf:"bytes,1,opt,name=trackid,proto3" json:"trackid,omitempty"`
	//
	Param string `protobuf:"bytes,2,opt,name=param,proto3" json:"param,omitempty"`
	//
	Show string `protobuf:"bytes,3,opt,name=show,proto3" json:"show,omitempty"`
	//
	Word string `protobuf:"bytes,4,opt,name=word,proto3" json:"word,omitempty"`
	//
	ShowFront int64 `protobuf:"varint,5,opt,name=show_front,json=showFront,proto3" json:"show_front,omitempty"`
	//
	ExpStr string `protobuf:"bytes,6,opt,name=exp_str,json=expStr,proto3" json:"exp_str,omitempty"`
	//
	Goto string `protobuf:"bytes,7,opt,name=goto,proto3" json:"goto,omitempty"`
	//
	Value string `protobuf:"bytes,8,opt,name=value,proto3" json:"value,omitempty"`
	//
	Uri                  string   `protobuf:"bytes,9,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DefaultWordsReply) Reset()         { *m = DefaultWordsReply{} }
func (m *DefaultWordsReply) String() string { return proto.CompactTextString(m) }
func (*DefaultWordsReply) ProtoMessage()    {}
func (*DefaultWordsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e899187b61530774, []int{0}
}

func (m *DefaultWordsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefaultWordsReply.Unmarshal(m, b)
}
func (m *DefaultWordsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefaultWordsReply.Marshal(b, m, deterministic)
}
func (m *DefaultWordsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultWordsReply.Merge(m, src)
}
func (m *DefaultWordsReply) XXX_Size() int {
	return xxx_messageInfo_DefaultWordsReply.Size(m)
}
func (m *DefaultWordsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultWordsReply.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultWordsReply proto.InternalMessageInfo

func (m *DefaultWordsReply) GetTrackid() string {
	if m != nil {
		return m.Trackid
	}
	return ""
}

func (m *DefaultWordsReply) GetParam() string {
	if m != nil {
		return m.Param
	}
	return ""
}

func (m *DefaultWordsReply) GetShow() string {
	if m != nil {
		return m.Show
	}
	return ""
}

func (m *DefaultWordsReply) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

func (m *DefaultWordsReply) GetShowFront() int64 {
	if m != nil {
		return m.ShowFront
	}
	return 0
}

func (m *DefaultWordsReply) GetExpStr() string {
	if m != nil {
		return m.ExpStr
	}
	return ""
}

func (m *DefaultWordsReply) GetGoto() string {
	if m != nil {
		return m.Goto
	}
	return ""
}

func (m *DefaultWordsReply) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *DefaultWordsReply) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

//
type NftFaceIcon struct {
	//
	RegionType int32 `protobuf:"varint,1,opt,name=region_type,json=regionType,proto3" json:"region_type,omitempty"`
	//
	Icon string `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon,omitempty"`
	//
	ShowStatus           int32    `protobuf:"varint,3,opt,name=show_status,json=showStatus,proto3" json:"show_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NftFaceIcon) Reset()         { *m = NftFaceIcon{} }
func (m *NftFaceIcon) String() string { return proto.CompactTextString(m) }
func (*NftFaceIcon) ProtoMessage()    {}
func (*NftFaceIcon) Descriptor() ([]byte, []int) {
	return fileDescriptor_e899187b61530774, []int{1}
}

func (m *NftFaceIcon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NftFaceIcon.Unmarshal(m, b)
}
func (m *NftFaceIcon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NftFaceIcon.Marshal(b, m, deterministic)
}
func (m *NftFaceIcon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NftFaceIcon.Merge(m, src)
}
func (m *NftFaceIcon) XXX_Size() int {
	return xxx_messageInfo_NftFaceIcon.Size(m)
}
func (m *NftFaceIcon) XXX_DiscardUnknown() {
	xxx_messageInfo_NftFaceIcon.DiscardUnknown(m)
}

var xxx_messageInfo_NftFaceIcon proto.InternalMessageInfo

func (m *NftFaceIcon) GetRegionType() int32 {
	if m != nil {
		return m.RegionType
	}
	return 0
}

func (m *NftFaceIcon) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *NftFaceIcon) GetShowStatus() int32 {
	if m != nil {
		return m.ShowStatus
	}
	return 0
}

//
type DefaultWordsReq struct {
	//
	From int64 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	//
	LoginEvent int64 `protobuf:"varint,2,opt,name=login_event,json=loginEvent,proto3" json:"login_event,omitempty"`
	//
	TeenagersMode int32 `protobuf:"varint,3,opt,name=teenagers_mode,json=teenagersMode,proto3" json:"teenagers_mode,omitempty"`
	//
	LessonsMode int32 `protobuf:"varint,4,opt,name=lessons_mode,json=lessonsMode,proto3" json:"lessons_mode,omitempty"`
	//
	Tab string `protobuf:"bytes,5,opt,name=tab,proto3" json:"tab,omitempty"`
	//
	EventId string `protobuf:"bytes,6,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	//
	Avid string `protobuf:"bytes,7,opt,name=avid,proto3" json:"avid,omitempty"`
	//
	Query string `protobuf:"bytes,8,opt,name=query,proto3" json:"query,omitempty"`
	//
	An int64 `protobuf:"varint,9,opt,name=an,proto3" json:"an,omitempty"`
	//
	IsFresh              int64    `protobuf:"varint,10,opt,name=is_fresh,json=isFresh,proto3" json:"is_fresh,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DefaultWordsReq) Reset()         { *m = DefaultWordsReq{} }
func (m *DefaultWordsReq) String() string { return proto.CompactTextString(m) }
func (*DefaultWordsReq) ProtoMessage()    {}
func (*DefaultWordsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e899187b61530774, []int{2}
}

func (m *DefaultWordsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefaultWordsReq.Unmarshal(m, b)
}
func (m *DefaultWordsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefaultWordsReq.Marshal(b, m, deterministic)
}
func (m *DefaultWordsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultWordsReq.Merge(m, src)
}
func (m *DefaultWordsReq) XXX_Size() int {
	return xxx_messageInfo_DefaultWordsReq.Size(m)
}
func (m *DefaultWordsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultWordsReq.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultWordsReq proto.InternalMessageInfo

func (m *DefaultWordsReq) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *DefaultWordsReq) GetLoginEvent() int64 {
	if m != nil {
		return m.LoginEvent
	}
	return 0
}

func (m *DefaultWordsReq) GetTeenagersMode() int32 {
	if m != nil {
		return m.TeenagersMode
	}
	return 0
}

func (m *DefaultWordsReq) GetLessonsMode() int32 {
	if m != nil {
		return m.LessonsMode
	}
	return 0
}

func (m *DefaultWordsReq) GetTab() string {
	if m != nil {
		return m.Tab
	}
	return ""
}

func (m *DefaultWordsReq) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *DefaultWordsReq) GetAvid() string {
	if m != nil {
		return m.Avid
	}
	return ""
}

func (m *DefaultWordsReq) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *DefaultWordsReq) GetAn() int64 {
	if m != nil {
		return m.An
	}
	return 0
}

func (m *DefaultWordsReq) GetIsFresh() int64 {
	if m != nil {
		return m.IsFresh
	}
	return 0
}

// 获取搜索建议-响应
type SuggestionResult3Reply struct {
	// 搜索追踪id
	Trackid string `protobuf:"bytes,1,opt,name=trackid,proto3" json:"trackid,omitempty"`
	// 搜索建议条目列表
	List []*ResultItem `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	// 搜索的abtest 实验信息
	ExpStr               string   `protobuf:"bytes,3,opt,name=exp_str,json=expStr,proto3" json:"exp_str,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuggestionResult3Reply) Reset()         { *m = SuggestionResult3Reply{} }
func (m *SuggestionResult3Reply) String() string { return proto.CompactTextString(m) }
func (*SuggestionResult3Reply) ProtoMessage()    {}
func (*SuggestionResult3Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e899187b61530774, []int{3}
}

func (m *SuggestionResult3Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuggestionResult3Reply.Unmarshal(m, b)
}
func (m *SuggestionResult3Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuggestionResult3Reply.Marshal(b, m, deterministic)
}
func (m *SuggestionResult3Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuggestionResult3Reply.Merge(m, src)
}
func (m *SuggestionResult3Reply) XXX_Size() int {
	return xxx_messageInfo_SuggestionResult3Reply.Size(m)
}
func (m *SuggestionResult3Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_SuggestionResult3Reply.DiscardUnknown(m)
}

var xxx_messageInfo_SuggestionResult3Reply proto.InternalMessageInfo

func (m *SuggestionResult3Reply) GetTrackid() string {
	if m != nil {
		return m.Trackid
	}
	return ""
}

func (m *SuggestionResult3Reply) GetList() []*ResultItem {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *SuggestionResult3Reply) GetExpStr() string {
	if m != nil {
		return m.ExpStr
	}
	return ""
}

// 获取搜索建议-请求
type SuggestionResult3Req struct {
	// 关键字
	Keyword string `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	// 是否语法高亮
	// 0:不显示 1:显示
	Highlight int32 `protobuf:"varint,2,opt,name=highlight,proto3" json:"highlight,omitempty"`
	// 是否青少年模式
	// 1:开启青少年模式
	TeenagersMode        int32    `protobuf:"varint,3,opt,name=teenagers_mode,json=teenagersMode,proto3" json:"teenagers_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuggestionResult3Req) Reset()         { *m = SuggestionResult3Req{} }
func (m *SuggestionResult3Req) String() string { return proto.CompactTextString(m) }
func (*SuggestionResult3Req) ProtoMessage()    {}
func (*SuggestionResult3Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_e899187b61530774, []int{4}
}

func (m *SuggestionResult3Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuggestionResult3Req.Unmarshal(m, b)
}
func (m *SuggestionResult3Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuggestionResult3Req.Marshal(b, m, deterministic)
}
func (m *SuggestionResult3Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuggestionResult3Req.Merge(m, src)
}
func (m *SuggestionResult3Req) XXX_Size() int {
	return xxx_messageInfo_SuggestionResult3Req.Size(m)
}
func (m *SuggestionResult3Req) XXX_DiscardUnknown() {
	xxx_messageInfo_SuggestionResult3Req.DiscardUnknown(m)
}

var xxx_messageInfo_SuggestionResult3Req proto.InternalMessageInfo

func (m *SuggestionResult3Req) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func (m *SuggestionResult3Req) GetHighlight() int32 {
	if m != nil {
		return m.Highlight
	}
	return 0
}

func (m *SuggestionResult3Req) GetTeenagersMode() int32 {
	if m != nil {
		return m.TeenagersMode
	}
	return 0
}

// 搜索建议条目
type ResultItem struct {
	// 来源
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// 显示结果(语法高亮)
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// 搜索关键字
	Keyword string `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
	// 序号
	Position int32 `protobuf:"varint,4,opt,name=position,proto3" json:"position,omitempty"`
	// 图片
	Cover string `protobuf:"bytes,5,opt,name=cover,proto3" json:"cover,omitempty"`
	// 图片尺寸
	CoverSize float64 `protobuf:"fixed64,6,opt,name=cover_size,json=coverSize,proto3" json:"cover_size,omitempty"`
	// sug词类型
	SugType string `protobuf:"bytes,7,opt,name=sug_type,json=sugType,proto3" json:"sug_type,omitempty"`
	// 词条大类型
	TermType int32 `protobuf:"varint,8,opt,name=term_type,json=termType,proto3" json:"term_type,omitempty"`
	// 跳转类型
	Goto string `protobuf:"bytes,9,opt,name=goto,proto3" json:"goto,omitempty"`
	// 跳转uri
	Uri string `protobuf:"bytes,10,opt,name=uri,proto3" json:"uri,omitempty"`
	// 认证信息
	OfficialVerify *OfficialVerify `protobuf:"bytes,11,opt,name=official_verify,json=officialVerify,proto3" json:"official_verify,omitempty"`
	// 跳转参数
	Param string `protobuf:"bytes,12,opt,name=param,proto3" json:"param,omitempty"`
	// up主mid
	Mid int64 `protobuf:"varint,13,opt,name=mid,proto3" json:"mid,omitempty"`
	// 粉丝数
	Fans int32 `protobuf:"varint,14,opt,name=fans,proto3" json:"fans,omitempty"`
	// up主等级
	Level int32 `protobuf:"varint,15,opt,name=level,proto3" json:"level,omitempty"`
	// up主稿件数
	Archives int32 `protobuf:"varint,16,opt,name=archives,proto3" json:"archives,omitempty"`
	// 投稿时间
	Ptime int64 `protobuf:"varint,17,opt,name=ptime,proto3" json:"ptime,omitempty"`
	// season类型名称
	SeasonTypeName string `protobuf:"bytes,18,opt,name=season_type_name,json=seasonTypeName,proto3" json:"season_type_name,omitempty"`
	// 地区
	Area string `protobuf:"bytes,19,opt,name=area,proto3" json:"area,omitempty"`
	// 作品风格
	Style string `protobuf:"bytes,20,opt,name=style,proto3" json:"style,omitempty"`
	// 描述信息
	Label string `protobuf:"bytes,21,opt,name=label,proto3" json:"label,omitempty"`
	// 评分
	Rating float64 `protobuf:"fixed64,22,opt,name=rating,proto3" json:"rating,omitempty"`
	// 投票数
	Vote int32 `protobuf:"varint,23,opt,name=vote,proto3" json:"vote,omitempty"`
	// 角标
	Badges []*ReasonStyle `protobuf:"bytes,24,rep,name=badges,proto3" json:"badges,omitempty"`
	//
	Styles string `protobuf:"bytes,25,opt,name=styles,proto3" json:"styles,omitempty"`
	//
	ModuleId int64 `protobuf:"varint,26,opt,name=module_id,json=moduleId,proto3" json:"module_id,omitempty"`
	//
	LiveLink string `protobuf:"bytes,27,opt,name=live_link,json=liveLink,proto3" json:"live_link,omitempty"`
	//
	FaceNftNew int32 `protobuf:"varint,28,opt,name=face_nft_new,json=faceNftNew,proto3" json:"face_nft_new,omitempty"`
	//
	NftFaceIcon          *NftFaceIcon `protobuf:"bytes,29,opt,name=nft_face_icon,json=nftFaceIcon,proto3" json:"nft_face_icon,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ResultItem) Reset()         { *m = ResultItem{} }
func (m *ResultItem) String() string { return proto.CompactTextString(m) }
func (*ResultItem) ProtoMessage()    {}
func (*ResultItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_e899187b61530774, []int{5}
}

func (m *ResultItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultItem.Unmarshal(m, b)
}
func (m *ResultItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultItem.Marshal(b, m, deterministic)
}
func (m *ResultItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultItem.Merge(m, src)
}
func (m *ResultItem) XXX_Size() int {
	return xxx_messageInfo_ResultItem.Size(m)
}
func (m *ResultItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultItem.DiscardUnknown(m)
}

var xxx_messageInfo_ResultItem proto.InternalMessageInfo

func (m *ResultItem) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ResultItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ResultItem) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func (m *ResultItem) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *ResultItem) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ResultItem) GetCoverSize() float64 {
	if m != nil {
		return m.CoverSize
	}
	return 0
}

func (m *ResultItem) GetSugType() string {
	if m != nil {
		return m.SugType
	}
	return ""
}

func (m *ResultItem) GetTermType() int32 {
	if m != nil {
		return m.TermType
	}
	return 0
}

func (m *ResultItem) GetGoto() string {
	if m != nil {
		return m.Goto
	}
	return ""
}

func (m *ResultItem) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *ResultItem) GetOfficialVerify() *OfficialVerify {
	if m != nil {
		return m.OfficialVerify
	}
	return nil
}

func (m *ResultItem) GetParam() string {
	if m != nil {
		return m.Param
	}
	return ""
}

func (m *ResultItem) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *ResultItem) GetFans() int32 {
	if m != nil {
		return m.Fans
	}
	return 0
}

func (m *ResultItem) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *ResultItem) GetArchives() int32 {
	if m != nil {
		return m.Archives
	}
	return 0
}

func (m *ResultItem) GetPtime() int64 {
	if m != nil {
		return m.Ptime
	}
	return 0
}

func (m *ResultItem) GetSeasonTypeName() string {
	if m != nil {
		return m.SeasonTypeName
	}
	return ""
}

func (m *ResultItem) GetArea() string {
	if m != nil {
		return m.Area
	}
	return ""
}

func (m *ResultItem) GetStyle() string {
	if m != nil {
		return m.Style
	}
	return ""
}

func (m *ResultItem) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *ResultItem) GetRating() float64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *ResultItem) GetVote() int32 {
	if m != nil {
		return m.Vote
	}
	return 0
}

func (m *ResultItem) GetBadges() []*ReasonStyle {
	if m != nil {
		return m.Badges
	}
	return nil
}

func (m *ResultItem) GetStyles() string {
	if m != nil {
		return m.Styles
	}
	return ""
}

func (m *ResultItem) GetModuleId() int64 {
	if m != nil {
		return m.ModuleId
	}
	return 0
}

func (m *ResultItem) GetLiveLink() string {
	if m != nil {
		return m.LiveLink
	}
	return ""
}

func (m *ResultItem) GetFaceNftNew() int32 {
	if m != nil {
		return m.FaceNftNew
	}
	return 0
}

func (m *ResultItem) GetNftFaceIcon() *NftFaceIcon {
	if m != nil {
		return m.NftFaceIcon
	}
	return nil
}

// 认证信息
type OfficialVerify struct {
	// 认证类型
	// 127:未认证 0:个人 1:机构
	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// 认证描述
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OfficialVerify) Reset()         { *m = OfficialVerify{} }
func (m *OfficialVerify) String() string { return proto.CompactTextString(m) }
func (*OfficialVerify) ProtoMessage()    {}
func (*OfficialVerify) Descriptor() ([]byte, []int) {
	return fileDescriptor_e899187b61530774, []int{6}
}

func (m *OfficialVerify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OfficialVerify.Unmarshal(m, b)
}
func (m *OfficialVerify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OfficialVerify.Marshal(b, m, deterministic)
}
func (m *OfficialVerify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OfficialVerify.Merge(m, src)
}
func (m *OfficialVerify) XXX_Size() int {
	return xxx_messageInfo_OfficialVerify.Size(m)
}
func (m *OfficialVerify) XXX_DiscardUnknown() {
	xxx_messageInfo_OfficialVerify.DiscardUnknown(m)
}

var xxx_messageInfo_OfficialVerify proto.InternalMessageInfo

func (m *OfficialVerify) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *OfficialVerify) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

// 角标
type ReasonStyle struct {
	// 角标文案
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// 文案日间色值
	TextColor string `protobuf:"bytes,2,opt,name=text_color,json=textColor,proto3" json:"text_color,omitempty"`
	// 文案夜间色值
	TextColorNight string `protobuf:"bytes,3,opt,name=text_color_night,json=textColorNight,proto3" json:"text_color_night,omitempty"`
	// 背景日间色值
	BgColor string `protobuf:"bytes,4,opt,name=bg_color,json=bgColor,proto3" json:"bg_color,omitempty"`
	// 背景夜间色值
	BgColorNight string `protobuf:"bytes,5,opt,name=bg_color_night,json=bgColorNight,proto3" json:"bg_color_night,omitempty"`
	// 边框日间色值
	BorderColor string `protobuf:"bytes,6,opt,name=border_color,json=borderColor,proto3" json:"border_color,omitempty"`
	// 边框夜间色值
	BorderColorNight string `protobuf:"bytes,7,opt,name=border_color_night,json=borderColorNight,proto3" json:"border_color_night,omitempty"`
	// 角标样式
	// 1:填充模式 2:镂空模式
	BgStyle              int32    `protobuf:"varint,8,opt,name=bg_style,json=bgStyle,proto3" json:"bg_style,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReasonStyle) Reset()         { *m = ReasonStyle{} }
func (m *ReasonStyle) String() string { return proto.CompactTextString(m) }
func (*ReasonStyle) ProtoMessage()    {}
func (*ReasonStyle) Descriptor() ([]byte, []int) {
	return fileDescriptor_e899187b61530774, []int{7}
}

func (m *ReasonStyle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReasonStyle.Unmarshal(m, b)
}
func (m *ReasonStyle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReasonStyle.Marshal(b, m, deterministic)
}
func (m *ReasonStyle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReasonStyle.Merge(m, src)
}
func (m *ReasonStyle) XXX_Size() int {
	return xxx_messageInfo_ReasonStyle.Size(m)
}
func (m *ReasonStyle) XXX_DiscardUnknown() {
	xxx_messageInfo_ReasonStyle.DiscardUnknown(m)
}

var xxx_messageInfo_ReasonStyle proto.InternalMessageInfo

func (m *ReasonStyle) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *ReasonStyle) GetTextColor() string {
	if m != nil {
		return m.TextColor
	}
	return ""
}

func (m *ReasonStyle) GetTextColorNight() string {
	if m != nil {
		return m.TextColorNight
	}
	return ""
}

func (m *ReasonStyle) GetBgColor() string {
	if m != nil {
		return m.BgColor
	}
	return ""
}

func (m *ReasonStyle) GetBgColorNight() string {
	if m != nil {
		return m.BgColorNight
	}
	return ""
}

func (m *ReasonStyle) GetBorderColor() string {
	if m != nil {
		return m.BorderColor
	}
	return ""
}

func (m *ReasonStyle) GetBorderColorNight() string {
	if m != nil {
		return m.BorderColorNight
	}
	return ""
}

func (m *ReasonStyle) GetBgStyle() int32 {
	if m != nil {
		return m.BgStyle
	}
	return 0
}

func init() {
	proto.RegisterType((*DefaultWordsReply)(nil), "bilibili.app.interface.v1.DefaultWordsReply")
	proto.RegisterType((*NftFaceIcon)(nil), "bilibili.app.interface.v1.NftFaceIcon")
	proto.RegisterType((*DefaultWordsReq)(nil), "bilibili.app.interface.v1.DefaultWordsReq")
	proto.RegisterType((*SuggestionResult3Reply)(nil), "bilibili.app.interface.v1.SuggestionResult3Reply")
	proto.RegisterType((*SuggestionResult3Req)(nil), "bilibili.app.interface.v1.SuggestionResult3Req")
	proto.RegisterType((*ResultItem)(nil), "bilibili.app.interface.v1.ResultItem")
	proto.RegisterType((*OfficialVerify)(nil), "bilibili.app.interface.v1.OfficialVerify")
	proto.RegisterType((*ReasonStyle)(nil), "bilibili.app.interface.v1.ReasonStyle")
}

func init() {
	proto.RegisterFile("grpc_api/bilibili/app/interfaces/v1/search.proto", fileDescriptor_e899187b61530774)
}

var fileDescriptor_e899187b61530774 = []byte{
	// 1102 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0x97, 0xe3, 0xc4, 0x7f, 0x9e, 0x53, 0xd7, 0x1d, 0x42, 0x3b, 0x49, 0x5b, 0x91, 0x5a, 0x14,
	0x05, 0x54, 0xd9, 0x34, 0xbd, 0xc0, 0x85, 0x0b, 0xb4, 0x52, 0x10, 0x18, 0x69, 0x5d, 0xc1, 0x71,
	0x35, 0xde, 0x7d, 0x5e, 0x8f, 0xb2, 0xde, 0xd9, 0xce, 0x8c, 0x9d, 0xb8, 0xe2, 0xcc, 0x07, 0xe0,
	0xc2, 0xc7, 0x43, 0xe2, 0xce, 0x77, 0x40, 0xef, 0xcd, 0x3a, 0xd9, 0xa0, 0x12, 0xca, 0x85, 0x43,
	0x94, 0xf7, 0x7e, 0xf3, 0xfe, 0xff, 0x5b, 0xc3, 0xe7, 0x99, 0x2d, 0x93, 0x58, 0x95, 0x7a, 0x3c,
	0xd3, 0xb9, 0xa6, 0xbf, 0xb1, 0x2a, 0xcb, 0xb1, 0x2e, 0x3c, 0xda, 0xb9, 0x4a, 0xd0, 0x8d, 0xd7,
	0xcf, 0xc7, 0x0e, 0x95, 0x4d, 0x16, 0xa3, 0xd2, 0x1a, 0x6f, 0xc4, 0xe1, 0x56, 0x70, 0xa4, 0xca,
	0x72, 0x74, 0x25, 0x38, 0x5a, 0x3f, 0x1f, 0xfe, 0xde, 0x80, 0x7b, 0xdf, 0xe0, 0x5c, 0xad, 0x72,
	0xff, 0x93, 0xb1, 0xa9, 0x8b, 0xb0, 0xcc, 0x37, 0x42, 0x42, 0xdb, 0x5b, 0x95, 0x9c, 0xeb, 0x54,
	0x36, 0x8e, 0x1b, 0x27, 0xdd, 0x68, 0xcb, 0x8a, 0x03, 0xd8, 0x2b, 0x95, 0x55, 0x4b, 0xb9, 0xc3,
	0x78, 0x60, 0x84, 0x80, 0x5d, 0xb7, 0x30, 0x17, 0xb2, 0xc9, 0x20, 0xd3, 0x84, 0x5d, 0x18, 0x9b,
	0xca, 0xdd, 0x80, 0x11, 0x2d, 0x1e, 0x03, 0xd0, 0x5b, 0x3c, 0xb7, 0xa6, 0xf0, 0x72, 0xef, 0xb8,
	0x71, 0xd2, 0x8c, 0xba, 0x84, 0xbc, 0x22, 0x40, 0x3c, 0x80, 0x36, 0x5e, 0x96, 0xb1, 0xf3, 0x56,
	0xb6, 0x58, 0xab, 0x85, 0x97, 0xe5, 0xd4, 0x5b, 0xb2, 0x95, 0x19, 0x6f, 0x64, 0x3b, 0xd8, 0x22,
	0x9a, 0x22, 0x59, 0xab, 0x7c, 0x85, 0xb2, 0x13, 0x22, 0x61, 0x46, 0x0c, 0xa0, 0xb9, 0xb2, 0x5a,
	0x76, 0x19, 0x23, 0x72, 0x98, 0x40, 0x6f, 0x32, 0xf7, 0xaf, 0x54, 0x82, 0x67, 0x89, 0x29, 0xc4,
	0x47, 0xd0, 0xb3, 0x98, 0x69, 0x53, 0xc4, 0x7e, 0x53, 0x22, 0xa7, 0xb7, 0x17, 0x41, 0x80, 0x5e,
	0x6f, 0x4a, 0x24, 0x5f, 0x3a, 0x31, 0x45, 0x95, 0x20, 0xd3, 0xa4, 0xc4, 0x71, 0x3b, 0xaf, 0xfc,
	0xca, 0x71, 0x9a, 0x7b, 0x11, 0xa7, 0x32, 0x65, 0x64, 0xf8, 0xeb, 0x0e, 0xdc, 0xbd, 0x59, 0xc6,
	0x37, 0x64, 0x68, 0x6e, 0xcd, 0x92, 0x5d, 0x34, 0x23, 0xa6, 0xc9, 0x50, 0x6e, 0x32, 0x5d, 0xc4,
	0xb8, 0xc6, 0xc2, 0xb3, 0x8f, 0x66, 0x04, 0x0c, 0xbd, 0x24, 0x44, 0x3c, 0x85, 0xbe, 0x47, 0x2c,
	0x54, 0x86, 0xd6, 0xc5, 0x4b, 0x93, 0x62, 0xe5, 0xec, 0xce, 0x15, 0xfa, 0xbd, 0x49, 0x51, 0x3c,
	0x81, 0xfd, 0x1c, 0x9d, 0x33, 0x45, 0x25, 0xb4, 0xcb, 0x42, 0xbd, 0x0a, 0x63, 0x91, 0x01, 0x34,
	0xbd, 0x9a, 0x71, 0x91, 0xbb, 0x11, 0x91, 0xe2, 0x10, 0x3a, 0xec, 0x36, 0xd6, 0x69, 0x55, 0xdf,
	0x36, 0xf3, 0x67, 0x29, 0xc5, 0xaa, 0xd6, 0x3a, 0xdd, 0x16, 0x98, 0x68, 0x2a, 0xf0, 0x9b, 0x15,
	0xda, 0xcd, 0xb6, 0xc0, 0xcc, 0x88, 0x3e, 0xec, 0xa8, 0x82, 0xeb, 0xdb, 0x8c, 0x76, 0x54, 0x41,
	0x46, 0xb5, 0x8b, 0xe7, 0x16, 0xdd, 0x42, 0x02, 0xa3, 0x6d, 0xed, 0x5e, 0x11, 0x3b, 0xfc, 0xa5,
	0x01, 0xf7, 0xa7, 0xab, 0x2c, 0x43, 0xe7, 0xb5, 0x29, 0x22, 0x74, 0xab, 0xdc, 0xbf, 0xf8, 0xb7,
	0x01, 0xfb, 0x12, 0x76, 0x73, 0xed, 0xa8, 0x34, 0xcd, 0x93, 0xde, 0xe9, 0xd3, 0xd1, 0x3f, 0x8e,
	0xee, 0x28, 0x18, 0x3c, 0xf3, 0xb8, 0x8c, 0x58, 0xa5, 0x3e, 0x3e, 0xcd, 0xfa, 0xf8, 0x0c, 0x57,
	0x70, 0xf0, 0x8e, 0x38, 0xde, 0x50, 0x14, 0xe7, 0xb8, 0xe1, 0x29, 0xad, 0xa2, 0xa8, 0x58, 0xf1,
	0x08, 0xba, 0x0b, 0x9d, 0x2d, 0x72, 0x9d, 0x2d, 0x42, 0x97, 0xf6, 0xa2, 0x6b, 0xe0, 0x3d, 0x9b,
	0x34, 0xfc, 0xb3, 0x05, 0x70, 0x1d, 0xe4, 0x8d, 0x79, 0xe8, 0x56, 0xf3, 0x70, 0x00, 0x7b, 0x5e,
	0xfb, 0x1c, 0xb7, 0xeb, 0xc4, 0x4c, 0x3d, 0xae, 0xe6, 0xcd, 0xb8, 0x8e, 0xa0, 0x53, 0x1a, 0xa7,
	0x29, 0x8f, 0xaa, 0xe7, 0x57, 0x3c, 0xd9, 0x4a, 0xcc, 0x1a, 0x6d, 0xd5, 0xf2, 0xc0, 0xd0, 0xca,
	0x31, 0x11, 0x3b, 0xfd, 0x16, 0xb9, 0xed, 0x8d, 0xa8, 0xcb, 0xc8, 0x54, 0xbf, 0x45, 0x6a, 0x9f,
	0x5b, 0x65, 0x61, 0x17, 0x42, 0xf3, 0xdb, 0x6e, 0x95, 0xf1, 0x22, 0x3c, 0x84, 0xae, 0x47, 0xbb,
	0x0c, 0x6f, 0x9d, 0xe0, 0x8c, 0x80, 0xed, 0x96, 0xf0, 0x46, 0x76, 0x6b, 0x1b, 0x59, 0xed, 0x1e,
	0x5c, 0xed, 0x9e, 0x88, 0xe0, 0xae, 0x99, 0xcf, 0x75, 0xa2, 0x55, 0x1e, 0xaf, 0xd1, 0xea, 0xf9,
	0x46, 0xf6, 0x8e, 0x1b, 0x27, 0xbd, 0xd3, 0x4f, 0x6f, 0xe9, 0xeb, 0x0f, 0x95, 0xc6, 0x8f, 0xac,
	0x10, 0xf5, 0xcd, 0x0d, 0xfe, 0xfa, 0x02, 0xed, 0xd7, 0x2f, 0xd0, 0x00, 0x9a, 0x4b, 0x9d, 0xca,
	0x3b, 0x3c, 0x81, 0x44, 0x72, 0xb9, 0x55, 0xe1, 0x64, 0x9f, 0x23, 0x67, 0x9a, 0x74, 0x73, 0x5c,
	0x63, 0x2e, 0xef, 0x32, 0x18, 0x18, 0x2a, 0x2a, 0x1d, 0x4b, 0xbd, 0x46, 0x27, 0x07, 0x21, 0xcf,
	0x2d, 0xcf, 0xde, 0xbc, 0x5e, 0xa2, 0xbc, 0xc7, 0x96, 0x03, 0x23, 0x4e, 0x60, 0xe0, 0x50, 0xb9,
	0xea, 0x88, 0xc4, 0x85, 0x5a, 0xa2, 0x14, 0x1c, 0x4e, 0x3f, 0xe0, 0x54, 0xa3, 0x89, 0x5a, 0x72,
	0x9d, 0x94, 0x45, 0x25, 0x3f, 0xa8, 0x16, 0xcb, 0xa2, 0x22, 0x9b, 0xce, 0x6f, 0x72, 0x94, 0x07,
	0x21, 0x03, 0x66, 0x38, 0x36, 0x35, 0xc3, 0x5c, 0x7e, 0x18, 0x50, 0x66, 0xc4, 0x7d, 0x68, 0x59,
	0xe5, 0x75, 0x91, 0xc9, 0xfb, 0xdc, 0xba, 0x8a, 0x23, 0xbb, 0x6b, 0xe3, 0x51, 0x3e, 0x08, 0xd9,
	0x11, 0x2d, 0xbe, 0x82, 0xd6, 0x4c, 0xa5, 0x19, 0x3a, 0x29, 0x79, 0x79, 0x3e, 0xb9, 0x75, 0x79,
	0x28, 0xcc, 0x29, 0x79, 0x8e, 0x2a, 0x2d, 0xf2, 0xc5, 0xa1, 0x38, 0x79, 0x18, 0xd6, 0x27, 0x70,
	0x34, 0x08, 0x4b, 0x93, 0xae, 0x72, 0xa4, 0xc3, 0x71, 0xc4, 0x75, 0xe8, 0x04, 0xe0, 0x2c, 0xa5,
	0xc7, 0x5c, 0xaf, 0x31, 0xce, 0x75, 0x71, 0x2e, 0x1f, 0xb2, 0x5e, 0x87, 0x80, 0xef, 0x74, 0x71,
	0x2e, 0x8e, 0x61, 0x9f, 0x1c, 0xc6, 0xc5, 0xdc, 0xc7, 0x05, 0x5e, 0xc8, 0x47, 0xe1, 0x70, 0x12,
	0x36, 0x99, 0xfb, 0x09, 0x5e, 0x88, 0x6f, 0xe1, 0x0e, 0x3d, 0xb2, 0x14, 0x9f, 0xdd, 0xc7, 0x3c,
	0x1f, 0xb7, 0x85, 0x5e, 0xbb, 0xe6, 0x51, 0xaf, 0xb8, 0x66, 0x86, 0x5f, 0x40, 0xff, 0xe6, 0xec,
	0x50, 0x95, 0x6a, 0x57, 0x9e, 0x69, 0xc2, 0x52, 0x74, 0xc9, 0xf6, 0xbe, 0x13, 0x3d, 0xfc, 0x6d,
	0x07, 0x7a, 0xb5, 0x8a, 0xb0, 0x1e, 0x5e, 0xfa, 0xed, 0xaa, 0x12, 0x4d, 0x8b, 0x44, 0xff, 0xe3,
	0xc4, 0xe4, 0xc6, 0x56, 0xda, 0x5d, 0x42, 0xbe, 0x26, 0x80, 0x46, 0xe2, 0xfa, 0x39, 0x2e, 0xf8,
	0x70, 0x84, 0xe5, 0xed, 0x5f, 0x09, 0x4d, 0xf8, 0x7a, 0x1c, 0x42, 0x67, 0x96, 0x55, 0x66, 0xc2,
	0xc7, 0xb1, 0x3d, 0xcb, 0x82, 0x91, 0x8f, 0xa1, 0xbf, 0x7d, 0xaa, 0x4c, 0x84, 0x5d, 0xde, 0xaf,
	0x04, 0x82, 0x81, 0x27, 0xb0, 0x3f, 0x33, 0x36, 0x45, 0x5b, 0x19, 0x09, 0xb7, 0xbc, 0x17, 0xb0,
	0x60, 0xe8, 0x19, 0x88, 0xba, 0x48, 0x65, 0x2c, 0x2c, 0xf8, 0xa0, 0x26, 0x58, 0x8f, 0x28, 0xcc,
	0x64, 0x58, 0xf4, 0xf6, 0x2c, 0xe3, 0x4a, 0x9c, 0xfe, 0xd1, 0x80, 0xd6, 0x94, 0x7f, 0x4b, 0x88,
	0x02, 0x3a, 0xd5, 0x15, 0x7d, 0x21, 0xc6, 0xb7, 0xf4, 0xe7, 0x5d, 0xa7, 0xf6, 0xe8, 0xf9, 0x7f,
	0x53, 0xa0, 0x6f, 0xc4, 0x02, 0xf6, 0xeb, 0x9f, 0x54, 0xf1, 0xd9, 0x2d, 0x26, 0xfe, 0xf6, 0xed,
	0x3d, 0x7a, 0xf6, 0xde, 0xb2, 0x65, 0xbe, 0x39, 0xfd, 0x19, 0x20, 0xe4, 0xf8, 0x1a, 0x9d, 0xa7,
	0x3c, 0x27, 0xc6, 0xbf, 0xbc, 0xa4, 0x4f, 0xca, 0xff, 0x90, 0xe7, 0xac, 0xc5, 0x3f, 0xd2, 0x5e,
	0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0xd9, 0xad, 0xed, 0x8e, 0xd8, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SearchClient is the client API for Search service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchClient interface {
	// 获取搜索建议
	Suggest3(ctx context.Context, in *SuggestionResult3Req, opts ...grpc.CallOption) (*SuggestionResult3Reply, error)
	//
	DefaultWords(ctx context.Context, in *DefaultWordsReq, opts ...grpc.CallOption) (*DefaultWordsReply, error)
}

type searchClient struct {
	cc *grpc.ClientConn
}

func NewSearchClient(cc *grpc.ClientConn) SearchClient {
	return &searchClient{cc}
}

func (c *searchClient) Suggest3(ctx context.Context, in *SuggestionResult3Req, opts ...grpc.CallOption) (*SuggestionResult3Reply, error) {
	out := new(SuggestionResult3Reply)
	err := c.cc.Invoke(ctx, "/bilibili.app.interface.v1.Search/Suggest3", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) DefaultWords(ctx context.Context, in *DefaultWordsReq, opts ...grpc.CallOption) (*DefaultWordsReply, error) {
	out := new(DefaultWordsReply)
	err := c.cc.Invoke(ctx, "/bilibili.app.interface.v1.Search/DefaultWords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServer is the server API for Search service.
type SearchServer interface {
	// 获取搜索建议
	Suggest3(context.Context, *SuggestionResult3Req) (*SuggestionResult3Reply, error)
	//
	DefaultWords(context.Context, *DefaultWordsReq) (*DefaultWordsReply, error)
}

// UnimplementedSearchServer can be embedded to have forward compatible implementations.
type UnimplementedSearchServer struct {
}

func (*UnimplementedSearchServer) Suggest3(ctx context.Context, req *SuggestionResult3Req) (*SuggestionResult3Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Suggest3 not implemented")
}
func (*UnimplementedSearchServer) DefaultWords(ctx context.Context, req *DefaultWordsReq) (*DefaultWordsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DefaultWords not implemented")
}

func RegisterSearchServer(s *grpc.Server, srv SearchServer) {
	s.RegisterService(&_Search_serviceDesc, srv)
}

func _Search_Suggest3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestionResult3Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).Suggest3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.interface.v1.Search/Suggest3",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).Suggest3(ctx, req.(*SuggestionResult3Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_DefaultWords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefaultWordsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).DefaultWords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.interface.v1.Search/DefaultWords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).DefaultWords(ctx, req.(*DefaultWordsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Search_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.interface.v1.Search",
	HandlerType: (*SearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Suggest3",
			Handler:    _Search_Suggest3_Handler,
		},
		{
			MethodName: "DefaultWords",
			Handler:    _Search_DefaultWords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_api/bilibili/app/interfaces/v1/search.proto",
}

// SearchTestClient is the client API for SearchTest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchTestClient interface {
	//
	NotExist(ctx context.Context, in *SuggestionResult3Req, opts ...grpc.CallOption) (*SuggestionResult3Reply, error)
}

type searchTestClient struct {
	cc *grpc.ClientConn
}

func NewSearchTestClient(cc *grpc.ClientConn) SearchTestClient {
	return &searchTestClient{cc}
}

func (c *searchTestClient) NotExist(ctx context.Context, in *SuggestionResult3Req, opts ...grpc.CallOption) (*SuggestionResult3Reply, error) {
	out := new(SuggestionResult3Reply)
	err := c.cc.Invoke(ctx, "/bilibili.app.interface.v1.SearchTest/NotExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchTestServer is the server API for SearchTest service.
type SearchTestServer interface {
	//
	NotExist(context.Context, *SuggestionResult3Req) (*SuggestionResult3Reply, error)
}

// UnimplementedSearchTestServer can be embedded to have forward compatible implementations.
type UnimplementedSearchTestServer struct {
}

func (*UnimplementedSearchTestServer) NotExist(ctx context.Context, req *SuggestionResult3Req) (*SuggestionResult3Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotExist not implemented")
}

func RegisterSearchTestServer(s *grpc.Server, srv SearchTestServer) {
	s.RegisterService(&_SearchTest_serviceDesc, srv)
}

func _SearchTest_NotExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestionResult3Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchTestServer).NotExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bilibili.app.interface.v1.SearchTest/NotExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchTestServer).NotExist(ctx, req.(*SuggestionResult3Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchTest_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.app.interface.v1.SearchTest",
	HandlerType: (*SearchTestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotExist",
			Handler:    _SearchTest_NotExist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_api/bilibili/app/interfaces/v1/search.proto",
}
