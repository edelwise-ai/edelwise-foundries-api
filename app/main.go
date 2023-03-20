package main

import (
	"Foundries/domain"
	_userHTTPDelivery "Foundries/users/controllers"
	_userRepo "Foundries/users/repository"
	_userUsecase "Foundries/users/usecase"
	"github.com/gin-gonic/gin"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgrespw dbname=postgres port=32768 sslmode=disable TimeZone=Asia/Singapore"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	log.Println("Database connected")
	log.Println("Running migrations...")
	// Migrate the schema
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		panic(err)
	}
	log.Println("Migrations ran successfully")

	r := gin.Default()

	UserRepo := _userRepo.NewUserRepository(db)
	UserUsecase := _userUsecase.NewUserUsecase(UserRepo)
	_userHTTPDelivery.NewUserHandler(r, UserUsecase)

	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

}
