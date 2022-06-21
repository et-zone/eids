package main

import (
	"encoding/json"
	"github.com/et-zone/eids/public"
	"github.com/sony/sonyflake"
	"net/http"
)

var sf *sonyflake.Sonyflake

func init() {

}

func handler(w http.ResponseWriter, r *http.Request) {
	id, err := sf.NextID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(sonyflake.Decompose(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header()["Content-Type"] = []string{"application/json; charset=utf-8"}
	w.Write(body)
}

func main() {
	var st sonyflake.Settings
	st.MachineID = public.MachineID
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
	http.HandleFunc("/ids", handler)
	http.ListenAndServe(":9001", nil)
}