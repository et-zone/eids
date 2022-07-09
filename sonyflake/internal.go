package sonyflake

import "github.com/et-zone/eids/sonyflake/internal"

var EIDCliet EIDs

type eid struct {
}

func (e *eid)InitSonyFlake(machineID int32) error {
	internal.InitMachineID(machineID)
	return internal.InitSonyflake()
}

func (e *eid)MachineID() (int32, error) {
	return internal.MachineID()
}

type EIDs interface {
	InitSonyFlake(machineID int32) error
	MachineID()(int32,error)
}

func init(){
	EIDCliet=&eid{}
}

func NextID() (uint64, error) {
	return internal.NexitID()
}
