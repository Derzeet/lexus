package main

import (
	"fmt"
	"go-contacts/app"
	"go-contacts/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/register", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/gun", controllers.CreateGun).Methods("POST")
	router.HandleFunc("/profile", controllers.GetGunsFor).Methods("GET")
	router.HandleFunc("/order", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/order/{gun_id}", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/guns/{id:[0-9]+}", controllers.EditGunMethod).Methods("PUT")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
