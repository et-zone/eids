package internal

import (
	"errors"
	"fmt"
	"net"
)

//var ids *Sonyflake

func InitSonyflake(machineID *uint16,size string) (*Sonyflake, error) {

	st := Settings{}
	ids := newSonyflake(st,size)

	if ids == nil {
		return nil, errors.New("sonyflake not created")
	}
	if machineID != nil {
		ids.machineID = *machineID
	}
	ids.size=size
	//log.Println("ids machineID = ",ids.machineID)
	if ids.machineID < 0 || ids.machineID >= 1<<BitLenMachineID {
		return nil, errors.New(fmt.Sprintf("machineID out of range ,machineID must < %v", 1<<BitLenMachineID))
	}
	return ids, nil
}


func lowerPrivateIPv4() (uint16, error) {
	ip, err := privateIPv4()
	if err != nil {
		return 0, err
	}
	//fmt.Println(ip[2], ip[3])
	return uint16(ip[0]) + uint16(ip[1]) + uint16(ip[2]) + uint16(ip[3]), nil
}

func lower16BitPrivateIP() (uint16, error) {
	ip, err := privateIPv4()
	if err != nil {
		return 0, err
	}

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}

func privateIPv4() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip, nil
		}
	}
	return nil, errors.New("no private ip address")
}

func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

