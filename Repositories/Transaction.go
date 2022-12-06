package Repositories

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	FromCurrency string
	FromAmount   float64
	FromWalletID uint
	ToCurrency   string
	ToAmount     float64
	ToWalletID   uint
	Rate         float64
	UserID       uint
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

func (t *Transaction) Read(condition map[string]interface{}) (*[]Transaction, error) {
	var transactions []Transaction
	err := GetDBInstance().Where(condition).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return &transactions, nil

}
