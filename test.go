package eids

import (
	"fmt"
	"testing"
)

func Test(t testing.T){
	err:=Init(&Config{
		SrvID: 1001658977917,
		Addr:"http://127.0.0.1:19001",//default port = 19001
	})
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(NextID())
}
