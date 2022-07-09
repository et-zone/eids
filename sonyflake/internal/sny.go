package internal

import "errors"

var ids *Sonyflake
func InitSonyflake()error{
	if mID==nil{
		return errors.New(" init machineID fail")
	}
	st := Settings{
		MachineID:      machineID,
		CheckMachineID: checkMachineID,
	}
	ids = newSonyflake(st)
	if ids == nil {
		return errors.New("sonyflake not created")
	}
	return nil
}
func NexitID()(uint64,error){
	if ids==nil{
		return 0,errors.New("sonyflake not created")
	}
	return ids.NextID()
}
func MachineID()(int32,error){
	if ids==nil{
		return 0,errors.New("machineID not find")
	}
	return int32(ids.machineID),nil
}
