package service_discovery

import(
  "fmt"
  "log"
  "github.com/andrewtj/dnssd"
  "github.com/bobbyduhbrain/go-play/airplay/devices"
)

type Discoverer struct{}

func (d *Discoverer) Initialize(){
  op, err := dnssd.StartBrowseOp("_airplay._tcp", d.discover)
  if err != nil {
    // op is now inactive
    log.Printf("Browse operation failed: %s", err)
    return
  }
  // later...
  op.Stop()
}

func (d *Discoverer) discover(op *dnssd.BrowseOp, err error, add bool, interfaceIndex int, name string, serviceType string, domain string){
  dev := devices.Device{ Data: name}
  fmt.Println(dev.Data)
}

