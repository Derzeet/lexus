package controllers

import (
	"encoding/json"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create() //Create account
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}

var GetUserInfo = func(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the request parameters
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid user ID"))
		return
	}

	// Retrieve the user record from the database
	user := models.GetUser(uint(userID))
	if user == nil {
		u.Respond(w, u.Message(false, "User not found"))
		return
	}

	// Remove the password field from the user record
	user.Password = ""

	// Convert the user record to JSON and send it back as the response
	resp := u.Message(true, "User found")
	resp["user"] = user
	u.Respond(w, resp)
}
