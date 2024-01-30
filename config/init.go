package config

func InitializeAll() {
	loadEnvVariables()
	initializeConfigParameters()
	initializeSQLiteDatabase()
	runSQLiteMigration()
}
