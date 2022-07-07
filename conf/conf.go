package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)
var ConMsg *Conf
type Conf struct {
	ServID int `json:"servID"`
	Port int `json:"port"`
}

func InitConf()error{
	b,err:=ioutil.ReadFile("config.json")
	if err!=nil{
		log.Fatalln(err)
	}
	tmp:=&Conf{}
	err=json.Unmarshal(b,tmp)
	if err!=nil{
		log.Fatalln(err)
	}
	ConMsg=tmp
	return err
}