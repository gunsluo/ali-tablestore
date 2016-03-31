package ots

import "github.com/gunsluo/ali-tablestore/ots/proto/pb"

type GetRowResponse struct {
	CapacityUnit CapacityUnit
	Row          Row
}

func NewGetRowResponse(p *pb.GetRowResponse) *GetRowResponse {

	resp := &GetRowResponse{
		CapacityUnit: NewCapacityUnit(p.GetConsumed().GetCapacityUnit()),
		Row:          *(NewRow(p.GetRow())),
	}

	return resp
}
