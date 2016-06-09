package main

import (
	"log"
	"net/http"

	"github.com/ljcastro/taskmanager/common"
	"github.com/ljcastro/taskmanager/routers"
	"github.com/urfave/negroni"
)

// Entry point of the program
func main() {

	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()
	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
