package database

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func GetDb() *bun.DB {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:mw.db?_journal_mode=WAL&_synchronous=NORMAL&_cache_size=10000&_foreign_keys=on")
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())
	return db
}

func GetTestDb() *bun.DB {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:mw.db?_journal_mode=WAL&_synchronous=NORMAL&_cache_size=10000&_foreign_keys=on")
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())
	return db
}
