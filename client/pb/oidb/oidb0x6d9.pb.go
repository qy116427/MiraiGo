// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/oidb/oidb0x6d9.proto

package oidb

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type GroupFileFeedsInfo struct {
	BusId     proto.Option[uint32] `protobuf:"varint,1,opt"`
	FileId    proto.Option[string] `protobuf:"bytes,2,opt"`
	MsgRandom proto.Option[uint32] `protobuf:"varint,3,opt"`
	Ext       []byte               `protobuf:"bytes,4,opt"`
	FeedFlag  proto.Option[uint32] `protobuf:"varint,5,opt"`
}

type CopyFromReqBody struct {
	GroupCode       proto.Option[uint64] `protobuf:"varint,1,opt"`
	AppId           proto.Option[uint32] `protobuf:"varint,2,opt"`
	SrcBusId        proto.Option[uint32] `protobuf:"varint,3,opt"`
	SrcParentFolder []byte               `protobuf:"bytes,4,opt"`
	SrcFilePath     []byte               `protobuf:"bytes,5,opt"`
	DstBusId        proto.Option[uint32] `protobuf:"varint,6,opt"`
	DstFolderId     []byte               `protobuf:"bytes,7,opt"`
	FileSize        proto.Option[uint64] `protobuf:"varint,8,opt"`
	LocalPath       proto.Option[string] `protobuf:"bytes,9,opt"`
	FileName        proto.Option[string] `protobuf:"bytes,10,opt"`
	SrcUin          proto.Option[uint64] `protobuf:"varint,11,opt"`
	Md5             []byte               `protobuf:"bytes,12,opt"`
}

type CopyFromRspBody struct {
	RetCode       proto.Option[int32]  `protobuf:"varint,1,opt"`
	RetMsg        proto.Option[string] `protobuf:"bytes,2,opt"`
	ClientWording proto.Option[string] `protobuf:"bytes,3,opt"`
	SaveFilePath  []byte               `protobuf:"bytes,4,opt"`
	BusId         proto.Option[uint32] `protobuf:"varint,5,opt"`
}

type CopyToReqBody struct {
	GroupCode             proto.Option[uint64] `protobuf:"varint,1,opt"`
	AppId                 proto.Option[uint32] `protobuf:"varint,2,opt"`
	SrcBusId              proto.Option[uint32] `protobuf:"varint,3,opt"`
	SrcFileId             proto.Option[string] `protobuf:"bytes,4,opt"`
	DstBusId              proto.Option[uint32] `protobuf:"varint,5,opt"`
	DstUin                proto.Option[uint64] `protobuf:"varint,6,opt"`
	NewFileName           proto.Option[string] `protobuf:"bytes,40,opt"`
	TimCloudPdirKey       []byte               `protobuf:"bytes,100,opt"`
	TimCloudPpdirKey      []byte               `protobuf:"bytes,101,opt"`
	TimCloudExtensionInfo []byte               `protobuf:"bytes,102,opt"`
	TimFileExistOption    proto.Option[uint32] `protobuf:"varint,103,opt"`
}

type CopyToRspBody struct {
	RetCode       proto.Option[int32]  `protobuf:"varint,1,opt"`
	RetMsg        proto.Option[string] `protobuf:"bytes,2,opt"`
	ClientWording proto.Option[string] `protobuf:"bytes,3,opt"`
	SaveFilePath  proto.Option[string] `protobuf:"bytes,4,opt"`
	BusId         proto.Option[uint32] `protobuf:"varint,5,opt"`
	FileName      proto.Option[string] `protobuf:"bytes,40,opt"`
}

type FeedsReqBody struct {
	GroupCode     proto.Option[uint64]  `protobuf:"varint,1,opt"`
	AppId         proto.Option[uint32]  `protobuf:"varint,2,opt"`
	FeedsInfoList []*GroupFileFeedsInfo `protobuf:"bytes,3,rep"`
	MultiSendSeq  proto.Option[uint32]  `protobuf:"varint,4,opt"`
}

type FeedsRspBody struct {
	RetCode       proto.Option[int32]  `protobuf:"varint,1,opt"`
	RetMsg        proto.Option[string] `protobuf:"bytes,2,opt"`
	ClientWording proto.Option[string] `protobuf:"bytes,3,opt"`
	//repeated C8639group_file_common.FeedsResult feedsResultList = 4;
	SvrbusyWaitTime proto.Option[uint32] `protobuf:"varint,5,opt"`
}

type D6D9ReqBody struct {
	TransFileReq *TransFileReqBody `protobuf:"bytes,1,opt"`
	CopyFromReq  *CopyFromReqBody  `protobuf:"bytes,2,opt"`
	CopyToReq    *CopyToReqBody    `protobuf:"bytes,3,opt"`
	FeedsInfoReq *FeedsReqBody     `protobuf:"bytes,5,opt"`
}

type D6D9RspBody struct {
	TransFileRsp *TransFileRspBody `protobuf:"bytes,1,opt"`
	CopyFromRsp  *CopyFromRspBody  `protobuf:"bytes,2,opt"`
	CopyToRsp    *CopyToRspBody    `protobuf:"bytes,3,opt"`
	FeedsInfoRsp *FeedsRspBody     `protobuf:"bytes,5,opt"`
}

type TransFileReqBody struct {
	GroupCode proto.Option[uint64] `protobuf:"varint,1,opt"`
	AppId     proto.Option[uint32] `protobuf:"varint,2,opt"`
	BusId     proto.Option[uint32] `protobuf:"varint,3,opt"`
	FileId    proto.Option[string] `protobuf:"bytes,4,opt"`
}

type TransFileRspBody struct {
	RetCode       proto.Option[int32]  `protobuf:"varint,1,opt"`
	RetMsg        proto.Option[string] `protobuf:"bytes,2,opt"`
	ClientWording proto.Option[string] `protobuf:"bytes,3,opt"`
	SaveBusId     proto.Option[uint32] `protobuf:"varint,4,opt"`
	SaveFilePath  proto.Option[string] `protobuf:"bytes,5,opt"`
}
