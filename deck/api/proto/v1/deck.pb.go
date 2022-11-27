// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: deck.proto

package v1

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

// The request message containing the user's name.
type GetDeckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeckId string `protobuf:"bytes,1,opt,name=deck_id,json=deckId,proto3" json:"deck_id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetDeckRequest) Reset() {
	*x = GetDeckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeckRequest) ProtoMessage() {}

func (x *GetDeckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeckRequest.ProtoReflect.Descriptor instead.
func (*GetDeckRequest) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{0}
}

func (x *GetDeckRequest) GetDeckId() string {
	if x != nil {
		return x.DeckId
	}
	return ""
}

func (x *GetDeckRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetDeckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deck *Deck `protobuf:"bytes,1,opt,name=deck,proto3" json:"deck,omitempty"`
}

func (x *GetDeckResponse) Reset() {
	*x = GetDeckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeckResponse) ProtoMessage() {}

func (x *GetDeckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeckResponse.ProtoReflect.Descriptor instead.
func (*GetDeckResponse) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{1}
}

func (x *GetDeckResponse) GetDeck() *Deck {
	if x != nil {
		return x.Deck
	}
	return nil
}

type GetDecksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetDecksRequest) Reset() {
	*x = GetDecksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDecksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDecksRequest) ProtoMessage() {}

func (x *GetDecksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDecksRequest.ProtoReflect.Descriptor instead.
func (*GetDecksRequest) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{2}
}

func (x *GetDecksRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetDecksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Decks []*Deck `protobuf:"bytes,1,rep,name=decks,proto3" json:"decks,omitempty"`
}

func (x *GetDecksResponse) Reset() {
	*x = GetDecksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDecksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDecksResponse) ProtoMessage() {}

func (x *GetDecksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDecksResponse.ProtoReflect.Descriptor instead.
func (*GetDecksResponse) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{3}
}

func (x *GetDecksResponse) GetDecks() []*Deck {
	if x != nil {
		return x.Decks
	}
	return nil
}

type CreateDeckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deck *Deck `protobuf:"bytes,1,opt,name=deck,proto3" json:"deck,omitempty"`
}

func (x *CreateDeckRequest) Reset() {
	*x = CreateDeckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeckRequest) ProtoMessage() {}

func (x *CreateDeckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeckRequest.ProtoReflect.Descriptor instead.
func (*CreateDeckRequest) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{4}
}

func (x *CreateDeckRequest) GetDeck() *Deck {
	if x != nil {
		return x.Deck
	}
	return nil
}

type CreateDeckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deck *Deck `protobuf:"bytes,1,opt,name=deck,proto3" json:"deck,omitempty"`
}

func (x *CreateDeckResponse) Reset() {
	*x = CreateDeckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeckResponse) ProtoMessage() {}

func (x *CreateDeckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeckResponse.ProtoReflect.Descriptor instead.
func (*CreateDeckResponse) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{5}
}

func (x *CreateDeckResponse) GetDeck() *Deck {
	if x != nil {
		return x.Deck
	}
	return nil
}

type DeleteDeckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDeckRequest) Reset() {
	*x = DeleteDeckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeckRequest) ProtoMessage() {}

func (x *DeleteDeckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeckRequest.ProtoReflect.Descriptor instead.
func (*DeleteDeckRequest) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteDeckRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteDeckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteDeckResponse) Reset() {
	*x = DeleteDeckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeckResponse) ProtoMessage() {}

func (x *DeleteDeckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeckResponse.ProtoReflect.Descriptor instead.
func (*DeleteDeckResponse) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{7}
}

type Deck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId    string  `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title       string  `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Cards       []*Card `protobuf:"bytes,5,rep,name=cards,proto3" json:"cards,omitempty"`
}

func (x *Deck) Reset() {
	*x = Deck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deck) ProtoMessage() {}

func (x *Deck) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deck.ProtoReflect.Descriptor instead.
func (*Deck) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{8}
}

func (x *Deck) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Deck) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Deck) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Deck) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Deck) GetCards() []*Card {
	if x != nil {
		return x.Cards
	}
	return nil
}

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeckId          string    `protobuf:"bytes,2,opt,name=deck_id,json=deckId,proto3" json:"deck_id,omitempty"`
	Title           string    `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	PossibleAnswers []*Answer `protobuf:"bytes,4,rep,name=possible_answers,json=possibleAnswers,proto3" json:"possible_answers,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{9}
}

func (x *Card) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Card) GetDeckId() string {
	if x != nil {
		return x.DeckId
	}
	return ""
}

func (x *Card) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Card) GetPossibleAnswers() []*Answer {
	if x != nil {
		return x.PossibleAnswers
	}
	return nil
}

type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CardId    string `protobuf:"bytes,2,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`
	Text      string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	IsCorrect bool   `protobuf:"varint,4,opt,name=is_correct,json=isCorrect,proto3" json:"is_correct,omitempty"`
}

func (x *Answer) Reset() {
	*x = Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deck_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_deck_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_deck_proto_rawDescGZIP(), []int{10}
}

func (x *Answer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Answer) GetCardId() string {
	if x != nil {
		return x.CardId
	}
	return ""
}

func (x *Answer) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Answer) GetIsCorrect() bool {
	if x != nil {
		return x.IsCorrect
	}
	return false
}

