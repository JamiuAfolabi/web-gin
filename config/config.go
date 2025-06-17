package config

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Dbname   string
}

type AppConfig struct {
	DSN          string
	DBSchemaFile string
}
