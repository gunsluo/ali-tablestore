package ots

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

type CreateTableRequest struct {
	TableName    string
	PrimaryKey   []*ColumnSchema
	CapacityUnit CapacityUnit
}

func (this *CreateTableRequest) AddPrimaryKey(cloumn string, ct pb.ColumnType) *CreateTableRequest {

	cs := new(ColumnSchema)
	cs.Name = cloumn
	cs.Type = ct
	this.PrimaryKey = append(this.PrimaryKey, cs)

	return this
}

func (this *CreateTableRequest) PBStruct() *pb.CreateTableRequest {

	p := &pb.CreateTableRequest{
		TableMeta: &pb.TableMeta{
			TableName: proto.String(this.TableName),
		},
		ReservedThroughput: &pb.ReservedThroughput{
			CapacityUnit: this.CapacityUnit.PBStruct(),
		},
	}

	for _, item := range this.PrimaryKey {
		p.TableMeta.PrimaryKey = append(p.TableMeta.PrimaryKey, item.PBStruct())
	}

	return p
}
