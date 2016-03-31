package ots

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

type DescribeTableRequest struct {
	TableName string
}

func (this *DescribeTableRequest) PBStruct() *pb.DescribeTableRequest {

	p := &pb.DescribeTableRequest{
		TableName: proto.String(this.TableName),
	}

	return p
}
