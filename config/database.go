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
		initializeSQLiteDatabase()
	}
	return sqliteDB
}

func initializeSQLiteDatabase() {
	currentDir := util.GetWorkingDir()

	folderName := GetParams().DbFolderName
	fileName := GetParams().DbFileName

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

// Runs the migration script(s), panics if it fails
func runSQLiteMigration() {
	mustRunSQLiteScript("schema.sql")
}

// Runs a SQL script in sql/ folder
func runSQLiteScript(scriptName string) error {
	currentDir := util.GetWorkingDir()

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
func mustRunSQLiteScript(scriptName string) {
	util.PanicIfError(runSQLiteScript(scriptName))
}
