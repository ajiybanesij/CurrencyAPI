package Repositories

import (
	"errors"
	"gorm.io/gorm"
)

type Balance struct {
	gorm.Model
	Currency string
	Amount   float64
	WalletID uint
}

func (b *Balance) TableName() string {
	return "balance"
}

func (b *Balance) Create() error {
	if err := GetDBInstance().Create(b).Error; err != nil {
		return err
	}
	return nil
}

func (b *Balance) Read(condition map[string]interface{}) error {
	result := GetDBInstance().Where(condition).Find(b)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return nil
}
