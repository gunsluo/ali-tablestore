package ots

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

type DeleteRowRequest struct {
	TableName    string
	RowExistence pb.RowExistenceExpectation
	PrimaryKey   []*Column
}

func (this *DeleteRowRequest) AddPrimaryKey(name string, value ColumnValue) *DeleteRowRequest {

	c := new(Column)
	c.Name = name
	c.Value = value
	this.PrimaryKey = append(this.PrimaryKey, c)

	return this
}

func (this *DeleteRowRequest) PBStruct() *pb.DeleteRowRequest {

	p := &pb.DeleteRowRequest{
		TableName: proto.String(this.TableName),
		Condition: &pb.Condition{
			RowExistence: &this.RowExistence,
		},
	}

	for _, item := range this.PrimaryKey {
		p.PrimaryKey = append(p.PrimaryKey, item.PBStruct())
	}

	return p
}
