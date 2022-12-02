package Repositories

import (
	"errors"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Currency    string
	UserID      uint
	Balance     []Balance
	Transaction []Transaction
}

func (w *Wallet) TableName() string {
	return "wallets"
}

func (w *Wallet) Create() error {
	if err := GetDBInstance().Create(w).Error; err != nil {
		return err
	}
	return nil
}

func (w *Wallet) Read(condition map[string]interface{}) error {
	result := GetDBInstance().Where(condition).Find(w)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return nil
}
