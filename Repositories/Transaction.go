package Repositories

import (
	"errors"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	FromCurrency string
	ToCurrency   string
	InAmount     float64
	OutAmount    float64
	Rate         float64
	WalletID     uint
}

func (t *Transaction) TableName() string {
	return "transactions"
}

func (t *Transaction) Create() error {
	if err := GetDBInstance().Create(t).Error; err != nil {
		return err
	}
	return nil
}

func (t *Transaction) Read(condition map[string]interface{}) error {
	result := GetDBInstance().Where(condition).Find(t)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return nil
}
