package handlers

import (
	"net/http"
	"os"
)

func Error(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	htmlContent, _ := os.ReadFile("./html/404.html")
	w.Write(htmlContent)
}
