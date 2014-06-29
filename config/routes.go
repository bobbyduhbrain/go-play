package config

import(
  "net/http"
  "github.com/bobbyduhbrain/go-play/airplay"
)

type Routes struct{}

func (config *Routes) Draw(){
  http.HandleFunc("/devices", new(airplay.Browser).List)
}