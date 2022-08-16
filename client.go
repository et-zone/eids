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
	Num  int64 `gorm:"column:num" json:"num"`
	Msg  int64 `gorm:"column:msg" json:"msg"`
}

type Config struct {
	SrvID int64  `json:"srv_id"`
	Addr  string `json:"addr"` //"https://127.0.0.1:19001"
}

func getNum(addr string, serID int64) (size string, num uint16, err error) {
	//r,err:=http.Get("http://127.0.0.1:19001/getnum?srvID=1001658977917")
	r, err := http.Get(addr + "/getnum?srvID=" + fmt.Sprintf("%v", serID))
	if err != nil {
		return "", 0, err
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", 0, err
	}
	defer r.Body.Close()
	rv := &req{}
	json.Unmarshal(b, rv)
	if rv.Size == 18 {
		return sonyflake.BSize18e7, uint16(rv.Num), nil
	}
	if rv.Size == 19 {
		return sonyflake.BSize19e6, uint16(rv.Num), nil
	}
	return "", 0, errors.New("nservID = " + fmt.Sprintf("%v", serID) + " not find ")
}

func New(cfg *Config) (sonyflake.Client,error) {
	size, num, err := getNum(cfg.Addr, cfg.SrvID)
	if err != nil {
		return nil,err
	}
	return sonyflake.InitSonyFlake(&num, size)
}
