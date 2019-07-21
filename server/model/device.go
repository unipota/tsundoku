package model

func NewDevice() *Device {
	device := &Device{}
	db.Create(device)
	return device
}
