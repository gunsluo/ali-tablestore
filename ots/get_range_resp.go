package ots

import "github.com/gunsluo/ali-tablestore/ots/proto/pb"

type GetRangeResponse struct {
	CapacityUnit        CapacityUnit
	NextStartPrimaryKey []*Column
	Rows                []*Row
}

func NewGetRangeResponse(p *pb.GetRangeResponse) *GetRangeResponse {

	resp := &GetRangeResponse{
		CapacityUnit: NewCapacityUnit(p.GetConsumed().GetCapacityUnit()),
	}

	nextStartPrimaryKey := p.GetNextStartPrimaryKey()
	if nextStartPrimaryKey != nil {
		for _, item := range nextStartPrimaryKey {
			resp.NextStartPrimaryKey = append(resp.NextStartPrimaryKey, NewColumn(item))
		}
	}

	rows := p.GetRows()
	if rows != nil {
		for _, item := range rows {
			resp.Rows = append(resp.Rows, NewRow(item))
		}
	}

	return resp
}

func (this *GetRangeResponse) IsDone() bool {
	if len(this.NextStartPrimaryKey) == 0 {
		return true
	}

	return false
}
