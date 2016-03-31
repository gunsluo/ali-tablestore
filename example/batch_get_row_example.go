package main

import (
	"fmt"

	"github.com/gunsluo/ali-tablestore/ots"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

func test_batch_get_row() {

	otsClient := ots.NewClient("DeCtic3Hz7dOJysd", "").
		SetInstance("nt-test").SetRegion("cn-hangzhou")

	req := &ots.BatchGetRowRequest{}

	row := &ots.RowInBatchGetRowRequest{}
	row.AddPrimaryKey("id", ots.ColumnValue{Type: pb.ColumnType_INTEGER, VInt: 111})

	row2 := &ots.RowInBatchGetRowRequest{}
	row2.AddPrimaryKey("id", ots.ColumnValue{Type: pb.ColumnType_INTEGER, VInt: 115})
	table := &ots.TableInBatchGetRowRequest{
		TableName: "log_trace",
	}
	table.AddRow(row)
	table.AddRow(row2)

	req.AddTable(table)

	info, err := otsClient.BatchGetRow(req)
	if err != nil {
		fmt.Println("batch get row - errmsg:", err)
	} else {
		fmt.Println("batch get row success.")
		for _, table := range info.Tables {
			fmt.Printf("table=[%s]:\n", table.TableName)
			for _, row := range table.Rows {
				fmt.Printf("\trow data isok[%t]\n", row.IsOk)
				for idx, info := range row.Row.PrimaryKeyColumns {
					attrInfo := row.Row.AttributeColumns[idx]
					fmt.Printf("\tpk.name=[%s] pk.val=[%d] atrr.name=[%s] atrr.val=[%s]\n", info.Name, info.Value.VInt, attrInfo.Name, attrInfo.Value.VString)
				}
			}
		}
	}
}

func main() {

	test_batch_get_row()
}
