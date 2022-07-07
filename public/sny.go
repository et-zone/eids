package public

var ids *Sonyflake
func InitIDS(){
	st :=Settings{
		MachineID: machineID,
		CheckMachineID: checkMachineID,
	}
	ids = NewSonyflake(st)
	if ids == nil {
		panic("sonyflake not created")
	}
}

func GetCli()*Sonyflake{
	return ids
}
