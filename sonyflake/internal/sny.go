package internal

import (
	"errors"
	"fmt"
	"log"
)

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
	log.Println("ids machineID = ",ids.machineID)
	if ids.machineID<0||ids.machineID>=1<<BitLenMachineID{
		return errors.New(fmt.Sprintf( "machineID out of range ,machineID must < %v",1<<BitLenMachineID))
	}
	return nil
}

func InitSonyflakeDefault()error{
	st := Settings{}
	ids = newSonyflake(st)
	if ids == nil {
		return errors.New("sonyflake not created")
	}
	log.Println("ids machineID = ",ids.machineID)
	mid:=int32(ids.machineID)
	mID=&mid
	if ids.machineID<0||ids.machineID>=1<<BitLenMachineID{
		return errors.New(fmt.Sprintf( "machineID out of range ,machineID must < %v",1<<BitLenMachineID))
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

func SetByteSzie(byteSize string)(error){
	if byteSize==b_e18{
		sonyflakeTimeUnit=byteSize_e7
		return nil
	}
	if byteSize==b_e19{
		sonyflakeTimeUnit=byteSize_e6
		return nil
	}
	return errors.New("SetByteSzie err , args val err")
}