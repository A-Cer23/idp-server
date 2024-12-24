package handlers

import (
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/signup"))
}
