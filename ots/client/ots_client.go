package client

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

type OTSClient struct {
	AccessKeyID     string
	AccessKeySecret string
	timeout         int
	builder         *RequestBuilder
	author          *ResponseAuthorization
}

func NewOTSClient(accessKeyID, accessKeySecret string) *OTSClient {
	client := new(OTSClient)
	client.AccessKeyID = accessKeyID
	client.AccessKeySecret = accessKeySecret
	client.builder = NewRequestBuilder(accessKeyID, accessKeySecret)
	client.author = NewResponseAuthorization(accessKeyID, accessKeySecret)
	client.timeout = DEFAULT_TIMEOUT
	return client
}

func (this *OTSClient) SetInternal(internal bool) *OTSClient {
	this.builder.internal = internal
	return this
}

func (this *OTSClient) SetInstance(instance string) *OTSClient {
	this.builder.instance = instance
	return this
}

func (this *OTSClient) SetRegion(region string) *OTSClient {
	this.builder.region = region
	return this
}

func (this *OTSClient) SetTimeout(timeout int) *OTSClient {
	this.timeout = timeout
	return this
}

func (this *OTSClient) GetRow(request *pb.GetRowRequest) (response pb.GetRowResponse, err error) {
	err = this.sendToOTSServer("GetRow", request, &response)
	return
}

func (this *OTSClient) PutRow(request *pb.PutRowRequest) (response pb.PutRowResponse, err error) {
	err = this.sendToOTSServer("PutRow", request, &response)
	return
}

func (this *OTSClient) UpdateRow(request *pb.UpdateRowRequest) (response pb.UpdateRowResponse, err error) {
	err = this.sendToOTSServer("UpdateRow", request, &response)
	return
}

func (this *OTSClient) DeleteRow(request *pb.DeleteRowRequest) (response pb.DeleteRowResponse, err error) {
	err = this.sendToOTSServer("DeleteRow", request, &response)
	return
}

func (this *OTSClient) GetRange(request *pb.GetRangeRequest) (response pb.GetRangeResponse, err error) {
	err = this.sendToOTSServer("GetRange", request, &response)
	return
}

func (this *OTSClient) BatchGetRow(request *pb.BatchGetRowRequest) (response pb.BatchGetRowResponse, err error) {
	err = this.sendToOTSServer("BatchGetRow", request, &response)
	return
}

func (this *OTSClient) BatchWriteRow(request *pb.BatchWriteRowRequest) (response pb.BatchWriteRowResponse, err error) {
	err = this.sendToOTSServer("BatchWriteRow", request, &response)
	return
}

func (this *OTSClient) CreateTable(request *pb.CreateTableRequest) (response pb.CreateTableResponse, err error) {

	err = this.sendToOTSServer("CreateTable", request, &response)
	return
}

func (this *OTSClient) ListTable() (response pb.ListTableResponse, err error) {

	request := new(pb.ListTableRequest)
	err = this.sendToOTSServer("ListTable", request, &response)
	return
}

func (this *OTSClient) DeleteTable(request *pb.DeleteTableRequest) (response pb.DeleteTableResponse, err error) {
	err = this.sendToOTSServer("DeleteTable", request, &response)
	return
}

func (this *OTSClient) UpdateTable(request *pb.UpdateTableRequest) (response pb.UpdateTableResponse, err error) {
	err = this.sendToOTSServer("UpdateTable", request, &response)
	return
}

func (this *OTSClient) DescribeTable(request *pb.DescribeTableRequest) (response pb.DescribeTableResponse, err error) {
	err = this.sendToOTSServer("DescribeTable", request, &response)
	return
}

func (this *OTSClient) sendMsg(request *http.Request) (response *http.Response, body []byte, err error) {
	if request == nil {
		return nil, nil, errors.New("request is nil")
	}

	timeout := time.Duration(this.timeout) * time.Second
	client := &http.Client{Timeout: timeout}
	response, err = client.Do(request)
	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(response.Body)

	return
}

func (this *OTSClient) sendToOTSServer(operation string, request proto.Message, ret proto.Message) (err error) {

	// to protocol buffer
	buffer, err := proto.Marshal(request)
	if err != nil {
		return
	}

	// send msg
	req := this.builder.NewRequest(operation, buffer)
	resp, respBody, err := this.sendMsg(req)
	if err != nil {
		return
	}

	// auth
	err = this.author.Auth(operation, respBody, resp)
	if err != nil {
		return err
	}

	err = proto.Unmarshal(respBody, ret)
	if err != nil {
		return
	}

	return
}
