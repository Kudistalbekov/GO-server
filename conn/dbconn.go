package conn

import (
	"database/sql"
	"log"
	_"github.com/jmoiron/sqlx"
)

func DBconnect()*sqlx.DB {
	op:="DBconnect"
	db,err:=sqlx.Open("postgres","user=posgres dbname=test password=164137 sslmode=disable")
	err!=nil{
		log.Fatal(err+op)
	}
	err=db.Ping()
	if  err!={
		log.Fatal(err+op)
	}
	fmt.Println("db connected")
	return db
}
