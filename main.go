package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! You are connected to DB")
}
func main() {
	connStr := "postgresql://XXXXX:XXXXX@10.5.0.5/XXXXX?sslmode=disable"

	_, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
