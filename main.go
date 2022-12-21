package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	controller "github.com/haseebarifseecs/rest-api/controllers"
	middleware "github.com/haseebarifseecs/rest-api/middleware"
)

func listener(r *mux.Router) {
	server := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8081",
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server Started Listening on Port 8081")

}

func main() {
	router := mux.NewRouter()
	// router.PathPrefix("/user")
	// router.HandleFunc("/create", controller.CreateHandler).Methods("GET").PathPrefix("/user")
	// r.PathPrefix("/user/")
	subRouter := router.PathPrefix("/user").Subrouter()
	subRouter.HandleFunc("/create", controller.CreateHandler).Methods("POST")
	subRouter.HandleFunc("/get", controller.ReadHandler).Methods("GET")
	subRouter.HandleFunc("/update", controller.UpdateHandler).Methods("PUT")
	subRouter.HandleFunc("/delete", controller.DeleteHandler).Methods("DELETE")
	// createRoute := router.HandleFunc("/create", controller.CreateHandler).Methods("GET")
	// createRoute.PathPrefix("/user").HandlerFunc(controller.CreateHandler)
	// updateRoute := router.HandleFunc("/update", controller.UpdateHandler)
	// updateRoute.PathPrefix("/user").HandlerFunc(controller.UpdateHandler)
	// deleteRoute := router.HandleFunc("/delete", controller.DeleteHandler)
	// deleteRoute.PathPrefix("/user").HandlerFunc(controller.DeleteHandler)
	// readRoute := router.HandleFunc("/get", controller.ReadHandler)
	// readRoute.PathPrefix("/user").HandlerFunc(controller.ReadHandler)
	router.NotFoundHandler = http.HandlerFunc(middleware.Logs)

	listener(router)
}
