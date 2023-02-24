package models

import (
	"fmt"
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
func GetOrder(id uint) (*Order, error) {

	order := &Order{}
	err := GetDB().Table("orders").Where("id = ?", id).First(order).Error
	if err != nil {
		return nil, nil
	}
	return order, nil
}

func GetUserOrder(user uint) []*Order {
	order := make([]*Order, 0)
	err := GetDB().Table("orders").Where("account_id = ?", user).Find(&order).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return order
}

func EditOrder(orderID uint, updates Order) error {
	// Retrieve the gun record from the database
	var order Order
	result := GetDB().First(&order, orderID)
	if result.Error != nil {
		return result.Error
	}

	// Update the gun record with the provided updates
	result = db.Model(&order).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
