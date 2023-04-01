package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"gorm_._model"`
	Name       string `json:"name,omitempty" gorm:"not null"`
	Email      string `json:"email,omitempty" gorm:"unique;not null'"`
	Password   string `json:"password,omitempty" gorm:"not null"`
	Address    string `json:"address,omitempty" gorm:"not null"`
}

func (u *User) TableName() string {
	return "users"
}
