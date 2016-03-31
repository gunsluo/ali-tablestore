package ots

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

type ColumnValue struct {
	Type    pb.ColumnType
	VInt    int64
	VString string
	VBool   bool
	VDouble float64
	VBinary []byte
}

func (this *ColumnValue) PBStruct() *pb.ColumnValue {
	p := &pb.ColumnValue{
		Type: &this.Type,
	}

	switch this.Type {
	case pb.ColumnType_INF_MIN:
		p.VInt = proto.Int64(this.VInt)
	case pb.ColumnType_INF_MAX:
		p.VInt = proto.Int64(this.VInt)
	case pb.ColumnType_INTEGER:
		p.VInt = proto.Int64(this.VInt)
	case pb.ColumnType_STRING:
		p.VString = proto.String(this.VString)
	case pb.ColumnType_BOOLEAN:
		p.VBool = proto.Bool(this.VBool)
	case pb.ColumnType_DOUBLE:
		p.VDouble = proto.Float64(this.VDouble)
	case pb.ColumnType_BINARY:
		p.VBinary = this.VBinary
	}

	return p
}

func NewColumnValue(p *pb.ColumnValue) *ColumnValue {
	resp := &ColumnValue{
		Type: p.GetType(),
	}

	switch resp.Type {
	case pb.ColumnType_INF_MIN:
		resp.VInt = p.GetVInt()
	case pb.ColumnType_INF_MAX:
		resp.VInt = p.GetVInt()
	case pb.ColumnType_INTEGER:
		resp.VInt = p.GetVInt()
	case pb.ColumnType_STRING:
		resp.VString = p.GetVString()
	case pb.ColumnType_BOOLEAN:
		resp.VBool = p.GetVBool()
	case pb.ColumnType_DOUBLE:
		resp.VDouble = p.GetVDouble()
	case pb.ColumnType_BINARY:
		resp.VBinary = p.VBinary
	}

	return resp
}
