package database

import (
	"log"
	"os"

	"github.com/toticavalcanti/fiber-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	var dsn = "root:mysql0401@/go_api?charset=utf8mb4&parseTime=True&loc=Local"
	var v = "Não conseguiu conectar ao banco de dados\n"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(v, err.Error())
		os.Exit(2)
	}

	log.Println("Conectado ao banco de dados!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Executando as migrações")
	// Add migrations
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{Db: db}
}
