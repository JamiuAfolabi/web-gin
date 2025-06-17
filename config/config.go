package config

type DatabaseConfig struct {
	Host         string
	Port         int
	Username     string
	Password     string
	Dbname       string
	DSN          string
	DBSchemaFile string
}

type AppConfig struct {
}
