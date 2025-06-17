package config

// This includes the database configuration
type DatabaseConfig struct {
	Host         string
	Port         int
	Username     string
	Password     string
	Dbname       string
	DSN          string
	DBSchemaFile string
}

// This includes the app related configuration
type AppConfig struct {
}
