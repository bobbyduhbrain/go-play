package devices

import(
  "net/http"
)

type DevicesRoutes struct{}

//
//  Draw routes to an action in the devicesController.go file
//
func (devices *DevicesRoutes) Draw(d *DevicesController){
  devices_controller := new(DevicesController)
  http.HandleFunc("/devices", devices_controller.Index)
}