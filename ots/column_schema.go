package ots

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

type ColumnSchema struct {
	Name string
	Type pb.ColumnType
}

func (this *ColumnSchema) PBStruct() *pb.ColumnSchema {
	p := &pb.ColumnSchema{
		Name: proto.String(this.Name),
		Type: &this.Type,
	}

	return p
}

func NewColumnSchema(p *pb.ColumnSchema) *ColumnSchema {
	resp := &ColumnSchema{
		Name: p.GetName(),
		Type: p.GetType(),
	}
	return resp
}
