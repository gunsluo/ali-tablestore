package ots

import "github.com/gunsluo/ali-tablestore/ots/proto/pb"

type RowInBatchWriteRowResponse struct {
	IsOk         bool
	Error        Error
	CapacityUnit CapacityUnit
}

func NewRowInBatchWriteRowResponse(p *pb.RowInBatchWriteRowResponse) *RowInBatchWriteRowResponse {
	resp := &RowInBatchWriteRowResponse{
		IsOk:         p.GetIsOk(),
		Error:        *(NewError(p.Error)),
		CapacityUnit: NewCapacityUnit(p.GetConsumed().GetCapacityUnit()),
	}

	return resp
}

type TableInBatchWriteRowResponse struct {
	TableName  string
	PutRows    []*RowInBatchWriteRowResponse
	UpdateRows []*RowInBatchWriteRowResponse
	DeleteRows []*RowInBatchWriteRowResponse
}

func NewTableInBatchWriteRowResponse(p *pb.TableInBatchWriteRowResponse) *TableInBatchWriteRowResponse {

	resp := &TableInBatchWriteRowResponse{
		TableName: p.GetTableName(),
	}

	for _, item := range p.PutRows {
		resp.PutRows = append(resp.PutRows, NewRowInBatchWriteRowResponse(item))
	}

	for _, item := range p.UpdateRows {
		resp.UpdateRows = append(resp.UpdateRows, NewRowInBatchWriteRowResponse(item))
	}

	for _, item := range p.UpdateRows {
		resp.DeleteRows = append(resp.DeleteRows, NewRowInBatchWriteRowResponse(item))
	}

	return resp
}

type BatchWriteRowResponse struct {
	Tables []*TableInBatchWriteRowResponse
}

func NewBatchWriteRowResponse(p *pb.BatchWriteRowResponse) *BatchWriteRowResponse {

	resp := &BatchWriteRowResponse{}

	for _, item := range p.Tables {
		resp.Tables = append(resp.Tables, NewTableInBatchWriteRowResponse(item))
	}

	return resp
}
