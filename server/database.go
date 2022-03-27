package server

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func openDatabase() {
	db, err := sql.Open("sqlite3", "yrc.db")
	if err != nil {
		log.Fatal(err)
	}
	database = db
}

func getUserByUsername(username string) (userModel, error) {

	rows, err := database.Query(fmt.Sprintf("SELECT * FROM Users WHERE Username = '%s'", username))
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	var user userModel

	for rows.Next() {
		err = rows.Scan(&user.id, &user.username, &user.password)
		return user, err
	}

	return user, errors.New(fmt.Sprintf(`User "%s" not found`, username))
}
