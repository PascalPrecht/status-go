// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat_message.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ImageMessage_ImageType int32

const (
	ImageMessage_UNKNOWN_IMAGE_TYPE ImageMessage_ImageType = 0
	ImageMessage_PNG                ImageMessage_ImageType = 1
	ImageMessage_JPEG               ImageMessage_ImageType = 2
	ImageMessage_WEBP               ImageMessage_ImageType = 3
	ImageMessage_GIF                ImageMessage_ImageType = 4
)

var ImageMessage_ImageType_name = map[int32]string{
	0: "UNKNOWN_IMAGE_TYPE",
	1: "PNG",
	2: "JPEG",
	3: "WEBP",
	4: "GIF",
}

var ImageMessage_ImageType_value = map[string]int32{
	"UNKNOWN_IMAGE_TYPE": 0,
	"PNG":                1,
	"JPEG":               2,
	"WEBP":               3,
	"GIF":                4,
}

func (x ImageMessage_ImageType) String() string {
	return proto.EnumName(ImageMessage_ImageType_name, int32(x))
}

func (ImageMessage_ImageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_263952f55fd35689, []int{1, 0}
}

type AudioMessage_AudioType int32

const (
	AudioMessage_UNKNOWN_AUDIO_TYPE AudioMessage_AudioType = 0
	AudioMessage_AAC                AudioMessage_AudioType = 1
	AudioMessage_AMR                AudioMessage_AudioType = 2
)

var AudioMessage_AudioType_name = map[int32]string{
	0: "UNKNOWN_AUDIO_TYPE",
	1: "AAC",
	2: "AMR",
}

var AudioMessage_AudioType_value = map[string]int32{
	"UNKNOWN_AUDIO_TYPE": 0,
	"AAC":                1,
	"AMR":                2,
}

func (x AudioMessage_AudioType) String() string {
	return proto.EnumName(AudioMessage_AudioType_name, int32(x))
}

func (AudioMessage_AudioType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_263952f55fd35689, []int{2, 0}
}

type ChatMessage_ContentType int32

const (
	ChatMessage_UNKNOWN_CONTENT_TYPE ChatMessage_ContentType = 0
	ChatMessage_TEXT_PLAIN           ChatMessage_ContentType = 1
	ChatMessage_STICKER              ChatMessage_ContentType = 2
	ChatMessage_STATUS               ChatMessage_ContentType = 3
	ChatMessage_EMOJI                ChatMessage_ContentType = 4
	ChatMessage_TRANSACTION_COMMAND  ChatMessage_ContentType = 5
	// Only local
	ChatMessage_SYSTEM_MESSAGE_CONTENT_PRIVATE_GROUP ChatMessage_ContentType = 6
	ChatMessage_IMAGE                                ChatMessage_ContentType = 7
	ChatMessage_AUDIO                                ChatMessage_ContentType = 8
	ChatMessage_ORGANISATION                         ChatMessage_ContentType = 9
)

var ChatMessage_ContentType_name = map[int32]string{
	0: "UNKNOWN_CONTENT_TYPE",
	1: "TEXT_PLAIN",
	2: "STICKER",
	3: "STATUS",
	4: "EMOJI",
	5: "TRANSACTION_COMMAND",
	6: "SYSTEM_MESSAGE_CONTENT_PRIVATE_GROUP",
	7: "IMAGE",
	8: "AUDIO",
	9: "ORGANISATION",
}

var ChatMessage_ContentType_value = map[string]int32{
	"UNKNOWN_CONTENT_TYPE":                 0,
	"TEXT_PLAIN":                           1,
	"STICKER":                              2,
	"STATUS":                               3,
	"EMOJI":                                4,
	"TRANSACTION_COMMAND":                  5,
	"SYSTEM_MESSAGE_CONTENT_PRIVATE_GROUP": 6,
	"IMAGE":                                7,
	"AUDIO":                                8,
	"ORGANISATION":                         9,
}

func (x ChatMessage_ContentType) String() string {
	return proto.EnumName(ChatMessage_ContentType_name, int32(x))
}

func (ChatMessage_ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_263952f55fd35689, []int{3, 0}
}

