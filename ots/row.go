package ots

import "github.com/gunsluo/ali-tablestore/ots/proto/pb"

type Row struct {
	PrimaryKeyColumns []*Column
	AttributeColumns  []*Column
}

func NewRow(p *pb.Row) *Row {
	resp := new(Row)

	for _, item := range p.PrimaryKeyColumns {
		resp.PrimaryKeyColumns = append(resp.PrimaryKeyColumns, NewColumn(item))
	}

	for _, item := range p.AttributeColumns {
		resp.AttributeColumns = append(resp.AttributeColumns, NewColumn(item))
	}

	return resp
}
