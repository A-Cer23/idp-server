package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"

	"idp-server/handlers"
	"idp-server/utils"
)

var hostPort = ":2345"

func main() {

	logger := utils.GetLogger()

	// // // // // // // // // // // // // // // // // // //
	// TODO: TO BE MOVED TO SEPERATE FILE
	logger.Info("Instantiating a connection to database")

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres sslmode=disable")
	if err != nil {
		logger.Error("Problem connecting to DB", err)
		return
	}

	err = db.Ping()
	if err != nil {
		logger.Error("Problem pinging DB", err)
		return
	}

	logger.Info("Connection established to database")

	logger.Info("Checking if users table exists")

	content, err := os.ReadFile("./db/sql/create_tables.sql")
	if err != nil {
		logger.Error("Problem reading create_tables.sql", err)
		return
	}

	statements := strings.Split(string(content), ";")

	for _, statement := range statements {
		trimmedStatement := strings.TrimSpace(statement)
		if trimmedStatement == "" {
			continue
		}
		_, err := db.Exec(trimmedStatement)
		if err != nil {
			logger.Error("Issue executing statement", err)
		}
	}

	stmts, _ := db.Prepare("INSERT INTO users (username, password) VALUES ($1, $2)")

	defer stmts.Close()

	stmts.Exec("testuser", "123456")

	getAllUsersQuery := "SELECT * FROM users"

	stmt, err := db.Prepare(getAllUsersQuery)

	if err != nil {
		logger.Error("Can't create getAllUser stmt", err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		logger.Error("Problem executing query", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id       []uint8
			uname    string
			password string
		)

		err := rows.Scan(&id, &uname, &password)
		if err != nil {
			logger.Error("Issue scanning row", err)
		}

		logger.Info(fmt.Sprintf("ID: %s, Username: %s, Password: %s", id, uname, password))
	}
	// // // // // // // // // // // // // // // // // // //

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
