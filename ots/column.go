package ots

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

type Column struct {
	Name  string
	Value ColumnValue
}

func (this *Column) PBStruct() *pb.Column {
	p := &pb.Column{
		Name:  proto.String(this.Name),
		Value: this.Value.PBStruct(),
	}

	return p
}

func NewColumn(p *pb.Column) *Column {
	resp := &Column{
		Name:  p.GetName(),
		Value: *(NewColumnValue(p.Value)),
	}

	return resp
}
