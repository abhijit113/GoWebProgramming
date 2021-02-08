package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAll(w http.ResponseWriter, r *http.Request) {
	results, err := Db.Query("SELECT * FROM persons")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	psList := []Person{}
	for results.Next() {
		ps := Person{}
		err = results.Scan(&ps.ID, &ps.FirstName, &ps.LastName, &ps.Age)
		if err != nil {
			panic(err.Error())
		}
		log.Println(ps.ID)
		psList = append(psList, ps)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	jsonValue, _ := json.Marshal(psList)
	w.Write(jsonValue)
}

func returnSingle(w http.ResponseWriter, r *http.Request) {
	ps := Person{}
	vars := mux.Vars(r)
	key := vars["id"]
	err := Db.QueryRow("SELECT * from persons where id = ?", key).Scan(&ps.ID, &ps.FirstName, &ps.LastName, &ps.Age)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	jsonValue, _ := json.Marshal(ps)
	w.Write(jsonValue)
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/person/{id}", returnSingle)
	myRouter.HandleFunc("/persons", returnAll)
	log.Fatal(http.ListenAndServe(":8080", myRouter))

	defer Db.Close()
}
