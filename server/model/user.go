package model

import "github.com/google/uuid"

func NewUser(screenName string, iconURL string) (*User, error) {
	user := &User{
		Name:    screenName,
		IconURL: iconURL,
	}

	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByUserUUID(userID uuid.UUID) (*User, error) {
	user := &User{}
	if err := db.Where("id = ?", userID).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByDeviceUUID(deviceID uuid.UUID) (*User, error) {
	user := &User{}
	if err := db.Where("device_users.device_id = ?", deviceID).Joins("LEFT JOIN device_users ON users.id = device_users.user_id").First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
