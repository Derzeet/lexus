package models

import (
	"fmt"
	u "go-contacts/utils"

	"github.com/jinzhu/gorm"
)

type Gun struct {
	gorm.Model
	Name         string `json:"name"`
	Modelka      string `json:"model"`
	Caliber      string `json:"caliber"`
	Price        int    `json:"price"`
	Availability bool   `json:"availability"`
	Type         string `json:"type"`
	UserId       uint   `json:"user_id"` //The user that this contact belongs to
}

func (gun *Gun) Validate() (map[string]interface{}, bool) {

	if gun.Name == "" {
		return u.Message(false, "Gun name should be on the payload"), false
	}

	if gun.Modelka == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}

	if gun.Caliber == "" {
		return u.Message(false, "User is not recognized"), false
	}

	if gun.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}
	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (gun *Gun) Create() map[string]interface{} {

	if resp, ok := gun.Validate(); !ok {
		return resp
	}

	GetDB().Create(gun)

	resp := u.Message(true, "success")
	resp["contact"] = gun
	return resp
}

func GetGun(id uint) *Gun {

	gun := &Gun{}
	err := GetDB().Table("guns").Where("id = ?", id).First(gun).Error
	if err != nil {
		return nil
	}
	return gun
}

func GetGuns(user uint) []*Gun {

	gun := make([]*Gun, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&gun).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return gun
}