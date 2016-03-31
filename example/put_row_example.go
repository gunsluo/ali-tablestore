package main

import (
	"fmt"

	"github.com/gunsluo/ali-tablestore/ots"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

func test_put_row() {

	otsClient := ots.NewClient("DeCtic3Hz7dOJysd", "").
		SetInstance("nt-test").SetRegion("cn-hangzhou")

	req := &ots.PutRowRequest{
		TableName:    "log_trace",
		RowExistence: pb.RowExistenceExpectation_IGNORE,
	}
	req.AddPrimaryKey("id", ots.ColumnValue{Type: pb.ColumnType_INTEGER, VInt: 111})
	//req.AddPrimaryKey("val", ots.ColumnValue{Type: pb.ColumnType_STRING, VString: "this is a test"})
	req.AddAttributeColumns("value", ots.ColumnValue{Type: pb.ColumnType_STRING, VString: "this is a test"})

	count, err := otsClient.PutRow(req)
	if err != nil {
		fmt.Println("put row - errmsg:", err)
	} else {
		fmt.Printf("put row[%d] success\n", count)
	}
}

func main() {

	test_put_row()
}
