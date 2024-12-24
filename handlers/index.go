package handlers

import (
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	htmlContent, _ := os.ReadFile("./html/index.html")
	w.Write(htmlContent)
}
