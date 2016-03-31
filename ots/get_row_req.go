package ots

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

type GetRowRequest struct {
	TableName    string
	PrimaryKey   []*Column
	ColumnsToGet []string
}

func (this *GetRowRequest) AddPrimaryKey(name string, value ColumnValue) *GetRowRequest {

	c := new(Column)
	c.Name = name
	c.Value = value
	this.PrimaryKey = append(this.PrimaryKey, c)

	return this
}

func (this *GetRowRequest) PBStruct() *pb.GetRowRequest {

	p := &pb.GetRowRequest{
		TableName:    proto.String(this.TableName),
		ColumnsToGet: this.ColumnsToGet,
	}

	for _, item := range this.PrimaryKey {
		p.PrimaryKey = append(p.PrimaryKey, item.PBStruct())
	}

	return p
}
