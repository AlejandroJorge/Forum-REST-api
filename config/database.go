package config

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	sqliteDB     *sql.DB
	sqliteDBName string = "data.db"
)

func SQLiteDatabase() *sql.DB {
	if sqliteDB == nil {
		newDB, err := sql.Open("sqlite3", sqliteDBName)
		if err != nil {
			panic(err)
		}
		if err = newDB.Ping(); err != nil {
			panic(err)
		}
		sqliteDB = newDB
	}
	return sqliteDB
}
