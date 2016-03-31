package ots

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

type DeleteTableRequest struct {
	TableName string
}

func (this *DeleteTableRequest) PBStruct() *pb.DeleteTableRequest {

	p := &pb.DeleteTableRequest{
		TableName: proto.String(this.TableName),
	}

	return p
}
