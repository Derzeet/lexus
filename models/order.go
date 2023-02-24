package models

import (
	"github.com/jinzhu/gorm"
	u "go-contacts/utils"
)

type Order struct {
	gorm.Model
	GunID      uint    `json:"gun_id"`
	AccountID  uint    `json:"account_id"`
	TotalPrice int     `json:"total_price"`
	Gun        Gun     `gorm:"foreignKey:GunID"`
	Account    Account `gorm:"foreignKey:AccountID"`
}

func (order *Order) CreateOrder() map[string]interface{} {
	GetDB().Create(order)
	resp := u.Message(true, "success")
	resp["orders"] = order
	return resp
}
