package config

import (
	"database/sql"
	"fmt"
	"os"
	"path"

	"github.com/AlejandroJorge/forum-rest-api/util"
	_ "github.com/mattn/go-sqlite3"
)

var sqliteDB *sql.DB

func SQLiteDatabase() *sql.DB {
	if sqliteDB == nil {
		currentDir, err := os.Getwd()
		util.PanicIfError(err)

		folderName := os.Getenv("SQLITE_DB_FOLDER_NAME")
		fileName := os.Getenv("SQLITE_DB_FILE_NAME")

		folderPath := path.Join(currentDir, folderName)
		err = os.Mkdir(folderPath, 0755)
		if err != nil && !os.IsExist(err) {
			fmt.Println(err)
			panic(err)
		}

		dbPath := path.Join(folderPath, fileName)
		newDB, err := sql.Open("sqlite3", dbPath)
		util.PanicIfError(err)

		sqliteDB = newDB
	}
	return sqliteDB
}

// Runs the migration script(s), panics if it fails
func RunSQLiteMigration() {
	MustRunSQLiteScript(SQLiteDatabase(), "schema.sql")
}

// Runs a SQL script in sql/ folder
func RunSQLiteScript(db *sql.DB, scriptName string) error {
	filePath := path.Join("sql", scriptName)
	scriptBytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	script := string(scriptBytes)

	_, err = db.Exec(script)
	if err != nil {
		return err
	}

	return nil
}

// Runs a SQL script in sql/ folder, panics when it fails
func MustRunSQLiteScript(db *sql.DB, scriptName string) {
	util.PanicIfError(RunSQLiteScript(db, scriptName))
}
