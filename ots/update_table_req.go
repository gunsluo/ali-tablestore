package ots

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

type UpdateTableRequest struct {
	TableName    string
	CapacityUnit CapacityUnit
}

func (this *UpdateTableRequest) PBStruct() *pb.UpdateTableRequest {

	p := &pb.UpdateTableRequest{
		TableName: proto.String(this.TableName),
		ReservedThroughput: &pb.ReservedThroughput{
			CapacityUnit: this.CapacityUnit.PBStruct(),
		},
	}

	return p
}
