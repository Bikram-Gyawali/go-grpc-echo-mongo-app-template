// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: book.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetBookRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookId        string                 `protobuf:"bytes,1,opt,name=bookId,proto3" json:"bookId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookRequest) Reset() {
	*x = GetBookRequest{}
	mi := &file_book_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookRequest) ProtoMessage() {}

func (x *GetBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookRequest.ProtoReflect.Descriptor instead.
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{0}
}

func (x *GetBookRequest) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

type GetBookResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookId        string                 `protobuf:"bytes,1,opt,name=bookId,proto3" json:"bookId,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Author        string                 `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	PublishedDate string                 `protobuf:"bytes,5,opt,name=publishedDate,proto3" json:"publishedDate,omitempty"`
	Image         string                 `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookResponse) Reset() {
	*x = GetBookResponse{}
	mi := &file_book_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookResponse) ProtoMessage() {}

func (x *GetBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookResponse.ProtoReflect.Descriptor instead.
func (*GetBookResponse) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{1}
}

func (x *GetBookResponse) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *GetBookResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetBookResponse) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *GetBookResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetBookResponse) GetPublishedDate() string {
	if x != nil {
		return x.PublishedDate
	}
	return ""
}

func (x *GetBookResponse) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type CreateBookRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Author        string                 `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	PublishedDate string                 `protobuf:"bytes,4,opt,name=publishedDate,proto3" json:"publishedDate,omitempty"`
	Image         string                 `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	mi := &file_book_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookRequest.ProtoReflect.Descriptor instead.
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBookRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBookRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *CreateBookRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateBookRequest) GetPublishedDate() string {
	if x != nil {
		return x.PublishedDate
	}
	return ""
}

func (x *CreateBookRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type CreateBookResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookId        string                 `protobuf:"bytes,1,opt,name=bookId,proto3" json:"bookId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBookResponse) Reset() {
	*x = CreateBookResponse{}
	mi := &file_book_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookResponse) ProtoMessage() {}

func (x *CreateBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookResponse.ProtoReflect.Descriptor instead.
func (*CreateBookResponse) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBookResponse) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

type UpdateBookRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookId        string                 `protobuf:"bytes,1,opt,name=bookId,proto3" json:"bookId,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Author        string                 `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	PublishedDate string                 `protobuf:"bytes,5,opt,name=publishedDate,proto3" json:"publishedDate,omitempty"`
	Image         string                 `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBookRequest) Reset() {
	*x = UpdateBookRequest{}
	mi := &file_book_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookRequest) ProtoMessage() {}

func (x *UpdateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateBookRequest) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *UpdateBookRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBookRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *UpdateBookRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateBookRequest) GetPublishedDate() string {
	if x != nil {
		return x.PublishedDate
	}
	return ""
}

func (x *UpdateBookRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type UpdateBookResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookId        string                 `protobuf:"bytes,1,opt,name=bookId,proto3" json:"bookId,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Author        string                 `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	PublishedDate string                 `protobuf:"bytes,5,opt,name=publishedDate,proto3" json:"publishedDate,omitempty"`
	Image         string                 `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBookResponse) Reset() {
	*x = UpdateBookResponse{}
	mi := &file_book_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookResponse) ProtoMessage() {}

func (x *UpdateBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookResponse.ProtoReflect.Descriptor instead.
func (*UpdateBookResponse) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBookResponse) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *UpdateBookResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBookResponse) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *UpdateBookResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateBookResponse) GetPublishedDate() string {
	if x != nil {
		return x.PublishedDate
	}
	return ""
}

func (x *UpdateBookResponse) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type DeleteBookRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookId        string                 `protobuf:"bytes,1,opt,name=bookId,proto3" json:"bookId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteBookRequest) Reset() {
	*x = DeleteBookRequest{}
	mi := &file_book_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookRequest) ProtoMessage() {}

func (x *DeleteBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookRequest.ProtoReflect.Descriptor instead.
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteBookRequest) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

type DeleteBookResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookId        string                 `protobuf:"bytes,1,opt,name=bookId,proto3" json:"bookId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteBookResponse) Reset() {
	*x = DeleteBookResponse{}
	mi := &file_book_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookResponse) ProtoMessage() {}

func (x *DeleteBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookResponse.ProtoReflect.Descriptor instead.
func (*DeleteBookResponse) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBookResponse) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

var File_book_proto protoreflect.FileDescriptor

var file_book_proto_rawDesc = string([]byte{
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70,
	0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x28, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0xb3, 0x01, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62,
	0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22,
	0x9d, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22,
	0x2c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0xb5, 0x01,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f,
	0x6f, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x2b,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x32, 0xfd, 0x02, 0x0a, 0x0b, 0x42, 0x6f,
	0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x7b, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x7d, 0x12,
	0x57, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x16, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x60, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a,
	0x01, 0x2a, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x2f, 0x7b, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x7d, 0x12, 0x5d, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x18, 0x2a, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x2f, 0x7b, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x7d, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x69, 0x6b, 0x72, 0x61, 0x6d, 0x2d, 0x47,
	0x79, 0x61, 0x77, 0x61, 0x6c, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x68, 0x6f,
	0x2d, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_book_proto_rawDescOnce sync.Once
	file_book_proto_rawDescData []byte
)

func file_book_proto_rawDescGZIP() []byte {
	file_book_proto_rawDescOnce.Do(func() {
		file_book_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_book_proto_rawDesc), len(file_book_proto_rawDesc)))
	})
	return file_book_proto_rawDescData
}

var file_book_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_book_proto_goTypes = []any{
	(*GetBookRequest)(nil),     // 0: api.GetBookRequest
	(*GetBookResponse)(nil),    // 1: api.GetBookResponse
	(*CreateBookRequest)(nil),  // 2: api.CreateBookRequest
	(*CreateBookResponse)(nil), // 3: api.CreateBookResponse
	(*UpdateBookRequest)(nil),  // 4: api.UpdateBookRequest
	(*UpdateBookResponse)(nil), // 5: api.UpdateBookResponse
	(*DeleteBookRequest)(nil),  // 6: api.DeleteBookRequest
	(*DeleteBookResponse)(nil), // 7: api.DeleteBookResponse
}
var file_book_proto_depIdxs = []int32{
	0, // 0: api.BookService.GetBook:input_type -> api.GetBookRequest
	2, // 1: api.BookService.CreateBook:input_type -> api.CreateBookRequest
	4, // 2: api.BookService.UpdateBook:input_type -> api.UpdateBookRequest
	6, // 3: api.BookService.DeleteBook:input_type -> api.DeleteBookRequest
	1, // 4: api.BookService.GetBook:output_type -> api.GetBookResponse
	3, // 5: api.BookService.CreateBook:output_type -> api.CreateBookResponse
	5, // 6: api.BookService.UpdateBook:output_type -> api.UpdateBookResponse
	7, // 7: api.BookService.DeleteBook:output_type -> api.DeleteBookResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_book_proto_init() }
func file_book_proto_init() {
	if File_book_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_book_proto_rawDesc), len(file_book_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_book_proto_goTypes,
		DependencyIndexes: file_book_proto_depIdxs,
		MessageInfos:      file_book_proto_msgTypes,
	}.Build()
	File_book_proto = out.File
	file_book_proto_goTypes = nil
	file_book_proto_depIdxs = nil
}
