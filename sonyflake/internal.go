package sonyflake

import (
	"github.com/et-zone/eids/sonyflake/internal"
	"log"
)
const (
	BSize18e7 = internal.BSize18e7
	BSize19e6 = internal.BSize19e6
)

type eid struct{
	*internal.Sonyflake
}

func InitSonyFlake(machineID *uint16,size string)(Client, error) {
	client,err:=internal.InitSonyflake(machineID,size)
	if err!=nil{
		return nil,err
	}
	log.Println("Init Succ ")
	return 	&eid{client},nil
}

type Client interface {
	GetMachineID() (uint16,error)
	NextID() (uint64, error)
}

