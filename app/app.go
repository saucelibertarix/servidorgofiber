package app

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	database "github.com/saucelibertarix/servidorgofiber/database"
	"github.com/saucelibertarix/servidorgofiber/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var httpServer *fiber.App

type App struct {
}

func (a *App) Initialize(port string) {
	fmt.Println(utils.GetEnvVariable("MYSQL_USER"))
	fmt.Println(utils.GetEnvVariable("MYSQL_PASSWORD"))
	fmt.Println(utils.GetEnvVariable("MYSQL_DATABASE"))

	/*
		InitializeDatabase(
			utils.GetEnvVariable("MYSQL_USER"),
			utils.GetEnvVariable("MYSQL_PASSWORD"),
			utils.GetEnvVariable("MYSQL_DATABASE"))

		database.Migrate()
		//fakedatabase.CreateFakeData()

		InitializeHttpServer(port)
	*/
}

func HandleRoutes(api fiber.Router) {

}

func InitializeHttpServer(port string) {

}

func InitializeDatabase(user, password, databaseName string) {
	var err error
	connectionString := fmt.Sprintf(
		"%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		databaseName,
	)

	database.GormDB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal(err)
	}
}
