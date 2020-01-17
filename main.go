package main

import _ "github.com/lib/pq"

import "database/sql"

import "fmt"

func main() {
	db, _ := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable")
	defer db.Close()

	result, _ := db.Exec("INSERT INTO mytable VALUES('shampoo', 420)")
	fmt.Println(result)
}
