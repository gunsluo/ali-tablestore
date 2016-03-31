package ots

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

type UpdateRowRequest struct {
	TableName        string
	RowExistence     pb.RowExistenceExpectation
	PrimaryKey       []*Column
	AttributeColumns []*ColumnUpdate
}

func (this *UpdateRowRequest) AddPrimaryKey(name string, value ColumnValue) *UpdateRowRequest {

	c := new(Column)
	c.Name = name
	c.Value = value
	this.PrimaryKey = append(this.PrimaryKey, c)

	return this
}

func (this *UpdateRowRequest) AddAttributeColumns(name string, ty pb.OperationType, value ColumnValue) *UpdateRowRequest {

	cu := new(ColumnUpdate)
	cu.Name = name
	cu.Type = ty
	cu.Value = value
	this.AttributeColumns = append(this.AttributeColumns, cu)

	return this
}

func (this *UpdateRowRequest) PBStruct() *pb.UpdateRowRequest {

	p := &pb.UpdateRowRequest{
		TableName: proto.String(this.TableName),
		Condition: &pb.Condition{
			RowExistence: &this.RowExistence,
		},
	}

	for _, item := range this.PrimaryKey {
		p.PrimaryKey = append(p.PrimaryKey, item.PBStruct())
	}

	for _, item := range this.AttributeColumns {
		p.AttributeColumns = append(p.AttributeColumns, item.PBStruct())
	}

	return p
}
