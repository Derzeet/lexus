package controllers

import (
	"encoding/json"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
)

var CreateGun = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
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
	data := models.GetGun(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
