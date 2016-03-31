package ots

import "github.com/gunsluo/ali-tablestore/ots/proto/pb"

type RowInBatchGetRowResponse struct {
	IsOk         bool
	Error        Error
	CapacityUnit CapacityUnit
	Row          Row
}

func NewRowInBatchGetRowResponse(p *pb.RowInBatchGetRowResponse) *RowInBatchGetRowResponse {
	resp := &RowInBatchGetRowResponse{
		IsOk:         p.GetIsOk(),
		Error:        *(NewError(p.Error)),
		CapacityUnit: NewCapacityUnit(p.GetConsumed().GetCapacityUnit()),
		Row:          *(NewRow(p.Row)),
	}

	return resp
}

type TableInBatchGetRowResponse struct {
	TableName string
	Rows      []*RowInBatchGetRowResponse
}

func NewTableInBatchGetRowResponse(p *pb.TableInBatchGetRowResponse) *TableInBatchGetRowResponse {

	resp := &TableInBatchGetRowResponse{
		TableName: p.GetTableName(),
	}

	for _, item := range p.Rows {
		resp.Rows = append(resp.Rows, NewRowInBatchGetRowResponse(item))
	}

	return resp
}

type BatchGetRowResponse struct {
	Tables []*TableInBatchGetRowResponse
}

func NewBatchGetRowResponse(p *pb.BatchGetRowResponse) *BatchGetRowResponse {

	resp := &BatchGetRowResponse{}

	for _, item := range p.Tables {
		resp.Tables = append(resp.Tables, NewTableInBatchGetRowResponse(item))
	}

	return resp
}
