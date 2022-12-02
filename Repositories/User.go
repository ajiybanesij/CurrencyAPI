package Repositories

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"model" `
	Name       string `json:"username" gorm:"unique"`
	Hash       string `gorm:"unique"`
	Wallet     []Wallet
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Create() error {
	if err := GetDBInstance().Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) Read(condition map[string]interface{}) error {
	result := GetDBInstance().Where(condition).Find(u)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return nil
}
