package config

//
// Import go-play packages that require publically accessible routes
//
import(
  "github.com/bobbyduhbrain/go-play/airplay/devices"
)

type Routes struct{}

//
//  Draw the routes for all controllers in the airplay package
//
func (config *Routes) Draw(){
  new(devices.DevicesRoutes).Draw(new(devices.DevicesController))
}