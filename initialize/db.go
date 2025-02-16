package initialize

import (
	"fmt"
	"log"
	"os"

	"github.com/raflinoob132/go-notes/initialize/dbmodel"
	"github.com/raflinoob132/go-notes/models"

	// Ganti dari mysql ke postgres
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var noteModel = &models.Note{}
var usersModel = &models.User{}
var favoriteModel = &models.Favorite{}

func ConnectDB(config *dbmodel.Config) {
	var err error
	// Format DSN untuk PostgreSQL di Railway
	// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	config.DBHost, config.DBPort, config.DBUserName, config.DBUserPassword, config.DBName)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Gunakan driver postgres
	if err != nil {
		log.Fatal("Gagal hubungkan ke database: " + err.Error())
	}

	log.Println("Migrasi...")
	if err := DB.AutoMigrate(noteModel, usersModel, favoriteModel); err != nil {
		log.Fatal("Gagal migrasi tabel: " + err.Error())
	}

	log.Println("Berhasil migrasi")
}

func LoadConfig() (config dbmodel.Config, err error) {
	config.DBHost = os.Getenv("PGHOST")             // Ganti MYSQLHOST -> PGHOST
	config.DBUserName = os.Getenv("PGUSER")         // Ganti MYSQLUSER -> PGUSER
	config.DBUserPassword = os.Getenv("PGPASSWORD") // Ganti MYSQL_PASSWORD -> PGPASSWORD
	config.DBName = os.Getenv("PGDATABASE")         // Ganti MYSQLNAME -> PGDATABASE
	config.DBPort = os.Getenv("PGPORT")             // Ganti MYSQLPORT -> PGPORT
	return
}
