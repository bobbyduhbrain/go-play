package service

import(
  "log"
  "github.com/andrewtj/dnssd"
  "github.com/bobbyduhbrain/go-play/airplay/devices"
)

type Browser struct{
  DeviceMap map[string]*devices.Device
}

//  Public:
//
//  Returns a pointer to a new Browser
//
func NewBrowser() *Browser {
  return &Browser{DeviceMap: make(map[string]*devices.Device)}
}

//  Public:
//
//  Polls the network for compatible devices for a specific service type
//
func (b *Browser) Poll(service_type string){
  poll, err := dnssd.StartBrowseOp(service_type, b.browse)
  if err != nil {
    log.Printf("Service browser failed on Poll: %s", err)
    return
  }
  poll.Stop()
}

//  Private:
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


//  Private:
//
//  DNSSD callback function for resolving service discovery
//
func (b *Browser) resolve_service_discovery(op *dnssd.ResolveOp, err error, host string, port int, txt map[string]string){
  if err != nil {
    log.Printf("Resolving discovery failed: %s", err)
    return
  }
  update_device_map(b.DeviceMap, op.Name(), host, port, txt)
}

//  Private:
//
//  Takes the DeviceMap of a Browser and updates the key-value with newly discovered information
//
func update_device_map(dm map[string]*devices.Device, name string, host string, port int, txt map[string]string){
  if dev,mapped := dm[name]; !mapped {
    dm[name] = &devices.Device { Host: host, Port: port, ServerInfo: txt}
  } else {
    dev.Update(host, port, txt)
    log.Printf("%v", *dev)
  }
}
