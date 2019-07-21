package model

import "github.com/google/uuid"

func NewDeviceUser(deviceID, userID uuid.UUID) (*DeviceUser, error) {
	deviceUser := &DeviceUser{
		DeviceID: deviceID,
		UserID:   userID,
	}
	if err := db.Create(deviceUser).Error; err != nil {
		return nil, err
	}
	return deviceUser, nil
}
