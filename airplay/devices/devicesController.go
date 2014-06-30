package devices

import(
  "fmt"
  //"log"
  "net/http"
)

//
//  Uses DNS-SD to detect Airplay devices on the LAN
//
type DevicesController struct{}

//
//  Lists the Airplay compatible devices on the local area network
//
func (devices *DevicesController) Index(w http.ResponseWriter, r *http.Request){
  fmt.Println(w, "Lists all airplay devices")
}