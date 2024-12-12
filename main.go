package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var hostPort = ":2345"

var methods = map[string]bool{
	"TOTP": true,
}

var allowedMethods = []string{
	"TOTP",
}

type LoginMethodError struct {
	Error          string   `json:"error"`
	AllowedMethods []string `json:"allowed-methods"`
}

func main() {

	logger.Info("Creating server mux")

	mux := http.NewServeMux()

	logger.Info("Creating routes")

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>IDP Server</h1>"))
	})

	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/signup"))
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		method := r.Header.Get("Method")

		// if method header missing
		if method == "" {
			e := LoginMethodError{"Missing `Method` header", allowedMethods}
			b, err := json.Marshal(e)
			if err != nil {
				logger.Error("Error marshaling in method header missing `/login`", err)
			}
			w.WriteHeader(400)
			w.Write([]byte(b))
			return
		}

		// if method header value incorrect
		if _, ok := methods[method]; !ok {
			e := LoginMethodError{"Incorrect `Method` header value", allowedMethods}
			b, err := json.Marshal(e)
			if err != nil {
				logger.Error("Error marshaling in incorrect method header value `/login`", err)
			}
			w.WriteHeader(400)
			w.Write([]byte(b))
			return
		}

		w.Write([]byte("/login"))
	})

	logger.Info(fmt.Sprintf("Starting server, listening on '%s'", hostPort))
	if err := http.ListenAndServe(hostPort, mux); err != nil {
		panic(err)
	}
}
