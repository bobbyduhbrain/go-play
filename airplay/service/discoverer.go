package service

import(
  "log"
  "github.com/andrewtj/dnssd"
  "github.com/bobbyduhbrain/go-play/airplay/devices"
)

type Discoverer struct{
  DeviceMap map[string]*devices.Device
}

//
//  Constructor for Discoverer
//
func NewDiscoverer() *Discoverer {
  return &Discoverer{DeviceMap: make(map[string]*devices.Device)}
}

//
// Polls the network for Airplay compatible devices
//
func (d *Discoverer) Poll(){
  op, err := dnssd.StartBrowseOp("_airplay._tcp", d.discovery)
  if err != nil {
    log.Printf("Service discoverer failed on Initialization: %s", err)
    return
  }
  op.Stop()
}

// 
//  DNSSD callback function for changes to services found on the network
//
func (d *Discoverer) discovery(op *dnssd.BrowseOp, err error, device_servicable bool, interfaceIndex int, device_name string, serviceType string, domain string){
  if err != nil {
    log.Printf("Service discoverer failed on discovery: %s", err)
    return
  }
  update_device_map(d.DeviceMap, device_name, device_servicable)
}

//
//  Takes the DeviceMap of a Discoverer and updates the key-value with newly discovered information
//
func update_device_map(dm map[string]*devices.Device, device_name string, device_servicable bool){
  if dev,mapped := dm[device_name]; !mapped {
    log.Printf("Creating new device with key: %v", device_name)
    dm[device_name] = &devices.Device { Data: device_name, Servicable: device_servicable }
  } else {
    log.Printf("Updating device at key %v", device_name)
    dev.Update(device_name, device_servicable)
  }
}
