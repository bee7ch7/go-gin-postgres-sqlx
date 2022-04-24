package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "172.30.243.242"
	port     = 5432
	user     = "root"
	password = "secret"
	dbname   = "transbank"
)

var DBConn *sqlx.DB

func InitDB() {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// db, err := sql.Open("postgres", dbinfo)
	db, err := sqlx.Open("postgres", dbinfo)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	DBConn = db

}
