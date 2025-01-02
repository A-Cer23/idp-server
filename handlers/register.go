package handlers

import (
	"idp-server/database"
	"idp-server/utils"
	"net/http"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var logger = utils.GetLogger()
var db = database.GetDB()

func Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		w.Header().Set("Content-Type", "text/html")
		htmlContent, _ := os.ReadFile("./html/register.html")
		w.Write(htmlContent)

	case "POST":
		err := r.ParseForm()
		if err != nil {
			logger.Error("error", err)
		}
		email := r.FormValue("email")
		// TODO: hash password server side
		password := r.FormValue("password")

		err = db.RegisterUser(email, password)

		if err != nil {
			logger.Error("Issue registering user", err)
		}

		w.Write([]byte("success"))
	}

}
