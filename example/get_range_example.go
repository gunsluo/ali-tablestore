package main

import (
	"fmt"

	"github.com/gunsluo/ali-tablestore/ots"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

func test_get_range_row() {

	otsClient := ots.NewClient("DeCtic3Hz7dOJysd", "").
		SetInstance("nt-test").SetRegion("cn-hangzhou")

	req := &ots.GetRangeRequest{
		TableName: "log_trace",
		Direction: pb.Direction_FORWARD,
		Limit:     20,
	}
	req.AddInclusiveStartPrimaryKey("id", ots.ColumnValue{Type: pb.ColumnType_INTEGER, VInt: 111})
	req.AddExclusiveEndPrimaryKey("id", ots.ColumnValue{Type: pb.ColumnType_INTEGER, VInt: 200})

	info, err := otsClient.GetRange(req)
	if err != nil {
		fmt.Println("get range row - errmsg:", err)
	} else {
		fmt.Printf("get range row success. [%+v] IsDone[%t]\n", info, info.IsDone())
	}
}

func main() {

	test_get_range_row()
}
