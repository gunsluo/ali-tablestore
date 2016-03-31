package main

import (
	"fmt"

	"github.com/gunsluo/ali-tablestore/ots"
)

func test_delete_table() {

	otsClient := ots.NewClient("DeCtic3Hz7dOJysd", "").
		SetInstance("nt-test").SetRegion("cn-hangzhou")

	req := &ots.DeleteTableRequest{
		TableName: "log_trace",
	}

	err := otsClient.DeleteTable(req)
	if err != nil {
		fmt.Println("delete table - errmsg:", err)
	} else {
		fmt.Println("delete table success")
	}
}

func main() {

	test_delete_table()
}
