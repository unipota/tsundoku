package model

import "github.com/google/uuid"

func NewUser(screenName string) (*User, error) {
	user := &User{
		Name: screenName,
	}

	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByID(userID uuid.UUID) (*User, error) {
	user := &User{}
	if err := db.Where("id = ?", userID).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
