package airplay

import(
  "fmt"
  "net/http"
)

type Controller struct{}

//
//  Lists the Airplay compatible devices on the network
//
func (airplay *Controller) Index(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "%s\n", "Listing Airplay compatible devices...")
}
