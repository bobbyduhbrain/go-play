package config

import(
  "github.com/bobbyduhbrain/go-play/lib/airplay/devices"
)

type Routes struct{}

func (config *Routes) Draw(){
 new(devices.DevicesRoutes).Draw()
}