package Repositories

import (
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

func (b *Balance) ReadBalance(walletId interface{}) (*float64, error) {
	var sum float64
	result := GetDBInstance().Table(b.TableName()).Where("wallet_id = ?", walletId).Select("sum(amount)").Row().Scan(&sum)

	if result != nil {
		return nil, result
	}

	return &sum, nil
}
