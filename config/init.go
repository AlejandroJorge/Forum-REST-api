package config

func Initialize() {
	LoadEnvVariables()
	RunSQLiteMigration()
}
