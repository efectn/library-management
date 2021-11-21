package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
	Phone    string
	City     string
	State    string
	Country  string
	ZipCode  int
	Address  string
	Roles    []Role `gorm:"many2many:user_roles;constraint:OnDelete:CASCADE;"`
}
