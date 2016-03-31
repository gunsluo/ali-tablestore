package ots

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

type PutRowInBatchWriteRowRequest struct {
	RowExistence     pb.RowExistenceExpectation
	PrimaryKey       []*Column
	AttributeColumns []*Column
}

func (this *PutRowInBatchWriteRowRequest) AddPrimaryKey(name string, value ColumnValue) *PutRowInBatchWriteRowRequest {

	c := new(Column)
	c.Name = name
	c.Value = value
	this.PrimaryKey = append(this.PrimaryKey, c)

	return this
}

func (this *PutRowInBatchWriteRowRequest) AddAttributeColumns(name string, value ColumnValue) *PutRowInBatchWriteRowRequest {

	c := new(Column)
	c.Name = name
	c.Value = value
	this.AttributeColumns = append(this.AttributeColumns, c)

	return this
}

func (this *PutRowInBatchWriteRowRequest) PBStruct() *pb.PutRowInBatchWriteRowRequest {
	p := &pb.PutRowInBatchWriteRowRequest{
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

type UpdateRowInBatchWriteRowRequest struct {
	RowExistence     pb.RowExistenceExpectation
	PrimaryKey       []*Column
	AttributeColumns []*ColumnUpdate
}

func (this *UpdateRowInBatchWriteRowRequest) AddPrimaryKey(name string, value ColumnValue) *UpdateRowInBatchWriteRowRequest {

	c := new(Column)
	c.Name = name
	c.Value = value
	this.PrimaryKey = append(this.PrimaryKey, c)

	return this
}

func (this *UpdateRowInBatchWriteRowRequest) AddAttributeColumns(name string, ty pb.OperationType, value ColumnValue) *UpdateRowInBatchWriteRowRequest {

	cu := new(ColumnUpdate)
	cu.Name = name
	cu.Type = ty
	cu.Value = value
	this.AttributeColumns = append(this.AttributeColumns, cu)

	return this
}

func (this *UpdateRowInBatchWriteRowRequest) PBStruct() *pb.UpdateRowInBatchWriteRowRequest {
	p := &pb.UpdateRowInBatchWriteRowRequest{
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

type DeleteRowInBatchWriteRowRequest struct {
	RowExistence pb.RowExistenceExpectation
	PrimaryKey   []*Column
}

func (this *DeleteRowInBatchWriteRowRequest) AddPrimaryKey(name string, value ColumnValue) *DeleteRowInBatchWriteRowRequest {

	c := new(Column)
	c.Name = name
	c.Value = value
	this.PrimaryKey = append(this.PrimaryKey, c)

	return this
}

func (this *DeleteRowInBatchWriteRowRequest) PBStruct() *pb.DeleteRowInBatchWriteRowRequest {
	p := &pb.DeleteRowInBatchWriteRowRequest{
		Condition: &pb.Condition{
			RowExistence: &this.RowExistence,
		},
	}

	for _, item := range this.PrimaryKey {
		p.PrimaryKey = append(p.PrimaryKey, item.PBStruct())
	}

	return p
}

type TableInBatchWriteRowRequest struct {
	TableName  string
	PutRows    []*PutRowInBatchWriteRowRequest
	UpdateRows []*UpdateRowInBatchWriteRowRequest
	DeleteRows []*DeleteRowInBatchWriteRowRequest
}

func (this *TableInBatchWriteRowRequest) AddPutRow(row *PutRowInBatchWriteRowRequest) *TableInBatchWriteRowRequest {

	this.PutRows = append(this.PutRows, row)
	return this
}

func (this *TableInBatchWriteRowRequest) AddUpdateRow(row *UpdateRowInBatchWriteRowRequest) *TableInBatchWriteRowRequest {

	this.UpdateRows = append(this.UpdateRows, row)
	return this
}

func (this *TableInBatchWriteRowRequest) AddDeleteRow(row *DeleteRowInBatchWriteRowRequest) *TableInBatchWriteRowRequest {

	this.DeleteRows = append(this.DeleteRows, row)
	return this
}

func (this *TableInBatchWriteRowRequest) PBStruct() *pb.TableInBatchWriteRowRequest {

	p := &pb.TableInBatchWriteRowRequest{
		TableName: proto.String(this.TableName),
	}

	for _, item := range this.PutRows {
		p.PutRows = append(p.PutRows, item.PBStruct())
	}

	for _, item := range this.UpdateRows {
		p.UpdateRows = append(p.UpdateRows, item.PBStruct())
	}

	for _, item := range this.DeleteRows {
		p.DeleteRows = append(p.DeleteRows, item.PBStruct())
	}

	return p
}

type BatchWriteRowRequest struct {
	Tables []*TableInBatchWriteRowRequest
}

func (this *BatchWriteRowRequest) AddTable(table *TableInBatchWriteRowRequest) *BatchWriteRowRequest {

	this.Tables = append(this.Tables, table)
	return this
}

func (this *BatchWriteRowRequest) PBStruct() *pb.BatchWriteRowRequest {

	p := &pb.BatchWriteRowRequest{}

	for _, item := range this.Tables {
		p.Tables = append(p.Tables, item.PBStruct())
	}

	return p
}
