package eids

import "github.com/et-zone/eids/sonyflake"


func NextID() (uint64, error) {
	return sonyflake.Cliet.NextID()
}


func GetMachineID() (uint16,error){
	return sonyflake.Cliet.GetMachineID()
}

//func InitSonyFlake(machineID int32) error {
//	return sonyflake.InitSonyFlake(machineID)
//}

func InitSonyFlakeByDefault() error {
	return sonyflake.InitSonyFlakeByDefault()
}