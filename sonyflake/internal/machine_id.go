package internal

var mID *int32

func InitMachineID(machine_id int32){
	if mID==nil{
		mID=&machine_id
	}
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