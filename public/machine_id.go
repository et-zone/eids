package public

import (
	"errors"
	"github.com/et-zone/eids/conf"
)

func machineID()(uint16,error){
	if conf.ConMsg==nil{
		return 0,errors.New("config not init , not get machineID ")
	}
	return uint16(conf.ConMsg.ServID), nil
}

func checkMachineID(maid uint16) bool{
	return maid==uint16(conf.ConMsg.ServID)
}