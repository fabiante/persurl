package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email string

	Domains []*Domain `gorm:"foreignKey:OwnerID"`

	Keys []*UserKey `gorm:"foreignKey:OwnerID"`
}

type UserKey struct {
	gorm.Model

	OwnerID uint

	Value string
}
