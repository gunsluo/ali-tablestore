package main

import (
	"fmt"

	"github.com/gunsluo/ali-tablestore/ots"
)

func test_list_table() {

	otsClient := ots.NewClient("DeCtic3Hz7dOJysd", "").
		SetInstance("nt-test").SetRegion("cn-hangzhou")

	tableNames, err := otsClient.ListTable()
	if err != nil {
		fmt.Println("list table - errmsg:", err)
	} else {
		fmt.Println("list table success", tableNames)
	}
}

func main() {

	test_list_table()
}