var File_deck_proto protoreflect.FileDescriptor

var file_deck_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31,
	0x22, 0x42, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x64, 0x65, 0x63, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x6b, 0x52,
	0x04, 0x64, 0x65, 0x63, 0x6b, 0x22, 0x2a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x32, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x05, 0x64, 0x65, 0x63, 0x6b, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x05,
	0x64, 0x65, 0x63, 0x6b, 0x73, 0x22, 0x31, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x64, 0x65,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x63, 0x6b, 0x52, 0x04, 0x64, 0x65, 0x63, 0x6b, 0x22, 0x32, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x04, 0x64, 0x65, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x04, 0x64, 0x65, 0x63, 0x6b, 0x22, 0x23, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x04, 0x44, 0x65, 0x63, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x05,
	0x63, 0x61, 0x72, 0x64, 0x73, 0x22, 0x7c, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x64, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x10,
	0x70, 0x6f, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x52, 0x0f, 0x70, 0x6f, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x73, 0x22, 0x64, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73,
	0x5f, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x32, 0xf7, 0x01, 0x0a, 0x08, 0x44, 0x65,
	0x63, 0x6b, 0x73, 0x41, 0x50, 0x49, 0x12, 0x34, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63,
	0x6b, 0x12, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x63, 0x6b, 0x73, 0x12, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x63, 0x6b, 0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65,
	0x63, 0x6b, 0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x58, 0x61, 0x76, 0x69, 0x46, 0x50, 0x2f, 0x74, 0x6f, 0x73, 0x68, 0x6f, 0x6b, 0x61,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x63, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deck_proto_rawDescOnce sync.Once
	file_deck_proto_rawDescData = file_deck_proto_rawDesc
)

func file_deck_proto_rawDescGZIP() []byte {
	file_deck_proto_rawDescOnce.Do(func() {
		file_deck_proto_rawDescData = protoimpl.X.CompressGZIP(file_deck_proto_rawDescData)
	})
	return file_deck_proto_rawDescData
}

var file_deck_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_deck_proto_goTypes = []interface{}{
	(*GetDeckRequest)(nil),     // 0: v1.GetDeckRequest
	(*GetDeckResponse)(nil),    // 1: v1.GetDeckResponse
	(*GetDecksRequest)(nil),    // 2: v1.GetDecksRequest
	(*GetDecksResponse)(nil),   // 3: v1.GetDecksResponse
	(*CreateDeckRequest)(nil),  // 4: v1.CreateDeckRequest
	(*CreateDeckResponse)(nil), // 5: v1.CreateDeckResponse
	(*DeleteDeckRequest)(nil),  // 6: v1.DeleteDeckRequest
	(*DeleteDeckResponse)(nil), // 7: v1.DeleteDeckResponse
	(*Deck)(nil),               // 8: v1.Deck
	(*Card)(nil),               // 9: v1.Card
	(*Answer)(nil),             // 10: v1.Answer
}
var file_deck_proto_depIdxs = []int32{
	8,  // 0: v1.GetDeckResponse.deck:type_name -> v1.Deck
	8,  // 1: v1.GetDecksResponse.decks:type_name -> v1.Deck
	8,  // 2: v1.CreateDeckRequest.deck:type_name -> v1.Deck
	8,  // 3: v1.CreateDeckResponse.deck:type_name -> v1.Deck
	9,  // 4: v1.Deck.cards:type_name -> v1.Card
	10, // 5: v1.Card.possible_answers:type_name -> v1.Answer
	0,  // 6: v1.DecksAPI.GetDeck:input_type -> v1.GetDeckRequest
	2,  // 7: v1.DecksAPI.GetDecks:input_type -> v1.GetDecksRequest
	4,  // 8: v1.DecksAPI.CreateDeck:input_type -> v1.CreateDeckRequest
	6,  // 9: v1.DecksAPI.DeleteDeck:input_type -> v1.DeleteDeckRequest
	1,  // 10: v1.DecksAPI.GetDeck:output_type -> v1.GetDeckResponse
	3,  // 11: v1.DecksAPI.GetDecks:output_type -> v1.GetDecksResponse
	5,  // 12: v1.DecksAPI.CreateDeck:output_type -> v1.CreateDeckResponse
	7,  // 13: v1.DecksAPI.DeleteDeck:output_type -> v1.DeleteDeckResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_deck_proto_init() }
func file_deck_proto_init() {
	if File_deck_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deck_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeckRequest); i {
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
		file_deck_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeckResponse); i {
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
		file_deck_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDecksRequest); i {
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
		file_deck_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDecksResponse); i {
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
		file_deck_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeckRequest); i {
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
		file_deck_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeckResponse); i {
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
		file_deck_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeckRequest); i {
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
		file_deck_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeckResponse); i {
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
		file_deck_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deck); i {
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
		file_deck_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
		file_deck_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer); i {
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
			RawDescriptor: file_deck_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deck_proto_goTypes,
		DependencyIndexes: file_deck_proto_depIdxs,
		MessageInfos:      file_deck_proto_msgTypes,
	}.Build()
	File_deck_proto = out.File
	file_deck_proto_rawDesc = nil
	file_deck_proto_goTypes = nil
	file_deck_proto_depIdxs = nil
}
