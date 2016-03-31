package main

import (
	"fmt"

	"github.com/gunsluo/ali-tablestore/ots"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

func test_update_row() {

	otsClient := ots.NewClient("DeCtic3Hz7dOJysd", "").
		SetInstance("nt-test").SetRegion("cn-hangzhou")

	req := &ots.UpdateRowRequest{
		TableName:    "log_trace",
		RowExistence: pb.RowExistenceExpectation_IGNORE,
	}
	req.AddPrimaryKey("id", ots.ColumnValue{Type: pb.ColumnType_INTEGER, VInt: 111})
	req.AddAttributeColumns("val", pb.OperationType_PUT, ots.ColumnValue{Type: pb.ColumnType_STRING, VString: "modify this is a test"})

	count, err := otsClient.UpdateRow(req)
	if err != nil {
		fmt.Println("update row - errmsg:", err)
	} else {
		fmt.Printf("update row[%d] success\n", count)
	}
}

func main() {

	test_update_row()
}
