package devices

import(
  "net/http"
)

type DevicesRoutes struct{}

func (devices *DevicesRoutes) Draw(){
  http.HandleFunc("/devices", new(DevicesController).Index)
}