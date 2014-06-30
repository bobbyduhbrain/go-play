package devices

import(
  "fmt"
  "net/http"
  //"github.com/andrewtj/dnssd"
)

//
//  Uses DNS-SD to detect Airplay devices on the LAN
//
type DevicesController struct{}

//
//  Draw Routes for DevicesController actions
//
func (airplay *DevicesController) Routes(){
  http.HandleFunc("/devices", new(DevicesController).Index)
}


//
//  Lists the Airplay compatible devices on the local area network
//
func (airplay *DevicesController) Index(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "%s\n", "Listing Airplay compatible devices...")
}
