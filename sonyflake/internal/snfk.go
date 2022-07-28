package internal

import (
	"errors"
	"net"
	"sync"
	"time"
)

// Package sonyflake implements Sonyflake, a distributed unique ID generator inspired by Twitter's Snowflake.
// default
// A Sonyflake ID is composed of
//     39 bits for time in units of 10 msec
//      8 bits for a sequence number
//     16 bits for a machine id

//     different machineID ,get id is different

// These constants are the bit lengths of Sonyflake ID parts.
const (
	BitLenTime      = 39                               // bit length of time
	BitLenSequence  = 13                               // bit length of sequence number
	BitLenMachineID = 63 - BitLenTime - BitLenSequence // bit length of machine id
)
const byteSize_e7 int64 = 1e7 //18   10msec
const byteSize_e6 int64 = 1e6 //19    1msec
const (
	b_e18 = "b_18"
	b_e19 = "b_19"
)

var sonyflakeTimeUnit = byteSize_e6

//const sonyflakeTimeUnit = 1e7 // nsec, i.e. 10 msec     size=18 count ~ 70
//const sonyflakeTimeUnit = 1e6 // nsec, i.e.  1 msec   size=19 count ~ 500

// Settings configures Sonyflake:
//
// StartTime is the time since which the Sonyflake time is defined as the elapsed time.
// If StartTime is 0, the start time of the Sonyflake is set to "2014-09-01 00:00:00 +0000 UTC".
// If StartTime is ahead of the current time, Sonyflake is not created.
//
// MachineID returns the unique ID of the Sonyflake instance.
// If MachineID returns an error, Sonyflake is not created.
// If MachineID is nil, default MachineID is used.
// Default MachineID returns the lower 16 bits of the private IP address.
//
// CheckMachineID validates the uniqueness of the machine ID.
// If CheckMachineID returns false, Sonyflake is not created.
// If CheckMachineID is nil, no validation is done.

type Settings struct {
	StartTime      time.Time
	MachineID      func() (uint16, error)
	CheckMachineID func(uint16) bool
}

// Sonyflake is a distributed unique ID generator.
type Sonyflake struct {
	mutex       *sync.Mutex
	startTime   int64
	elapsedTime int64
	sequence    uint16
	machineID   uint16
}

// NewSonyflake returns a new Sonyflake configured with the given Settings.
// NewSonyflake returns nil in the following cases:
// - Settings.StartTime is ahead of the current time.
// - Settings.MachineID returns an error.
// - Settings.CheckMachineID returns false.
func newSonyflake(st Settings) *Sonyflake {
	sf := new(Sonyflake)
	sf.mutex = new(sync.Mutex)
	sf.sequence = uint16(1<<BitLenSequence - 1)

	if st.StartTime.After(time.Now()) {
		return nil
	}
	if st.StartTime.IsZero() {
		sf.startTime = toSonyflakeTime(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	} else {
		sf.startTime = toSonyflakeTime(st.StartTime)
	}

	var err error
	if st.MachineID == nil {
		//sf.machineID, err = lower16BitPrivateIP()
		sf.machineID, err = lowerPrivateIPv4()
	} else {
		sf.machineID, err = st.MachineID()
	}
	if err != nil || (st.CheckMachineID != nil && !st.CheckMachineID(sf.machineID)) {
		return nil
	}
	//fmt.Println("machineID=",sf.machineID)
	return sf
}

func (sf *Sonyflake) GetMachineID() (uint16,error) {
	if sf==nil{
		return 0,errors.New("not init client")
	}
	return sf.machineID,nil
}

// NextID generates a next unique ID.
// After the Sonyflake time overflows, NextID returns an error.
func (sf *Sonyflake) NextID() (uint64, error) {
	if sf==nil{
		return 0,errors.New("not init client")
	}
	const maskSequence = uint16(1<<BitLenSequence - 1)

	sf.mutex.Lock()
	defer sf.mutex.Unlock()
Tag:
	current := currentElapsedTime(sf.startTime)
	if sf.elapsedTime < current {

		sf.elapsedTime = current
		sf.sequence = 0
	} else { // sf.elapsedTime >= current
		sf.sequence = (sf.sequence + 1) & maskSequence
		//fmt.Println(sf.sequence)
		if sf.sequence == 0 {
			//fmt.Println(sf.sequence)
			//overtime := sf.elapsedTime - current
			//time.Sleep(sleepTime((overtime)))
			time.Sleep(sleepTimeNS())
			goto Tag
		}
		//fmt.Println(sf.sequence)
	}

	return sf.toID()
}

func toSonyflakeTime(t time.Time) int64 {
	return t.UTC().UnixNano() / sonyflakeTimeUnit
}

func currentElapsedTime(startTime int64) int64 {
	return toSonyflakeTime(time.Now()) - startTime
}

func sleepTime(overtime int64) time.Duration {
	return time.Duration(overtime)*10*time.Millisecond -
		time.Duration(time.Now().UTC().UnixNano()%sonyflakeTimeUnit)*time.Nanosecond
}

func sleepTimeNS() time.Duration {
	return time.Duration(1*sonyflakeTimeUnit+10) * time.Nanosecond
}

func (sf *Sonyflake) toID() (uint64, error) {
	if sf.elapsedTime >= 1<<BitLenTime {
		return 0, errors.New("over the time limit")
	}
	//fmt.Println(uint64(sf.elapsedTime)<<(BitLenSequence+BitLenMachineID) ,uint64(sf.sequence)<<BitLenMachineID,uint64(sf.machineID))
	return uint64(sf.elapsedTime)<<(BitLenSequence+BitLenMachineID) |
		uint64(sf.sequence)<<BitLenMachineID |
		uint64(sf.machineID), nil
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

func lower16BitPrivateIP() (uint16, error) {
	ip, err := privateIPv4()
	if err != nil {
		return 0, err
	}

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}

func lowerPrivateIPv4() (uint16, error) {
	ip, err := privateIPv4()
	if err != nil {
		return 0, err
	}
	//fmt.Println(ip[2], ip[3])
	return uint16(ip[0]) + uint16(ip[1]) + uint16(ip[2]) + uint16(ip[3]), nil
}

// Decompose returns a set of Sonyflake ID parts.
func Decompose(id uint64) map[string]uint64 {
	const maskSequence = uint64((1<<BitLenSequence - 1) << BitLenMachineID)
	const maskMachineID = uint64(1<<BitLenMachineID - 1)

	msb := id >> 63
	time := id >> (BitLenSequence + BitLenMachineID)
	sequence := id & maskSequence >> BitLenMachineID
	machineID := id & maskMachineID
	return map[string]uint64{
		"id":        id,
		"msb":       msb,
		"time":      time,
		"sequence":  sequence,
		"machineID": machineID,
	}
}
