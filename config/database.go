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
		currentDir := GetWorkingDir()

		folderName, ok := os.LookupEnv("SQLITE_DB_FOLDER_NAME")
		if !ok {
			folderName = "data"
		}
		fileName, ok := os.LookupEnv("SQLITE_DB_FILE_NAME")
		if !ok {
			fileName = "database.sqlite"
		}

		folderPath := path.Join(currentDir, folderName)
		err := os.Mkdir(folderPath, 0755)
		if err != nil && !os.IsExist(err) {
			fmt.Println(err)
			panic(err)
		}

		dbPath := path.Join(folderPath, fileName)

		connectionStr := "file:" + dbPath + "?_journal=WAL&_foreign_keys=true"
		newDB, err := sql.Open("sqlite3", connectionStr)
		util.PanicIfError(err)

		sqliteDB = newDB
	}
	return sqliteDB
}

// Runs the migration script(s), panics if it fails
func RunSQLiteMigration() {
	MustRunSQLiteScript("schema.sql")
}

// Runs a SQL script in sql/ folder
func RunSQLiteScript(scriptName string) error {
	currentDir := GetWorkingDir()

	filePath := path.Join(currentDir, "sql", scriptName)
	scriptBytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	script := string(scriptBytes)

	_, err = SQLiteDatabase().Exec(script)
	if err != nil {
		return err
	}

	return nil
}

// Runs a SQL script in sql/ folder, panics when it fails
func MustRunSQLiteScript(scriptName string) {
	util.PanicIfError(RunSQLiteScript(scriptName))
}
