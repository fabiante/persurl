package models

import "gorm.io/gorm"

type Domain struct {
	gorm.Model

	Name string

	PURLs []*PURL `gorm:"foreignKey:DomainID"`
}

type PURL struct {
	gorm.Model

	DomainID uint

	Name   string
	Target string
}
