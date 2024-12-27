package main

import (
	"fmt"
	"net/http"

	"idp-server/handlers"
	"idp-server/utils"
)

var hostPort = ":2345"

func main() {

	logger := utils.GetLogger()

	logger.Info("Creating server mux")

	mux := http.NewServeMux()

	logger.Info("Creating routes")

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/", handlers.Error)

	mux.HandleFunc("/register", handlers.Register)

	mux.HandleFunc("/login", handlers.Login)

	logger.Info(fmt.Sprintf("Starting server, listening on '%s'", hostPort))
	if err := http.ListenAndServe(hostPort, mux); err != nil {
		panic(err)
	}
}
