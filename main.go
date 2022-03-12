package main

import (
	"database/sql"
	"fmt"
	"log"
	// "github.com/go-sql-driver/mysql" # This is MySQL driver for go
	// "github.com/mattn/go-sqlite3" # This is SQLite driver for go
)

// To install the database driver, run
// go get github.com/mattn/go-sqlite3
// or, go get github.com/go-sql-driver/mysql

// For other database usage, please visit https://github.com/golang/go/wiki/SQLDrivers

func main() {
	db, err := sql.Open("mysql", "root:password1@tcp(127.0.0.1:3306)/test")
	// Setup the database connection, if use sqlite, use the relative .db path
	if err != nil {
		log.Fatal("Unable to open connection to db")
	}

	results, err := db.Query("select * from product")
	// Use Query func to query data in SQL statement
	if err != nil {
		log.Fatal("Error when fetching product table rows:", err)
	}
	for results.Next() {
		var (
			id    int
			name  string
			price int
		)
		err = results.Scan(&id, &name, &price)
		// Matching the query result to the declared variable
		if err != nil {
			log.Fatal("Unable to parse row:", err)
		}
		fmt.Printf("ID: %d, Name: '%s', Price: %d\n", id, name, price)
	}

	var (
		id    int
		name  string
		price int
	)
	err = db.QueryRow("Select * from product where id = 1").Scan(&id, &name, &price)
	// Use QuerrRow func to get one row only
	if err != nil {
		log.Fatal("Unable to parse row:", err)
	}
	fmt.Printf("ID: %d, Name: '%s', Price: %d\n", id, name, price)

	products := []struct {
		name  string
		price int
	}{
		{"Light", 10},
		{"Mic", 30},
		{"Router", 90},
	}

	stmt, err := db.Prepare("INSERT INTO product (name, price) VALUES (?, ?)")
	// Use Prepare func to insert, delete or update data
	if err != nil {
		log.Fatal("Unable to prepare statement:", err)
	}
	for _, product := range products {
		_, err = stmt.Exec(product.name, product.price)
		if err != nil {
			log.Fatal("Unable to execute statement:", err)
		}
	}
}
