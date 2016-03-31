package ots

import "github.com/gunsluo/ali-tablestore/ots/proto/pb"

type Error struct {
	Code    string
	Message string
}

func NewError(p *pb.Error) *Error {
	e := &Error{
		Code:    p.GetCode(),
		Message: p.GetMessage(),
	}

	return e
}
