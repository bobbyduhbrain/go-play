package devices

type Device struct{
  Data string
  Servicable bool
}

func (d *Device) Update(data string, servicable bool){
  d.Data = data
  d.Servicable = servicable
}