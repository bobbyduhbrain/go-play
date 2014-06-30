package devices

import(
  "net/http"
)

type DevicesRoutes struct{}

//
//  Draw routes to publically accessible actions in the devicesController.go file
//
func (devices *DevicesRoutes) Draw(d *DevicesController){
  http.HandleFunc("/devices", d.Index)
}