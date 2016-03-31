package main

import (
	"fmt"

	"github.com/gunsluo/ali-tablestore/ots"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

func test_delete_row() {

	otsClient := ots.NewClient("DeCtic3Hz7dOJysd", "").
		SetInstance("nt-test").SetRegion("cn-hangzhou")

	req := &ots.DeleteRowRequest{
		TableName:    "log_trace",
		RowExistence: pb.RowExistenceExpectation_IGNORE,
	}
	req.AddPrimaryKey("id", ots.ColumnValue{Type: pb.ColumnType_INTEGER, VInt: 111})

	count, err := otsClient.DeleteRow(req)
	if err != nil {
		fmt.Println("delete row - errmsg:", err)
	} else {
		fmt.Printf("delete row[%d] success\n", count)
	}
}

func main() {

	test_delete_row()
}
