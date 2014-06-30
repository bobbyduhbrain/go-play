package config

//
// Import any packages with a RESTful API interface
//
import(
  "github.com/bobbyduhbrain/go-play/airplay/devices"
)

type Routes struct{}

//
//  Draw the routes for each domain objects in the airplay folder
//  To add new actions, add the route to the {packages}/{packages}Routes.go
//
func (config *Routes) Draw(){
  new(devices.DevicesRoutes).Draw()
}