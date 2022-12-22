package main

import (
	"fmt"
	"net"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/alramdein/user-service/config"
	httpDelivery "github.com/alramdein/user-service/delivery/http"
	"github.com/alramdein/user-service/model"
	"github.com/alramdein/user-service/pb"
	"github.com/alramdein/user-service/repository"
	"github.com/alramdein/user-service/usecase"
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

	// run http server in goroutine
	go func() {
		e := echo.New()

		httpDelivery.NewUserHandler(e, userUsecase)

		e.Use(middleware.Logger())
		e.Start(getHTTPServerAddress())
	}()

	// grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPCPort()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &model.UserServiceServer{})

	log.Info("grpc listen from ", config.GRPCPort())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return
	}
}

func getDatabaseDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost(), config.DBUsername(), config.DBPassword(), config.DBName(), config.DBPort(),
	)
}

func getHTTPServerAddress() string {
	return fmt.Sprintf("%v:%v", config.HTTPHost(), config.HTTPPort())
}
