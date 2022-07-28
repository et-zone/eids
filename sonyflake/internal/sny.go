package internal

import (
	"errors"
	"fmt"
	"log"
)

//var ids *Sonyflake

func InitSonyflake()(*Sonyflake,error){
	if mID==nil{
		return nil,errors.New(" init machineID fail")
	}
	st := Settings{
		MachineID:      machineID,
		CheckMachineID: checkMachineID,
	}
	ids := newSonyflake(st)
	if ids == nil {
		return nil,errors.New("sonyflake not created")
	}
	log.Println("ids machineID = ",ids.machineID)
	if ids.machineID<0||ids.machineID>=1<<BitLenMachineID{
		return nil,errors.New(fmt.Sprintf( "machineID out of range ,machineID must < %v",1<<BitLenMachineID))
	}
	return ids,nil
}

func InitSonyflakeDefault()(*Sonyflake,error){
	st := Settings{}
	ids := newSonyflake(st)
	if ids == nil {
		return nil,errors.New("sonyflake not created")
	}
	log.Println("ids machineID = ",ids.machineID)
	mid:=int32(ids.machineID)
	mID=&mid
	if ids.machineID<0||ids.machineID>=1<<BitLenMachineID{
		return nil,errors.New(fmt.Sprintf( "machineID out of range ,machineID must < %v",1<<BitLenMachineID))
	}
	return ids,nil
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