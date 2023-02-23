package controllers

import (
	"encoding/json"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var CreateGun = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	userProps := r.Context().Value("user")

	account, ok := userProps.(models.Account)
	if !ok {
		// Handle the case where userProps is not of type Account
	}

	isSeller := account.Seller

	if !isSeller {
		u.Respond(w, u.Message(false, "You are not seller"))
		return
	}

	gun := &models.Gun{}

	err := json.NewDecoder(r.Body).Decode(gun)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	gun.UserId = user
	resp := gun.Create()
	u.Respond(w, resp)
}

var GetGunsFor = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	data := models.GetUserGuns(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var EditGunMethod = func(w http.ResponseWriter, r *http.Request) {
	// Get the gun ID from the request URL parameters
	gunID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		// Handle the case where the ID parameter is not valid
		resp := u.Message(false, "Invalid gun ID")
		u.Respond(w, resp)
		return
	}

	// Get the updates to apply to the gun record from the request body
	updateData := &models.Gun{}
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		// Handle the case where the request body is not valid
		resp := u.Message(false, "Invalid request body")
		u.Respond(w, resp)
		return
	}

	// Update the gun record in the database
	err = models.EditGun(uint(gunID), *updateData)
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
