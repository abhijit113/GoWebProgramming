package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Age       int    `json:"age"`
}

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:Root@123@tcp(localhost:3306)/sakila")
	if err != nil {
		panic(err)
	}
}

func retrieve(w http.ResponseWriter, r *http.Request) {

	var ps Person
	// Execute the query
	id := 3
	err := Db.QueryRow("SELECT * from persons where id = ?", id).Scan(&ps.ID, &ps.FirstName, &ps.LastName, &ps.Age)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	log.Println(ps.ID)
	log.Println(ps.Age)
	log.Println(ps.FirstName)
	log.Println(ps.LastName)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	jsonValue, _ := json.Marshal(ps)
	w.Write(jsonValue)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/get", http.HandlerFunc(retrieve))
	server.ListenAndServe()

	defer Db.Close()

}
