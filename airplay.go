package main

import (
  "flag"
  "fmt"
  "net/http"
  "github.com/bobbyduhbrain/go-play/config"
)

var port int

func init() {
  flag.IntVar(&port, "port", 8080, "port to run the server on")
  flag.Parse()
}

func main(){
  new(config.Routes).Draw()                             //  Draw the HTTP accessible routes 
  http.ListenAndServe(fmt.Sprintf(":%d", port), nil)    //  Begin listening for requests on port
}




