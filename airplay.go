package main

import (
  "github.com/bobbyduhbrain/go-play/airplay"
  
  "flag"
  "time"
)

var port int

func init() {
  flag.IntVar(&port, "port", 8080, "port to run the server on")
  flag.Parse()
}

func main(){
  browser := airplay.NewBrowser()
  for _ = range time.Tick(2 * time.Second) { 
    browser.Poll("_airplay._tcp")
  }
}