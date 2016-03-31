package ots

import "github.com/gunsluo/ali-tablestore/ots/proto/pb"

type UpdateTableResponse struct {
	CapacityUnit           CapacityUnit
	LastIncreaseTime       int64
	LastDecreaseTime       int64
	NumberOfDecreasesToday int32
}

func NewUpdateTableResponse(p *pb.UpdateTableResponse) *UpdateTableResponse {

	resp := &UpdateTableResponse{
		CapacityUnit:           NewCapacityUnit(p.GetReservedThroughputDetails().GetCapacityUnit()),
		LastIncreaseTime:       p.GetReservedThroughputDetails().GetLastIncreaseTime(),
		LastDecreaseTime:       p.GetReservedThroughputDetails().GetLastDecreaseTime(),
		NumberOfDecreasesToday: p.GetReservedThroughputDetails().GetNumberOfDecreasesToday(),
	}

	return resp
}
