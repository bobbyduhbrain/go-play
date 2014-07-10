package main

import (
  "github.com/bobbyduhbrain/go-play/airplay/config"
  "github.com/bobbyduhbrain/go-play/airplay/service"
  
  "flag"
  "fmt"
  "net/http"
  "time"
)

var port int

func init() {
  flag.IntVar(&port, "port", 8080, "port to run the server on")
  flag.Parse()
  browser := service.NewBrowser()
  for _ = range time.Tick(2 * time.Second) { 
    browser.Poll("_airplay._tcp")
  }
}

func main(){
  new(config.Routes).Draw()                             //  Draw the HTTP accessible routes 
  http.ListenAndServe(fmt.Sprintf(":%d", port), nil)    //  Begin listening for requests on port
}