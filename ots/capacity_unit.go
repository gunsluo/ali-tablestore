package ots

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gunsluo/ali-tablestore/ots/proto/pb"
)

type CapacityUnit struct {
	Read  int32
	Write int32
}

func (this *CapacityUnit) PBStruct() *pb.CapacityUnit {
	p := &pb.CapacityUnit{
		Read:  proto.Int32(this.Read),
		Write: proto.Int32(this.Write),
	}

	return p
}

func NewCapacityUnit(p *pb.CapacityUnit) CapacityUnit {
	resp := CapacityUnit{
		Read:  p.GetRead(),
		Write: p.GetWrite(),
	}

	return resp
}
