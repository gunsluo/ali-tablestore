package ots

import (
	"errors"

	"github.com/gunsluo/ali-tablestore/ots/client"
)

type Client struct {
	otsClient *client.OTSClient
}

func NewClient(accessKeyID, accessKeySecret string) *Client {

	c := new(Client)
	c.otsClient = client.NewOTSClient(accessKeyID, accessKeySecret)

	return c
}

func (this *Client) UseInternal() *Client {
	this.otsClient.SetInternal(true)
	return this
}

func (this *Client) UseOuternal() *Client {
	this.otsClient.SetInternal(false)
	return this
}

func (this *Client) SetInstance(instance string) *Client {
	this.otsClient.SetInstance(instance)
	return this
}

func (this *Client) SetRegion(region string) *Client {
	this.otsClient.SetRegion(region)
	return this
}

func (this *Client) SetTimeout(timeout int) *Client {
	this.otsClient.SetTimeout(timeout)
	return this
}

func (this *Client) GetRow(req *GetRowRequest) (resp *GetRowResponse, err error) {

	response, err := this.otsClient.GetRow(req.PBStruct())
	if err != nil {
		return
	}

	return NewGetRowResponse(&response), nil
}

func (this *Client) PutRow(req *PutRowRequest) (c int32, err error) {

	response, err := this.otsClient.PutRow(req.PBStruct())
	if err != nil {
		return
	}

	c = response.GetConsumed().GetCapacityUnit().GetWrite()

	return
}

func (this *Client) UpdateRow(req *UpdateRowRequest) (c int32, err error) {

	response, err := this.otsClient.UpdateRow(req.PBStruct())
	if err != nil {
		return
	}

	c = response.GetConsumed().GetCapacityUnit().GetWrite()
	return
}

func (this *Client) DeleteRow(req *DeleteRowRequest) (c int32, err error) {

	response, err := this.otsClient.DeleteRow(req.PBStruct())
	if err != nil {
		return
	}

	c = response.GetConsumed().GetCapacityUnit().GetWrite()
	return
}

func (this *Client) GetRange(req *GetRangeRequest) (resp *GetRangeResponse, err error) {

	response, err := this.otsClient.GetRange(req.PBStruct())
	if err != nil {
		return
	}

	return NewGetRangeResponse(&response), nil
}

func (this *Client) BatchGetRow(req *BatchGetRowRequest) (resp *BatchGetRowResponse, err error) {

	response, err := this.otsClient.BatchGetRow(req.PBStruct())
	if err != nil {
		return
	}

	return NewBatchGetRowResponse(&response), nil
}

func (this *Client) BatchWriteRow(req *BatchWriteRowRequest) (resp *BatchWriteRowResponse, err error) {

	response, err := this.otsClient.BatchWriteRow(req.PBStruct())
	if err != nil {
		return
	}

	return NewBatchWriteRowResponse(&response), nil
}

func (this *Client) CreateTable(req *CreateTableRequest) (err error) {

	response, err := this.otsClient.CreateTable(req.PBStruct())
	if err != nil {
		return err
	}

	errmsg := response.String()
	if errmsg != "" {
		return errors.New(errmsg)
	}

	return
}

func (this *Client) ListTable() (tableNames []string, err error) {

	response, err := this.otsClient.ListTable()
	if err != nil {
		return
	}

	return response.GetTableNames(), nil
}

func (this *Client) DeleteTable(req *DeleteTableRequest) (err error) {

	response, err := this.otsClient.DeleteTable(req.PBStruct())
	if err != nil {
		return
	}

	errmsg := response.String()
	if errmsg != "" {
		return errors.New(errmsg)
	}

	return
}

func (this *Client) UpdateTable(req *UpdateTableRequest) (resp *UpdateTableResponse, err error) {

	response, err := this.otsClient.UpdateTable(req.PBStruct())
	if err != nil {
		return
	}

	return NewUpdateTableResponse(&response), nil
}

func (this *Client) DescribeTable(req *DescribeTableRequest) (resp *DescribeTableResponse, err error) {

	response, err := this.otsClient.DescribeTable(req.PBStruct())
	if err != nil {
		return
	}

	return NewDescribeTableResponse(&response), nil
}
