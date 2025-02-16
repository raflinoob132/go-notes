package initialize

import (
	"fmt"
	"log"
	"os"

	"github.com/raflinoob132/go-notes/initialize/dbmodel"
	"github.com/raflinoob132/go-notes/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var noteModel = &models.Note{}
var usersModel = &models.User{}
var favoriteModel = &models.Favorite{}

func ConnectDB(config *dbmodel.Config) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal hubungkan ke database" + err.Error())
		//os.Exit(1)
	}
	log.Println("Migrasi...")
	if err := DB.AutoMigrate(noteModel, usersModel, favoriteModel); err != nil {
		log.Fatal("Gagal migrasi tabel" + err.Error())
	}

	log.Println("Berhasil migrasi")
}
func LoadConfig() (config dbmodel.Config, err error) {
	config.DBHost = os.Getenv("MYSQLHOST")
	config.DBUserName = os.Getenv("MYSQLUSER")
	config.DBUserPassword = os.Getenv("MYSQL_PASSWORD")
	config.DBName = os.Getenv("MYSQLNAME")
	config.DBPort = os.Getenv("MYSQLPORT")
	return
}