type StickerMessage struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Pack                 int32    `protobuf:"varint,2,opt,name=pack,proto3" json:"pack,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StickerMessage) Reset()         { *m = StickerMessage{} }
func (m *StickerMessage) String() string { return proto.CompactTextString(m) }
func (*StickerMessage) ProtoMessage()    {}
func (*StickerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_263952f55fd35689, []int{0}
}

func (m *StickerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StickerMessage.Unmarshal(m, b)
}
func (m *StickerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StickerMessage.Marshal(b, m, deterministic)
}
func (m *StickerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StickerMessage.Merge(m, src)
}
func (m *StickerMessage) XXX_Size() int {
	return xxx_messageInfo_StickerMessage.Size(m)
}
func (m *StickerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StickerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StickerMessage proto.InternalMessageInfo

func (m *StickerMessage) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *StickerMessage) GetPack() int32 {
	if m != nil {
		return m.Pack
	}
	return 0
}

type ImageMessage struct {
	Payload              []byte                 `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Type                 ImageMessage_ImageType `protobuf:"varint,2,opt,name=type,proto3,enum=protobuf.ImageMessage_ImageType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ImageMessage) Reset()         { *m = ImageMessage{} }
func (m *ImageMessage) String() string { return proto.CompactTextString(m) }
func (*ImageMessage) ProtoMessage()    {}
func (*ImageMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_263952f55fd35689, []int{1}
}

func (m *ImageMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageMessage.Unmarshal(m, b)
}
func (m *ImageMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageMessage.Marshal(b, m, deterministic)
}
func (m *ImageMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageMessage.Merge(m, src)
}
func (m *ImageMessage) XXX_Size() int {
	return xxx_messageInfo_ImageMessage.Size(m)
}
func (m *ImageMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ImageMessage proto.InternalMessageInfo

func (m *ImageMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ImageMessage) GetType() ImageMessage_ImageType {
	if m != nil {
		return m.Type
	}
	return ImageMessage_UNKNOWN_IMAGE_TYPE
}

type AudioMessage struct {
	Payload              []byte                 `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Type                 AudioMessage_AudioType `protobuf:"varint,2,opt,name=type,proto3,enum=protobuf.AudioMessage_AudioType" json:"type,omitempty"`
	DurationMs           uint64                 `protobuf:"varint,3,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AudioMessage) Reset()         { *m = AudioMessage{} }
func (m *AudioMessage) String() string { return proto.CompactTextString(m) }
func (*AudioMessage) ProtoMessage()    {}
func (*AudioMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_263952f55fd35689, []int{2}
}

func (m *AudioMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AudioMessage.Unmarshal(m, b)
}
func (m *AudioMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AudioMessage.Marshal(b, m, deterministic)
}
func (m *AudioMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AudioMessage.Merge(m, src)
}
func (m *AudioMessage) XXX_Size() int {
	return xxx_messageInfo_AudioMessage.Size(m)
}
func (m *AudioMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AudioMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AudioMessage proto.InternalMessageInfo

func (m *AudioMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *AudioMessage) GetType() AudioMessage_AudioType {
	if m != nil {
		return m.Type
	}
	return AudioMessage_UNKNOWN_AUDIO_TYPE
}

func (m *AudioMessage) GetDurationMs() uint64 {
	if m != nil {
		return m.DurationMs
	}
	return 0
}

type ChatMessage struct {
	// Lamport timestamp of the chat message
	Clock uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	// Unix timestamps in milliseconds, currently not used as we use whisper as more reliable, but here
	// so that we don't rely on it
	Timestamp uint64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Text of the message
	Text string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	// Id of the message that we are replying to
	ResponseTo string `protobuf:"bytes,4,opt,name=response_to,json=responseTo,proto3" json:"response_to,omitempty"`
	// Ens name of the sender
	EnsName string `protobuf:"bytes,5,opt,name=ens_name,json=ensName,proto3" json:"ens_name,omitempty"`
	// Chat id, this field is symmetric for public-chats and private group chats,
	// but asymmetric in case of one-to-ones, as the sender will use the chat-id
	// of the received, while the receiver will use the chat-id of the sender.
	// Probably should be the concatenation of sender-pk & receiver-pk in alphabetical order
	ChatId string `protobuf:"bytes,6,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	// The type of message (public/one-to-one/private-group-chat)
	MessageType MessageType `protobuf:"varint,7,opt,name=message_type,json=messageType,proto3,enum=protobuf.MessageType" json:"message_type,omitempty"`
	// The type of the content of the message
	ContentType ChatMessage_ContentType `protobuf:"varint,8,opt,name=content_type,json=contentType,proto3,enum=protobuf.ChatMessage_ContentType" json:"content_type,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*ChatMessage_Sticker
	//	*ChatMessage_Image
	//	*ChatMessage_Audio
	//	*ChatMessage_Organisation
	Payload              isChatMessage_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ChatMessage) Reset()         { *m = ChatMessage{} }
func (m *ChatMessage) String() string { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()    {}
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_263952f55fd35689, []int{3}
}

func (m *ChatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMessage.Unmarshal(m, b)
}
func (m *ChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMessage.Marshal(b, m, deterministic)
}
func (m *ChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMessage.Merge(m, src)
}
func (m *ChatMessage) XXX_Size() int {
	return xxx_messageInfo_ChatMessage.Size(m)
}
func (m *ChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMessage proto.InternalMessageInfo

func (m *ChatMessage) GetClock() uint64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *ChatMessage) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ChatMessage) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *ChatMessage) GetResponseTo() string {
	if m != nil {
		return m.ResponseTo
	}
	return ""
}

func (m *ChatMessage) GetEnsName() string {
	if m != nil {
		return m.EnsName
	}
	return ""
}

func (m *ChatMessage) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *ChatMessage) GetMessageType() MessageType {
	if m != nil {
		return m.MessageType
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *ChatMessage) GetContentType() ChatMessage_ContentType {
	if m != nil {
		return m.ContentType
	}
	return ChatMessage_UNKNOWN_CONTENT_TYPE
}

type isChatMessage_Payload interface {
	isChatMessage_Payload()
}

type ChatMessage_Sticker struct {
	Sticker *StickerMessage `protobuf:"bytes,9,opt,name=sticker,proto3,oneof"`
}

type ChatMessage_Image struct {
	Image *ImageMessage `protobuf:"bytes,10,opt,name=image,proto3,oneof"`
}

type ChatMessage_Audio struct {
	Audio *AudioMessage `protobuf:"bytes,11,opt,name=audio,proto3,oneof"`
}

type ChatMessage_Organisation struct {
	Organisation []byte `protobuf:"bytes,12,opt,name=organisation,proto3,oneof"`
}

func (*ChatMessage_Sticker) isChatMessage_Payload() {}

func (*ChatMessage_Image) isChatMessage_Payload() {}

func (*ChatMessage_Audio) isChatMessage_Payload() {}

func (*ChatMessage_Organisation) isChatMessage_Payload() {}

func (m *ChatMessage) GetPayload() isChatMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ChatMessage) GetSticker() *StickerMessage {
	if x, ok := m.GetPayload().(*ChatMessage_Sticker); ok {
		return x.Sticker
	}
	return nil
}

func (m *ChatMessage) GetImage() *ImageMessage {
	if x, ok := m.GetPayload().(*ChatMessage_Image); ok {
		return x.Image
	}
	return nil
}

func (m *ChatMessage) GetAudio() *AudioMessage {
	if x, ok := m.GetPayload().(*ChatMessage_Audio); ok {
		return x.Audio
	}
	return nil
}

func (m *ChatMessage) GetOrganisation() []byte {
	if x, ok := m.GetPayload().(*ChatMessage_Organisation); ok {
		return x.Organisation
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ChatMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ChatMessage_Sticker)(nil),
		(*ChatMessage_Image)(nil),
		(*ChatMessage_Audio)(nil),
		(*ChatMessage_Organisation)(nil),
	}
}

func init() {
	proto.RegisterEnum("protobuf.ImageMessage_ImageType", ImageMessage_ImageType_name, ImageMessage_ImageType_value)
	proto.RegisterEnum("protobuf.AudioMessage_AudioType", AudioMessage_AudioType_name, AudioMessage_AudioType_value)
	proto.RegisterEnum("protobuf.ChatMessage_ContentType", ChatMessage_ContentType_name, ChatMessage_ContentType_value)
	proto.RegisterType((*StickerMessage)(nil), "protobuf.StickerMessage")
	proto.RegisterType((*ImageMessage)(nil), "protobuf.ImageMessage")
	proto.RegisterType((*AudioMessage)(nil), "protobuf.AudioMessage")
	proto.RegisterType((*ChatMessage)(nil), "protobuf.ChatMessage")
}

func init() {
	proto.RegisterFile("chat_message.proto", fileDescriptor_263952f55fd35689)
}

var fileDescriptor_263952f55fd35689 = []byte{
	// 650 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x6e, 0x9b, 0x4c,
	0x14, 0x35, 0x36, 0x36, 0xe6, 0x82, 0xa2, 0xd1, 0x7c, 0xf9, 0x12, 0x5a, 0x55, 0xaa, 0x6b, 0x65,
	0xe1, 0x95, 0x17, 0x69, 0x2a, 0x65, 0x4b, 0x1c, 0xea, 0x90, 0x14, 0x6c, 0x0d, 0xe3, 0xa6, 0x59,
	0xa1, 0x09, 0x9e, 0xc6, 0x28, 0xe1, 0x47, 0x06, 0x4b, 0xcd, 0x23, 0xf5, 0x05, 0xba, 0xe8, 0x03,
	0xf4, 0xb9, 0xaa, 0x19, 0x42, 0x20, 0x55, 0x77, 0x5d, 0xf9, 0xfe, 0x9c, 0x73, 0x7d, 0xee, 0xe5,
	0x0c, 0xe0, 0x68, 0xc3, 0xca, 0x30, 0xe1, 0x45, 0xc1, 0xee, 0xf8, 0x34, 0xdf, 0x66, 0x65, 0x86,
	0x87, 0xf2, 0xe7, 0x76, 0xf7, 0xf5, 0xb5, 0xc1, 0xd3, 0x5d, 0x52, 0x54, 0xe5, 0xf1, 0x29, 0xec,
	0x05, 0x65, 0x1c, 0xdd, 0xf3, 0xad, 0x57, 0xc1, 0x31, 0x06, 0x75, 0xc3, 0x8a, 0x8d, 0xa5, 0x8c,
	0x94, 0x89, 0x4e, 0x64, 0x2c, 0x6a, 0x39, 0x8b, 0xee, 0xad, 0xee, 0x48, 0x99, 0xf4, 0x89, 0x8c,
	0xc7, 0xdf, 0x15, 0x30, 0xdd, 0x84, 0xdd, 0xf1, 0x9a, 0x68, 0x81, 0x96, 0xb3, 0xc7, 0x87, 0x8c,
	0xad, 0x25, 0xd7, 0x24, 0x75, 0x8a, 0x4f, 0x40, 0x2d, 0x1f, 0x73, 0x2e, 0xe9, 0x7b, 0xc7, 0xa3,
	0x69, 0x2d, 0x65, 0xda, 0xe6, 0x57, 0x09, 0x7d, 0xcc, 0x39, 0x91, 0xe8, 0xb1, 0x0b, 0xfa, 0x73,
	0x09, 0x1f, 0x00, 0x5e, 0xf9, 0x57, 0xfe, 0xe2, 0xda, 0x0f, 0x5d, 0xcf, 0x9e, 0x3b, 0x21, 0xbd,
	0x59, 0x3a, 0xa8, 0x83, 0x35, 0xe8, 0x2d, 0xfd, 0x39, 0x52, 0xf0, 0x10, 0xd4, 0xcb, 0xa5, 0x33,
	0x47, 0x5d, 0x11, 0x5d, 0x3b, 0x67, 0x4b, 0xd4, 0x13, 0xcd, 0xb9, 0xfb, 0x11, 0xa9, 0xe3, 0x1f,
	0x0a, 0x98, 0xf6, 0x6e, 0x1d, 0x67, 0xff, 0xa0, 0xb5, 0xcd, 0xaf, 0x92, 0x46, 0x2b, 0x7e, 0x0b,
	0xc6, 0x7a, 0xb7, 0x65, 0x65, 0x9c, 0xa5, 0x61, 0x52, 0x58, 0xbd, 0x91, 0x32, 0x51, 0x09, 0xd4,
	0x25, 0xaf, 0x18, 0x7f, 0x00, 0xfd, 0x99, 0xd3, 0x5e, 0xc6, 0x5e, 0x9d, 0xbb, 0x8b, 0xd6, 0x32,
	0xb6, 0x3d, 0x43, 0x8a, 0x0c, 0x3c, 0x82, 0xba, 0xe3, 0x9f, 0x7d, 0x30, 0x66, 0x1b, 0x56, 0xd6,
	0xba, 0xf7, 0xa1, 0x1f, 0x3d, 0x64, 0xd1, 0xbd, 0x54, 0xad, 0x92, 0x2a, 0xc1, 0x6f, 0x40, 0x2f,
	0xe3, 0x84, 0x17, 0x25, 0x4b, 0x72, 0x29, 0x5c, 0x25, 0x4d, 0x41, 0x7c, 0xbc, 0x92, 0x7f, 0x2b,
	0xa5, 0x28, 0x9d, 0xc8, 0x58, 0xe8, 0xdd, 0xf2, 0x22, 0xcf, 0xd2, 0x82, 0x87, 0x65, 0x66, 0xa9,
	0xb2, 0x05, 0x75, 0x89, 0x66, 0xf8, 0x15, 0x0c, 0x79, 0x5a, 0x84, 0x29, 0x4b, 0xb8, 0xd5, 0x97,
	0x5d, 0x8d, 0xa7, 0x85, 0xcf, 0x12, 0x8e, 0x0f, 0x41, 0x93, 0xfe, 0x8a, 0xd7, 0xd6, 0x40, 0x76,
	0x06, 0x22, 0x75, 0xd7, 0xf8, 0x14, 0xcc, 0x27, 0xcf, 0x85, 0xf2, 0x84, 0x9a, 0x3c, 0xe1, 0xff,
	0xcd, 0x09, 0x9f, 0xb6, 0x90, 0x77, 0x33, 0x92, 0x26, 0xc1, 0xe7, 0x60, 0x46, 0x59, 0x5a, 0xf2,
	0xb4, 0xac, 0x98, 0x43, 0xc9, 0x7c, 0xd7, 0x30, 0x5b, 0x37, 0x98, 0xce, 0x2a, 0x64, 0x35, 0x25,
	0x6a, 0x12, 0x7c, 0x02, 0x5a, 0x51, 0x79, 0xd9, 0xd2, 0x47, 0xca, 0xc4, 0x38, 0xb6, 0x9a, 0x01,
	0x2f, 0x4d, 0x7e, 0xd1, 0x21, 0x35, 0x14, 0x4f, 0xa1, 0x1f, 0x0b, 0x9b, 0x59, 0x20, 0x39, 0x07,
	0x7f, 0x77, 0xe7, 0x45, 0x87, 0x54, 0x30, 0x81, 0x67, 0xe2, 0x4b, 0x5a, 0xc6, 0x9f, 0xf8, 0xb6,
	0x43, 0x04, 0x5e, 0xc2, 0xf0, 0x11, 0x98, 0xd9, 0xf6, 0x8e, 0xa5, 0x71, 0x21, 0xbd, 0x60, 0x99,
	0xc2, 0x6f, 0x17, 0x1d, 0xf2, 0xa2, 0x3a, 0xfe, 0xa5, 0x80, 0xd1, 0x5a, 0x0c, 0x5b, 0xb0, 0x5f,
	0x5b, 0x64, 0xb6, 0xf0, 0xa9, 0xe3, 0xd3, 0xda, 0x24, 0x7b, 0x00, 0xd4, 0xf9, 0x42, 0xc3, 0xe5,
	0x27, 0xdb, 0xf5, 0x91, 0x82, 0x0d, 0xd0, 0x02, 0xea, 0xce, 0xae, 0x1c, 0x82, 0xba, 0x18, 0x60,
	0x10, 0x50, 0x9b, 0xae, 0x02, 0xd4, 0xc3, 0x3a, 0xf4, 0x1d, 0x6f, 0x71, 0xe9, 0x22, 0x15, 0x1f,
	0xc2, 0x7f, 0x94, 0xd8, 0x7e, 0x60, 0xcf, 0xa8, 0xbb, 0x10, 0x13, 0x3d, 0xcf, 0xf6, 0xcf, 0x51,
	0x1f, 0x4f, 0xe0, 0x28, 0xb8, 0x09, 0xa8, 0xe3, 0x85, 0x9e, 0x13, 0x04, 0xe2, 0x5d, 0xd5, 0xff,
	0xb6, 0x24, 0xee, 0x67, 0x9b, 0x3a, 0xe1, 0x9c, 0x2c, 0x56, 0x4b, 0x34, 0x10, 0xd3, 0xe4, 0xc3,
	0x43, 0x9a, 0x08, 0xa5, 0x6d, 0xd1, 0x10, 0x23, 0x30, 0x17, 0x64, 0x6e, 0xfb, 0x6e, 0x60, 0x8b,
	0xc9, 0x48, 0x3f, 0xd3, 0x9f, 0x5f, 0xd6, 0xed, 0x40, 0x5e, 0xe6, 0xfd, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x16, 0x0f, 0x9f, 0x74, 0x8f, 0x04, 0x00, 0x00,
}
