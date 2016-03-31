package ots

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

type PutRowRequest struct {
	TableName        string
	RowExistence     pb.RowExistenceExpectation
	PrimaryKey       []*Column
	AttributeColumns []*Column
}

func (this *PutRowRequest) AddPrimaryKey(name string, value ColumnValue) *PutRowRequest {

	c := new(Column)
	c.Name = name
	c.Value = value
	this.PrimaryKey = append(this.PrimaryKey, c)

	return this
}

func (this *PutRowRequest) AddAttributeColumns(name string, value ColumnValue) *PutRowRequest {

	c := new(Column)
	c.Name = name
	c.Value = value
	this.AttributeColumns = append(this.AttributeColumns, c)

	return this
}

func (this *PutRowRequest) PBStruct() *pb.PutRowRequest {

	p := &pb.PutRowRequest{
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
