package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "db_name"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world, I'm running on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)
}

type MemberData struct {
	ID     int64
	Handle string
	Email  string
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	rawPath := r.URL.Path
	splittedString := strings.Split(rawPath, "/")
	username := splittedString[2]
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	// fmt.Println(db)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	queryString := "SELECT id, handle, email FROM member WHERE handle = $1"
	iterator, err := db.Query(queryString, username)
	defer iterator.Close()
	var rows []MemberData
	for iterator.Next() {
		var row = MemberData{}
		err = iterator.Scan(
			&row.ID, &row.Handle, &row.Email)
		if err != nil {
			err = errors.Wrapf(err, "Event row scanning failed (type=%s)", username)
			log.Fatal(err)
		}

		rows = append(rows, row)
	}
	fmt.Fprintf(w, "Hello %s. Your member id is %d and email is: %s", rows[0].Handle, rows[0].ID, rows[0].Email)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/user/", userHandler)
	fmt.Println("Go Boilerplate Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
