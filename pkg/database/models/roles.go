package models

type Role struct {
	ID          uint `gorm:"primarykey"`
	Name        string
	Permissions []Permission `gorm:"many2many:role_permissions;constraint:OnDelete:CASCADE;"`
}
