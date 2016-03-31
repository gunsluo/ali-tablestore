package ots

import "github.com/gunsluo/ali-tablestore/ots/proto/pb"

type DescribeTableResponse struct {
	TableName              string
	PrimaryKey             []*ColumnSchema
	CapacityUnit           CapacityUnit
	LastIncreaseTime       int64
	LastDecreaseTime       int64
	NumberOfDecreasesToday int32
}

func NewDescribeTableResponse(p *pb.DescribeTableResponse) *DescribeTableResponse {

	resp := &DescribeTableResponse{
		TableName:              p.GetTableMeta().GetTableName(),
		CapacityUnit:           NewCapacityUnit(p.GetReservedThroughputDetails().GetCapacityUnit()),
		LastIncreaseTime:       p.GetReservedThroughputDetails().GetLastIncreaseTime(),
		LastDecreaseTime:       p.GetReservedThroughputDetails().GetLastDecreaseTime(),
		NumberOfDecreasesToday: p.GetReservedThroughputDetails().GetNumberOfDecreasesToday(),
	}

	all := p.GetTableMeta().GetPrimaryKey()
	for _, item := range all {
		resp.PrimaryKey = append(resp.PrimaryKey, NewColumnSchema(item))
	}

	return resp
}
