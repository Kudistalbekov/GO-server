package conn

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

//DBconnect connects to database
func DBconnect() *sqlx.DB {
	db, err := sqlx.Open("postgres", "user=kudi dbname=test password = '164137' sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("db connected")
	return db
}
