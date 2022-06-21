package public

import (
	"errors"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

func MachineID()(uint16,error){
	ip,err:=getLocalIP()
	if err!=nil{
		return 0,err
	}
	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}


// TimeDifference returns the time difference between the localhost and the given NTP server.
func timeDifference(server string) (time.Duration, error) {
	output, err := exec.Command("/usr/sbin/ntpdate", "-q", server).CombinedOutput()
	if err != nil {
		return time.Duration(0), err
	}

	re, _ := regexp.Compile("offset (.*) sec")
	submatched := re.FindSubmatch(output)
	if len(submatched) != 2 {
		return time.Duration(0), errors.New("invalid ntpdate output")
	}

	f, err := strconv.ParseFloat(string(submatched[1]), 64)
	if err != nil {
		return time.Duration(0), err
	}
	return time.Duration(f*1000) * time.Millisecond, nil
}