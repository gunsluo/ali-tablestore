package ots

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

type ColumnUpdate struct {
	Type  pb.OperationType
	Name  string
	Value ColumnValue
}

func (this *ColumnUpdate) PBStruct() *pb.ColumnUpdate {

	p := &pb.ColumnUpdate{
		Type:  &this.Type,
		Name:  proto.String(this.Name),
		Value: this.Value.PBStruct(),
	}

	return p
}
