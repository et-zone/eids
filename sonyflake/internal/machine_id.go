package internal

import (
	"errors"
	"github.com/et-zone/eids/conf"
)
var mID *int32

func InitMachineID(machine_id int32){
	if machineID==nil{
		tmp:=machine_id
		mID=&tmp
	}
}

func machineID()(uint16,error){
	if conf.ConMsg==nil{
		return 0,errors.New("config not init , not get machineID ")
	}
	return uint16(conf.ConMsg.ServID), nil
}

func checkMachineID(maid uint16) bool{
	if mID==nil{
		return false
	}
	return maid==uint16(*mID)
}