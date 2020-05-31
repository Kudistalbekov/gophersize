package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	user     = "kudi"
	password = "164137"
	dbname   = "phonenormalizer"
)

func main() {
	psql := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)
	db, err := sqlx.Open("postgres", psql)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}

func normalize(given string) string {
	var exp bytes.Buffer
	for _, val := range given {
		if val >= '0' && val <= '9' {
			exp.WriteRune(val)
		}
	}
	return exp.String()
}
