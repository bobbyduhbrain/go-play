package config

//
// Import go-play packages that require publically accessible routes
//
import(
  "github.com/bobbyduhbrain/go-play/airplay/devices"
)

type Routes struct{}

//
//  Draw the routes for each package in the airplay folder
//  To add new actions, add the route to the {packages}/{packages}Routes.go
//
func (config *Routes) Draw(){
  devices_controller := new(devices.DevicesController)
  devices_router     := new(devices.DevicesRoutes)
  devices_router.Draw(devices_controller)
}