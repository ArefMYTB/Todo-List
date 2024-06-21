package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	DB_HOST string = "127.0.0.1"
	DB_USER string = "root"
	DB_PASS string = ""
)

func Database() *sql.DB {

	credentials := fmt.Sprintf(
		"%s:%s@(%s:4040)/?charset=utf8&parseTime=True",
		DB_USER, DB_PASS, DB_HOST)

	database, err := sql.Open("mysql", credentials)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database Connection Successful")
	}

	if _, err = database.Exec(`CREATE TABLE TODO`); err != nil {
		log.Fatal(err)
	}
	if _, err = database.Exec(`USE TODO`); err != nil {
		log.Fatal(err)
	}
	if _, err = database.Exec(`
		CREATE TABLE todos (
		    id INT AUTO_INCREMENT,
		    item TEXT NOT NULL,
		    completed BOOLEAN DEFAULT FALSE,
		    PRIMARY KEY (id)
		);
	`); err != nil {
		log.Fatal(err)
	}

	return database
}
