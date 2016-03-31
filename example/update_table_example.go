package main

import (
	"fmt"

	"github.com/gunsluo/ali-tablestore/ots"
)

func test_update_table() {

	otsClient := ots.NewClient("DeCtic3Hz7dOJysd", "").
		SetInstance("nt-test").SetRegion("cn-hangzhou")

	req := &ots.UpdateTableRequest{
		TableName: "log_trace",
		CapacityUnit: ots.CapacityUnit{
			Read:  10,
			Write: 10,
		},
	}

	info, err := otsClient.UpdateTable(req)
	if err != nil {
		fmt.Println("update table - errmsg:", err)
	} else {
		fmt.Println("update table success", info)
	}
}

func main() {

	test_update_table()
}
