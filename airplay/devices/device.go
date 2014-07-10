package devices

type Device struct{
  Host string
  Port int
  ServerInfo map[string]string
}

func (d *Device) Update(host string, port int, si map[string]string){
  d.Host = host
  d.Port = port
  d.ServerInfo = si
}