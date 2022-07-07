package main

import (
	"encoding/json"
	"github.com/et-zone/eids/conf"
	"github.com/et-zone/eids/public"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
	id, err := public.GetCli().NextID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	body, err := json.Marshal(public.Decompose(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header()["Content-Type"] = []string{"application/json; charset=utf-8"}
	w.Write(body)
}

func main() {
	conf.InitConf()
	public.InitIDS()
	http.HandleFunc("/ids", handler)
	http.ListenAndServe(":9001", nil)
}