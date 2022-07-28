package sonyflake

import (
	"github.com/et-zone/eids/sonyflake/internal"
	"log"
)

var Cliet EIDs

const (
	B_e18 = "b_18"
	B_e19 = "b_19"
)

type eid struct{
	*internal.Sonyflake
}

func InitSonyFlake(machineID *uint16) error {

	client,err:=internal.InitSonyflake(machineID)
	if err!=nil{
		return err
	}
	Cliet=&eid{client}
	return nil
}

func InitSonyFlakeWithSzie(machineID *uint16,id_size string) error {
	err:=internal.SetByteSzie(id_size)
	if err!=nil{
		return err
	}

	client,err:=internal.InitSonyflake(machineID)
	if err!=nil{
		return err
	}

	Cliet=&eid{client}

	log.Println("eids init succ ")
	return nil
}

func InitSonyFlakeByDefault() error {
	client,err:=internal.InitSonyflakeDefault()
	if err!=nil{
		return err
	}
	Cliet=&eid{client}
	return nil
}

type EIDs interface {
	GetMachineID() (uint16,error)
	NextID() (uint64, error)
}

