package model

import "github.com/google/uuid"

func NewDevice() *Device {
	device := &Device{}
	db.Create(device)
	return device
}

func GetDeviceByDeviceID(deviceID uuid.UUID) (*Device, error) {
	device := &Device{
		Base: Base{
			ID: deviceID,
		},
	}

	if err := db.Where(device).First(device).Error; err != nil {
		return nil, err
	}
	return device, nil
}
