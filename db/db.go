package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/fridauxd/cht/db/sqlc"
)

var Queries *sqlc.Queries

func init() {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		panic(err)
	}

	Queries = sqlc.New(db)
}
