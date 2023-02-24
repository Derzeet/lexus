package controllers

import (
	"encoding/json"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var CreateOrder = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gunID, err := strconv.Atoi(params["gun_id"])
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid gun ID"))
		return
	}

	user := r.Context().Value("user").(uint) // Grab the id of the user that sent the request

	order := &models.Order{}

	err = json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	order.AccountID = user
	order.GunID = uint(gunID)
	account := models.GetUser(order.AccountID)
	if account == nil {
		order.Account = *account
	}
	gun, err := models.GetGun(order.GunID)
	if err == nil {
		order.Gun = *gun
		order.TotalPrice = 1000 + gun.Price

	}
	resp := order.CreateOrder()
	u.Respond(w, resp)
}
