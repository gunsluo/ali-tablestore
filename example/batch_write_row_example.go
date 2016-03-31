package main

import (
	"fmt"

	"github.com/gunsluo/ali-tablestore/ots"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

func test_batch_write_row() {

	otsClient := ots.NewClient("DeCtic3Hz7dOJysd", "").
		SetInstance("nt-test").SetRegion("cn-hangzhou")

	req := &ots.BatchWriteRowRequest{}

	row := &ots.PutRowInBatchWriteRowRequest{
		RowExistence: pb.RowExistenceExpectation_IGNORE,
	}
	row.AddPrimaryKey("id", ots.ColumnValue{Type: pb.ColumnType_INTEGER, VInt: 211})
	row.AddAttributeColumns("val", ots.ColumnValue{Type: pb.ColumnType_STRING, VString: "modify value"})

	row2 := &ots.UpdateRowInBatchWriteRowRequest{
		RowExistence: pb.RowExistenceExpectation_IGNORE,
	}
	row2.AddPrimaryKey("id", ots.ColumnValue{Type: pb.ColumnType_INTEGER, VInt: 115})
	row2.AddAttributeColumns("val", pb.OperationType_PUT, ots.ColumnValue{Type: pb.ColumnType_STRING, VString: "modify value too"})

	row3 := &ots.DeleteRowInBatchWriteRowRequest{
		RowExistence: pb.RowExistenceExpectation_IGNORE,
	}
	row3.AddPrimaryKey("id", ots.ColumnValue{Type: pb.ColumnType_INTEGER, VInt: 113})

	table := &ots.TableInBatchWriteRowRequest{
		TableName: "log_trace",
	}
	table.AddPutRow(row)
	table.AddUpdateRow(row2)
	table.AddDeleteRow(row3)

	req.AddTable(table)

	info, err := otsClient.BatchWriteRow(req)
	if err != nil {
		fmt.Println("batch write row - errmsg:", err)
	} else {
		fmt.Println("batch write row success.")
		for _, table := range info.Tables {
			fmt.Printf("table=[%s]:\n", table.TableName)
			for _, row := range table.PutRows {
				fmt.Printf("\tPut row isok[%t][%d] Error.code[%s] Error.msg[%s]\n", row.IsOk, row.CapacityUnit.Write, row.Error.Code, row.Error.Message)
			}
			for _, row := range table.UpdateRows {
				fmt.Printf("\tUpd row isok[%t][%d] Error.code[%s] Error.msg[%s]\n", row.IsOk, row.CapacityUnit.Write, row.Error.Code, row.Error.Message)
			}
			for _, row := range table.DeleteRows {
				fmt.Printf("\tDel row isok[%t][%d] Error.code[%s] Error.msg[%s]\n", row.IsOk, row.CapacityUnit.Write, row.Error.Code, row.Error.Message)
			}
		}
	}
}

func main() {

	test_batch_write_row()
}
