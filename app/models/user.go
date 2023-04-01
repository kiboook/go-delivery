package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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

func (u *User) BeforeSave(*gorm.DB) error {
	hashedPassword, err := hashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword
	return nil
}

// hashPassword using bcrypt
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
