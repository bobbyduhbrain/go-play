package main

import (
  "flag"
  "fmt"
  "net/http"

  "github.com/bobbyduhbrain/go-play/airplay"
)

var port int

func init() {
  flag.IntVar(&port, "port", 8080, "port to run the server on")
  flag.Parse()
}

func main(){
  drawRoutes()
  http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

//
//  Draw the service routes
//
func drawRoutes(){
  http.HandleFunc("/devices", new(airplay.Controller).Index)
}




