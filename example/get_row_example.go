package main

import (
	"fmt"

	"github.com/gunsluo/ali-tablestore/ots"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

func test_get_row() {

	otsClient := ots.NewClient("DeCtic3Hz7dOJysd", "").
		SetInstance("nt-test").SetRegion("cn-hangzhou")

	req := &ots.GetRowRequest{
		TableName: "log_trace",
	}
	req.AddPrimaryKey("id", ots.ColumnValue{Type: pb.ColumnType_INTEGER, VInt: 111})
	//req.AddPrimaryKey("val", ots.ColumnValue{Type: pb.ColumnType_STRING, VString: "this is a test"})

	info, err := otsClient.GetRow(req)
	if err != nil {
		fmt.Println("get row - errmsg:", err)
	} else {
		fmt.Printf("get row success. [%+v]\n", info)
	}
}

func main() {

	test_get_row()
}
