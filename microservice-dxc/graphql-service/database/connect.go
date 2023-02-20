package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "postgres-graphql" // this host name is from the docker file
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "users"
)

func ConnectToPostgres() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	return sql.Open("postgres", psqlInfo)

}
