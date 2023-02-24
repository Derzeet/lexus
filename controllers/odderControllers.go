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
var GetOrderFor = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	data := models.GetUserOrder(id)
	resp := u.Message(true, "success")
	resp["data"] = data

	u.Respond(w, resp)
}
var EditOrderMethod = func(w http.ResponseWriter, r *http.Request) {
	orderID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		// Handle the case where the ID parameter is not valid
		resp := u.Message(false, "Invalid order ID")
		u.Respond(w, resp)
		return
	}

	// Get the updates to apply to the gun record from the request body
	updateData := &models.Order{}
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		// Handle the case where the request body is not valid
		resp := u.Message(false, "Invalid request body")
		u.Respond(w, resp)
		return
	}

	// Update the gun record in the database
	err = models.EditOrder(uint(orderID), *updateData)
	if err != nil {
		// Handle the case where there was an error updating the gun record
		resp := u.Message(false, "Error updating gun record")
		u.Respond(w, resp)
		return
	}

	// If the gun record was successfully updated, return a success response
	resp := u.Message(true, "Gun record updated successfully")
	u.Respond(w, resp)
}
