package Repositories

import (
	"errors"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Currency string
	UserID   uint
	Balance  []Balance
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
	result := GetDBInstance().Where(condition).Preload("Balance").Find(w)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return nil
}

func (w *Wallet) ReadAll(condition map[string]interface{}) ([]Wallet, error) {
	var wallets []Wallet
	err := GetDBInstance().Where(condition).Preload("Balance").Find(&wallets).Error
	if err != nil {
		return nil, err
	}
	return wallets, nil
}
