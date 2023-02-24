package main

import (
	"fmt"
	"go-contacts/app"
	"go-contacts/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	router := mux.NewRouter()

	router.HandleFunc("/register", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/gun", controllers.CreateGun).Methods("POST")
	router.HandleFunc("/profile", controllers.GetGunsFor).Methods("GET")

	router.HandleFunc("/user/{id}", controllers.GetUserInfo).Methods("GET")
	router.HandleFunc("/order", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/order/{gun_id}", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/guns/{id:[0-9]+}", controllers.EditGunMethod).Methods("PUT")

	router.HandleFunc("/store", controllers.ListStore).Methods("GET")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":8000", c.Handler(router)) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
