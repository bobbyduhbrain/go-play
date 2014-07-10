package main

import (
  "github.com/bobbyduhbrain/go-play/airplay"
  "time"
)

func main(){
  browser := airplay.NewBrowser()
  for _ = range time.Tick(2 * time.Second) { 
    browser.Poll("_airplay._tcp")
  }
}