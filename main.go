package main

import (
	"database/sql"
	"fmt"

	_ "github.com/marcboeker/go-duckdb"
)

func main() {
	// Open a connection to the database
	db, err := sql.Open("duckdb", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Create a table
	_, err = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name VARCHAR(50), age INTEGER)")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Insert some data
	_, err = db.Exec("INSERT INTO users (id, name, age) VALUES (1, 'Alice', 25), (2, 'Bob', 30), (3, 'Charlie', 35)")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Query the data
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	// Print the results
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("id: %d, name: %s, age: %d\n", id, name, age)
	}
}
