package main

import (
	"fmt"

	"github.com/gunsluo/ali-tablestore/ots"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

func test_create_table() {

	otsClient := ots.NewClient("DeCtic3Hz7dOJysd", "").
		SetInstance("nt-test").SetRegion("cn-hangzhou")

	req := &ots.CreateTableRequest{
		TableName: "log_trace",
		PrimaryKey: []*ots.ColumnSchema{
			{Name: "id", Type: pb.ColumnType_INTEGER},
		},
		CapacityUnit: ots.CapacityUnit{
			Read:  100,
			Write: 100,
		},
	}
	//req.AddPrimaryKey("col3", pb.ColumnType_STRING)

	err := otsClient.CreateTable(req)
	if err != nil {
		fmt.Println("create table - errmsg:", err)
	} else {
		fmt.Println("create table success")
	}
}

func main() {

	test_create_table()
}
