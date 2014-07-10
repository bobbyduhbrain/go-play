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
func (d *Discoverer) Poll(service_type string){
  poll, err := dnssd.StartBrowseOp(service_type, d.discovery)
  if err != nil {
    log.Printf("Service discoverer failed on Initialization: %s", err)
    return
  }
  poll.Stop()
}

// 
//  DNSSD callback function for changes to services found on the network
//
func (d *Discoverer) discovery(op *dnssd.BrowseOp, err error, device_servicable bool, interfaceIndex int, device_name string, serviceType string, domain string){
  if err != nil {
    log.Printf("Service discoverer failed on discovery: %s", err)
    return
  }
  resolve, err := dnssd.StartResolveOp(interfaceIndex, device_name, serviceType, domain, d.resolve_discovery)
  if err!= nil {
    log.Printf("Failed to resolve service discovery: %s", err)
  }
  resolve.Stop()
}


//
//  DNSSD callback function for resolving service discovery
//
func (d *Discoverer) resolve_discovery(op *dnssd.ResolveOp, err error, host string, port int, metadata map[string]string){
  if err != nil {
    log.Printf("Resolving discovery failed: %s", err)
    return
  }
  log.Printf("Resolved service to host %s port %d with meta info: %v", host, port, metadata)
  update_device_map(d.DeviceMap, op.Name(), metadata)
}

//
//  Takes the DeviceMap of a Discoverer and updates the key-value with newly discovered information
//
func update_device_map(dm map[string]*devices.Device, device_name string, device_metadata map[string]string){
  if dev,mapped := dm[device_name]; !mapped {
    log.Printf("Creating new device with key: %v", device_name)
    dm[device_name] = &devices.Device { Data: device_metadata}
  } else {
    log.Printf("Updating device at key %v", device_name)
    dev.Update(device_metadata)
  }
}
