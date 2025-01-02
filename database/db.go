package database

import (
	"database/sql"
	"fmt"
	"idp-server/utils"
	"os"
	"strings"
)

// TODO: get variables from environment variables instead of hard coded
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
)

type Database struct {
	db *sql.DB
}

var database Database
var logger = utils.GetLogger()

// Creates a connection to DB
// Pings DB for testing connection
// Migrates the DB if tables dont exist
func InitializeDB() error {

	// create db connection
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	db, err := sql.Open("postgres", config)
	if err != nil {
		return err
	}

	// ping db
	err = db.Ping()
	if err != nil {
		return err
	}

	// migrate tables
	content, err := os.ReadFile("./database/sql/create_tables.sql")
	if err != nil {
		return err
	}

	statements := strings.Split(string(content), ";")

	for _, statement := range statements {
		trimmedStatement := strings.TrimSpace(statement)

		if trimmedStatement == "" {
			continue
		}

		_, err := db.Exec(trimmedStatement)
		if err != nil {
			return err
		}
	}

	database = Database{db}

	return nil
}

func GetDB() *Database {
	return &database
}

func (d *Database) RegisterUser(email string, password string) error {
	registerUserQuery := "INSERT INTO users (email, password) VALUES ($1, $2)"

	stmts, err := d.db.Prepare(registerUserQuery)
	if err != nil {
		return err
	}

	defer stmts.Close()

	_, err = stmts.Exec(email, password)
	if err != nil {
		return err
	}

	return nil
}

// for logging without the need for pgadmin
func (d *Database) GetAllUsers() {
	getAllUsersQuery := "SELECT * FROM users"

	stmt, err := d.db.Prepare(getAllUsersQuery)

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
}
