package repository

import (
	"database/sql"
	"ilmostro.org/gin-tutorial/configuration"
)

var Connection *sql.DB

type Model struct {
	Id int

	CreatedOn int
}

func Setup() {
	Connection = configuration.GetConnection()
}
