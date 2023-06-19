package main

import (
	"github.com/gorilla/mux"
	"lld-tdd/controllers"
	"lld-tdd/service"
	"net/http"
)

func main() {
	service.Connect()
	service.Migrate()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/signup", controllers.Signup).Methods("POST")

	http.ListenAndServe(":8080", router)
}
