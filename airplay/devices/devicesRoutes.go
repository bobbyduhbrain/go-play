package devices

import(
  "net/http"
)

type DevicesRoutes struct{}

//
//  Draw routes to an action in the devicesController.go file
//
func (devices *DevicesRoutes) Draw(){
  http.HandleFunc("/devices", new(DevicesController).Index)
}