package dbmodel

//import "os"

type Config struct {
	DBHost         string `mapstructure:"PGHOST"`     // Ganti MYSQLHOST -> PGHOST
	DBUserName     string `mapstructure:"PGUSER"`     // Ganti MYSQLUSER -> PGUSER
	DBUserPassword string `mapstructure:"PGPASSWORD"` // Ganti MYSQL_PASSWORD -> PGPASSWORD
	DBName         string `mapstructure:"PGDATABASE"` // Ganti MYSQL_DATABASE -> PGDATABASE
	DBPort         string `mapstructure:"PGPORT"`     // Ganti MYSQL_PORT -> PGPORT

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}
