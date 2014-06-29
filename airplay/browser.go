package airplay

import(
  "fmt"
  "net/http"
)

type Browser struct{}

//
//  Lists the Airplay compatible devices on the network
//
func (airplay *Browser) List(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "%s\n", "Listing Airplay compatible devices...")
}
