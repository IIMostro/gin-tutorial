package repository

import (
	"database/sql"
	"ilmostro.org/gin-tutorial/configuration"
)

var Connection *sql.DB

type Model struct {
	Id int `json:"id"`

	CreatedOn int `json:"create_on"`
}

func Setup() {
	Connection, _ = configuration.GetConnection()
}
