package controllers

import (
	"encoding/json"
	"lld-tdd/models"
	"lld-tdd/service"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user) //decodes body of incoming json request and maps it to newly created user variable

	response := service.CreateNewUser(user)
	json.NewEncoder(w).Encode(response)
}
