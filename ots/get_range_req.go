package ots

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

type GetRangeRequest struct {
	TableName                string
	Direction                pb.Direction
	ColumnsToGet             []string
	Limit                    int32
	InclusiveStartPrimaryKey []*Column
	ExclusiveEndPrimaryKey   []*Column
}

func (this *GetRangeRequest) AddInclusiveStartPrimaryKey(name string, value ColumnValue) *GetRangeRequest {

	c := new(Column)
	c.Name = name
	c.Value = value
	this.InclusiveStartPrimaryKey = append(this.InclusiveStartPrimaryKey, c)

	return this
}

func (this *GetRangeRequest) AddExclusiveEndPrimaryKey(name string, value ColumnValue) *GetRangeRequest {

	c := new(Column)
	c.Name = name
	c.Value = value
	this.ExclusiveEndPrimaryKey = append(this.ExclusiveEndPrimaryKey, c)

	return this
}

func (this *GetRangeRequest) PBStruct() *pb.GetRangeRequest {

	p := &pb.GetRangeRequest{
		TableName:    proto.String(this.TableName),
		Direction:    &this.Direction,
		ColumnsToGet: this.ColumnsToGet,
		Limit:        proto.Int32(this.Limit),
	}

	for _, item := range this.InclusiveStartPrimaryKey {
		p.InclusiveStartPrimaryKey = append(p.InclusiveStartPrimaryKey, item.PBStruct())
	}

	for _, item := range this.ExclusiveEndPrimaryKey {
		p.ExclusiveEndPrimaryKey = append(p.ExclusiveEndPrimaryKey, item.PBStruct())
	}

	return p
}
