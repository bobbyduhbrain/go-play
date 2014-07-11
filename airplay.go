package main

import (
  "github.com/bobbyduhbrain/go-play/airplay"
  "github.com/codegangsta/cli"
  "os"
  "time"
)

func main(){
  app := cli.NewApp()
  app.Name = "go-play"
  app.Usage = "Stream videos to Airplay compatible devices."
  app.Commands = []cli.Command{
    {
      Name:      "listen",
      ShortName: "l",
      Usage:     "listen for Airplay compatible devices on the network",
      Action: func(c *cli.Context) {
        Listen()
      },
    },
  }
  app.Run(os.Args)
}

func Listen(){
  browser := airplay.NewBrowser()
  for _ = range time.Tick(2 * time.Second) { 
    browser.Poll("_airplay._tcp")
  }
}