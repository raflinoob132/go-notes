package dbmodel

//import "os"

type Config struct {
	DBHost         string `mapstructure:"MYSQL_HOST"`
	DBUserName     string `mapstructure:"MYSQL_USER"`
	DBUserPassword string `mapstructure:"MYSQL_PASSWORD"`
	DBName         string `mapstructure:"MYSQL_DATABASE"`
	DBPort         string `mapstructure:"MYSQL_PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}
