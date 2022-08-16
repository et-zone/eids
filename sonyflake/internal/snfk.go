package internal

import (
	"errors"
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

const (
	sonyflakeTimeUnit   = 1e6 //size = 19 , nsec, i.e.  1 msec
	sonyflakeTimeUnitE6 = 1e6 //size = 19 , nsec, i.e.  1 msec
	sonyflakeTimeUnitE7 = 1e7 //size = 18 , nsec, i.e. 10 msec
	BSize18e7           = "size_18" //size = 18 , nsec, i.e. 10 msec
	BSize19e6           = "size_19" //size = 19 , nsec, i.e.  1 msec
)

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
	size        string
}

// NewSonyflake returns a new Sonyflake configured with the given Settings.
// NewSonyflake returns nil in the following cases:
// - Settings.StartTime is ahead of the current time.
// - Settings.MachineID returns an error.
// - Settings.CheckMachineID returns false.
func newSonyflake(st Settings, size string) *Sonyflake {
	sf := new(Sonyflake)
	sf.mutex = new(sync.Mutex)
	sf.sequence = uint16(1<<BitLenSequence - 1)
	sf.size = size
	if st.StartTime.After(time.Now()) {
		return nil
	}
	if st.StartTime.IsZero() {
		sf.startTime = toSonyflakeTime(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC), sf.size)
	} else {
		sf.startTime = toSonyflakeTime(st.StartTime, sf.size)
	}
	return sf
}

func (sf *Sonyflake) GetMachineID() (uint16, error) {
	if sf == nil {
		return 0, errors.New("not init client")
	}
	return sf.machineID, nil
}

// NextID generates a next unique ID.
// After the Sonyflake time overflows, NextID returns an error.
func (sf *Sonyflake) NextID() (uint64, error) {
	if sf == nil {
		return 0, errors.New("not init client")
	}
	const maskSequence = uint16(1<<BitLenSequence - 1)

	sf.mutex.Lock()
	defer sf.mutex.Unlock()
Tag:
	current := sf.currentElapsedTime(sf.startTime)
	if sf.elapsedTime < current {

		sf.elapsedTime = current
		sf.sequence = 0
	} else { // sf.elapsedTime >= current
		sf.sequence = (sf.sequence + 1) & maskSequence
		if sf.sequence == 0 {
			time.Sleep(sleepTimeNS(sf.size))
			goto Tag
		}
	}

	return sf.toID()
}
func (sf *Sonyflake) currentElapsedTime(startTime int64) int64 {
	return toSonyflakeTime(time.Now(), sf.size) - sf.startTime
}

func toSonyflakeTime(t time.Time, size string) int64 {
	//return t.UTC().UnixNano() / sonyflakeTimeUnit
	switch size {
	case BSize18e7:
		return t.UTC().UnixNano() / sonyflakeTimeUnitE7
	case BSize19e6:
		return t.UTC().UnixNano() / sonyflakeTimeUnitE6
	}
	return 0
}

func sleepTimeNS(size string) time.Duration {
	//return time.Duration(1*sonyflakeTimeUnit+10) * time.Nanosecond
	switch size {
	case BSize18e7:
		return time.Duration(1*sonyflakeTimeUnitE7+10) * time.Nanosecond
	case BSize19e6:
		return time.Duration(1*sonyflakeTimeUnitE6+10) * time.Nanosecond
	}
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