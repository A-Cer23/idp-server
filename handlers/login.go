package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO: only post requests
	var loginRequest LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginRequest)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Wrong data in body err: %v", err)))
		return
	}
	w.Write([]byte(fmt.Sprintf("loginRequest: %v", loginRequest)))
}
