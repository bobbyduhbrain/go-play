package config

import(
  "net/http"
  "github.com/bobbyduhbrain/go-play/lib/airplay/devices"
)

type Routes struct{}

func (config *Routes) Draw(){
  http.HandleFunc("/devices", new(airplay.DevicesController).Index)
}