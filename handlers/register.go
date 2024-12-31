package handlers

import (
	"net/http"
	"os"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// GET - display form
	// POST - store info to db
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "text/html")
		htmlContent, _ := os.ReadFile("./html/register.html")
		w.Write(htmlContent)
	case "POST":
		// store username and pass to db
		w.Write([]byte("POST /register"))
	}

}
