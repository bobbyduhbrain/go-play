package devices

import(
  "fmt"
  "net/http"
)

type DevicesController struct{}

//  Public:
//
//  Lists the Airplay compatible devices on the local area network
//
func (devices *DevicesController) Index(w http.ResponseWriter, r *http.Request){
  fmt.Println(w, "Lists all airplay devices")
}