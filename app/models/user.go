package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email string

	Keys []*UserKey `gorm:"foreignKey:OwnerID"`
}

type UserKey struct {
	gorm.Model

	OwnerID uint

	Value string
}
