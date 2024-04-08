package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:Djm@1211@/mydatabase")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		name := r.FormValue("name")
		number := r.FormValue("number")
		phone := r.FormValue("phone")
		email := r.FormValue("email")

		_, err := db.Exec("INSERT INTO volunteers (name, number, phone, email) VALUES (?, ?, ?, ?)", name, number, phone, email)
		if err != nil {
			http.Error(w, "Failed to insert data", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Received: %s, %s, %s, %s", name, number, phone, email)
	})

	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
