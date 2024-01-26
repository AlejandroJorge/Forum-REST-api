package config

import (
	"database/sql"
	"fmt"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

var sqliteDB *sql.DB

func SQLiteDatabase() *sql.DB {
	if sqliteDB == nil {
		err := os.Mkdir(os.Getenv("SQLITE_DB_FOLDER_NAME"), 0755)
		if err != nil && !os.IsExist(err) {
			fmt.Println(err)
			panic(err)
		}
		dbPath := path.Join(os.Getenv("SQLITE_DB_FOLDER_NAME"), os.Getenv("SQLITE_DB_FILE_NAME"))
		newDB, err := sql.Open("sqlite3", dbPath)
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

func RunSQLiteMigration() {
	migrationFilePath := path.Join("sql", "schema.sql")

	migrationScriptBytes, err := os.ReadFile(migrationFilePath)
	if err != nil {
		panic(err)
	}

	migrationScript := string(migrationScriptBytes)

	db := SQLiteDatabase()

	_, err = db.Exec(migrationScript)
	if err != nil {
		panic(err)
	}

}
