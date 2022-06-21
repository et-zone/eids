package public

import (
	"fmt"
	"net"
)
var ip net.IP = nil
func getLocalIP()(net.IP,error){
	addrs,err:=net.InterfaceAddrs()
	if err!=nil{
		return nil,err
	}
	for _,add:=range addrs{
		if ipnet,ok:=add.(*net.IPNet);ok&&!ipnet.IP.IsLoopback(){
			if ipnet.IP.To4()!=nil{
				fmt.Println(add.String())
				return ipnet.IP,nil
			}
		}

	}
	return nil,err
}

func SetIP(ipstring string){
	ip=[]byte(ipstring)
}