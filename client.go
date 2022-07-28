package eids

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/et-zone/eids/sonyflake"
	"io/ioutil"
	"net/http"
)

type req struct {
	Size int64 `gorm:"column:size" json:"IDSize"`
	Num int64 `gorm:"column:num" json:"num"`
	Msg int64 `gorm:"column:msg" json:"msg"`
}

type Config struct {
	SrvID int64 `json:"srv_id"`
	Addr string `json:"addr"`//"https://127.0.0.1:8080"
}


func getNum(addr string,serID int64)(size string,num int32,err error){
	//r,err:=http.Get("http://127.0.0.1:8080/getnum?srvID=1001658977917")
	r,err:=http.Get(addr+"/getnum?srvID="+fmt.Sprintf("%v",serID))
	if err!=nil{
		fmt.Println(err.Error())
		return "",0,err
	}
	b,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println(err.Error())
		return "",0,err
	}
	rv:=&req{}
	json.Unmarshal(b,rv)
	if rv.Size==18{
		return sonyflake.B_e18,int32(rv.Num),nil
	}
	if rv.Size==19{
		return sonyflake.B_e19,int32(rv.Num),nil
	}

	return "",0,errors.New("nservID = "+fmt.Sprintf("%v",serID)+" not find ")
}

func Init(cfg *Config)error{
	size,num,err:=getNum(cfg.Addr,cfg.SrvID)
	if err!=nil{
		return err
	}
	return sonyflake.InitSonyFlakeWithSzie(num,size)
}