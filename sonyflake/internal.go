package sonyflake

import "github.com/et-zone/eids/sonyflake/internal"

var EIDCliet EIDs

const (
	b_e18 = "b_18"
	b_e19 = "b_19"
)

type eid struct{}

func (e *eid) InitSonyFlake(machineID int32) error {
	internal.InitMachineID(machineID)
	return internal.InitSonyflake()
}

func (e *eid) MachineID() (int32, error) {
	return internal.MachineID()
}

func (e *eid) SetByteSzie(byteSize string) error {
	return internal.SetByteSzie(byteSize)
}

type EIDs interface {
	InitSonyFlake(machineID int32) error
	MachineID() (int32, error)
	SetByteSzie(byteSize string) error //byteSize is const val
}

func init() {
	EIDCliet = &eid{}
}

func NextID() (uint64, error) {
	return internal.NexitID()
}
