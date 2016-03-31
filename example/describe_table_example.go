package main

import (
	"fmt"

	"github.com/gunsluo/ali-tablestore/ots"
)

func test_describe_table() {

	otsClient := ots.NewClient("DeCtic3Hz7dOJysd", "").
		SetInstance("nt-test").SetRegion("cn-hangzhou")

	req := &ots.DescribeTableRequest{
		TableName: "log_trace",
	}

	info, err := otsClient.DescribeTable(req)
	if err != nil {
		fmt.Println("describe table - errmsg:", err)
	} else {
		fmt.Println("describe table success", info)
	}
}

func main() {

	test_describe_table()
}
