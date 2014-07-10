package service

import(
  "log"
  "github.com/andrewtj/dnssd"
  "github.com/bobbyduhbrain/go-play/airplay/devices"
)

type Browser struct{
  DeviceMap map[string]*devices.Device
}

func NewBrowser() *Browser {
  return &Browser{DeviceMap: make(map[string]*devices.Device)}
}

//
// Polls the network for compatible devices for a specific service type
//
func (b *Browser) Poll(service_type string){
  poll, err := dnssd.StartBrowseOp(service_type, b.browse)
  if err != nil {
    log.Printf("Service browser failed on Poll: %s", err)
    return
  }
  poll.Stop()
}

// 
//  DNSSD callback function for changes to services found on the network
//
func (b *Browser) browse(op *dnssd.BrowseOp, err error, device_servicable bool, interfaceIndex int, device_name string, serviceType string, domain string){
  if err != nil {
    log.Printf("Service browser failed on browse: %s", err)
    return
  }
  resolve, err := dnssd.StartResolveOp(interfaceIndex, device_name, serviceType, domain, b.resolve_service_discovery)
  if err!= nil {
    log.Printf("Service browser failed on resolve service discovery: %s", err)
  }
  resolve.Stop()
}


//
//  DNSSD callback function for resolving service discovery
//
func (b *Browser) resolve_service_discovery(op *dnssd.ResolveOp, err error, host string, port int, metadata map[string]string){
  if err != nil {
    log.Printf("Resolving discovery failed: %s", err)
    return
  }
  update_device_map(b.DeviceMap, op.Name(), metadata)
}

//
//  Takes the DeviceMap of a Browser and updates the key-value with newly discovered information
//
func update_device_map(dm map[string]*devices.Device, device_name string, device_metadata map[string]string){
  if dev,mapped := dm[device_name]; !mapped {
    dm[device_name] = &devices.Device { Data: device_metadata}
  } else {
    dev.Update(device_metadata)
    log.Printf("%s: %v", device_name, dev.Data["features"])
  }
}
