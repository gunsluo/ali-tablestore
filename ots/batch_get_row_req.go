package ots

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

type RowInBatchGetRowRequest struct {
	PrimaryKey []*Column
}

func (this *RowInBatchGetRowRequest) AddPrimaryKey(name string, value ColumnValue) *RowInBatchGetRowRequest {

	c := new(Column)
	c.Name = name
	c.Value = value
	this.PrimaryKey = append(this.PrimaryKey, c)

	return this
}

func (this *RowInBatchGetRowRequest) PBStruct() *pb.RowInBatchGetRowRequest {

	p := &pb.RowInBatchGetRowRequest{}

	for _, item := range this.PrimaryKey {
		p.PrimaryKey = append(p.PrimaryKey, item.PBStruct())
	}

	return p

}

type TableInBatchGetRowRequest struct {
	TableName    string
	Rows         []*RowInBatchGetRowRequest
	ColumnsToGet []string
}

func (this *TableInBatchGetRowRequest) AddRow(row *RowInBatchGetRowRequest) *TableInBatchGetRowRequest {

	this.Rows = append(this.Rows, row)
	return this
}

func (this *TableInBatchGetRowRequest) PBStruct() *pb.TableInBatchGetRowRequest {

	p := &pb.TableInBatchGetRowRequest{
		TableName:    proto.String(this.TableName),
		ColumnsToGet: this.ColumnsToGet,
	}

	for _, item := range this.Rows {
		p.Rows = append(p.Rows, item.PBStruct())
	}

	return p
}

type BatchGetRowRequest struct {
	Tables []*TableInBatchGetRowRequest
}

func (this *BatchGetRowRequest) AddTable(table *TableInBatchGetRowRequest) *BatchGetRowRequest {

	this.Tables = append(this.Tables, table)
	return this
}

func (this *BatchGetRowRequest) PBStruct() *pb.BatchGetRowRequest {

	p := &pb.BatchGetRowRequest{}

	for _, item := range this.Tables {
		p.Tables = append(p.Tables, item.PBStruct())
	}

	return p
}
