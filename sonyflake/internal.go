package sonyflake

import (
	"fmt"
	"github.com/et-zone/eids/sonyflake/internal"
)

var Cliet EIDs

const (
	B_e18 = "b_18"
	B_e19 = "b_19"
)

type eid struct{
	*internal.Sonyflake
}

func InitSonyFlake(machineID int32) error {
	err:=internal.InitMachineID(machineID)
	if err!=nil{
		return err
	}
	client,err:=internal.InitSonyflake()
	if err!=nil{
		return err
	}
	Cliet=&eid{client}
	return nil
}

func InitSonyFlakeWithSzie(machineID int32,id_size string) error {
	err:=internal.InitMachineID(machineID)
	if err!=nil{
		return err
	}
	err=internal.SetByteSzie(id_size)
	if err!=nil{
		return err
	}

	client,err:=internal.InitSonyflake()
	if err!=nil{
		return err
	}

	Cliet=&eid{client}
	fmt.Printf("eids init succ ")
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

