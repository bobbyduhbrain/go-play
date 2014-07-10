package devices

type Device struct{
  Data map[string]string
}

func (d *Device) Update(data map[string]string){
  d.Data = data
}