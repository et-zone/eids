package sonyflake

import "github.com/et-zone/eids/sonyflake/internal"

var EIDCliet EIDs

const (
	B_e18 = "b_18"
	B_e19 = "b_19"
)

type eid struct{}

func (e *eid) InitSonyFlake(machineID int32) error {
	err:=internal.InitMachineID(machineID)
	if err!=nil{
		return err
	}
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
	// byteSize is const val,before do InitSonyFlake
	SetByteSzie(byteSize string) error
}

func init() {
	EIDCliet = &eid{}
}

func NextID() (uint64, error) {
	return internal.NexitID()
}
