package main

import (
	"github.com/gorilla/mux"
	"lld-tdd/controllers"
	"lld-tdd/service"
	"log"
	"net/http"
)

func main() {

	db, err := service.ConnectDB()
	if err != nil {
		panic("Failed to connect to database")
	}
	log.Print("Connected to DB")
	service.MigrateDB(db)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/signup", controllers.Signup).Methods("POST")

	http.ListenAndServe(":8080", router)
}
