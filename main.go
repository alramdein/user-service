package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/alramdein/config"
	httpDelivery "github.com/alramdein/delivery/http"
	"github.com/alramdein/repository"
	"github.com/alramdein/usecase"
)

func main() {
	config.GetConf()
	db, err := gorm.Open(postgres.Open(getDatabaseDSN()), &gorm.Config{})
	if err != nil {
		log.Error(err.Error())
		panic("failed to connect database")
	}

	userRepo := repository.NewUserRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepo)

	e := echo.New()

	httpDelivery.NewUserHandler(e, userUsecase)

	e.Use(middleware.Logger())

	e.Start(getHTTPServerAddress())
}

func getDatabaseDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost(), config.DBUsername(), config.DBPassword(), config.DBName(), config.DBPort(),
	)
}

func getHTTPServerAddress() string {
	return fmt.Sprintf("%v:%v", config.HTTPHost(), config.HTTPPort())
}
