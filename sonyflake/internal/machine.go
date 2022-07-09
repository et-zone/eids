package internal

import "errors"

var mID *int32

const MaxNum int32=256

func InitMachineID(machine_id int32)error{
	if machine_id<0||machine_id>=MaxNum{
		return errors.New("machineID out of range ,machineID must < 256")
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