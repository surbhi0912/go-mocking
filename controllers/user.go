package controllers

import (
	"encoding/json"
	"lld-tdd/models"
	"lld-tdd/service"
	"net/http"
)

func SignupUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user) //decodes body of incoming json request and maps it to newly created user variable

	err := service.CreateNewUser(&user)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode("User signed up successfully")
	}
}
