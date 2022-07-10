package internal

import (
	"errors"
	"fmt"
)

var mID *int32


func InitMachineID(machine_id int32)error{
	if machine_id<0||machine_id>=1<<BitLenMachineID{
		return errors.New(fmt.Sprintf( "machineID out of range ,machineID must < %v",1<<BitLenMachineID))
	}
	if mID==nil{
		mID=&machine_id
	}
	return nil
}

func machineID()(uint16,error){
	return uint16(*mID), nil
}

func checkMachineID(maid uint16) bool{
	if mID==nil{
		return false
	}
	return maid==uint16(*mID)
}