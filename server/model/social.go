package model

import "github.com/google/uuid"

func NewSocial(userID uuid.UUID, typeName, identifier string) (*Social, error) {
	social := &Social{
		UserID:     userID,
		Type:       typeName,
		Identifier: identifier,
	}
	if err := db.Create(social).Error; err != nil {
		return nil, err
	}
	return social, nil
}

func GetSocial(typeName, identifier string) (*Social, error) {
	social := &Social{
		Type:       typeName,
		Identifier: identifier,
	}
	if err := db.Where(social).First(social).Error; err != nil {
		return nil, err
	}
	return social, nil
}
