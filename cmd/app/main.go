package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/handler"
	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/handler/route"
	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/repository"
	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/service"
	"github.com/AkbarFikri/hackfestuc2024_backend/internal/pkg/database/postgres"

)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("error loading env file!")
	}

	db := postgres.NewPostgres()
	app := gin.New()

	// Repository
	UserRepository := repository.NewUser(db)

	// Service
	AuthService := service.NewAuth(UserRepository)
	UserService := service.NewUser(UserRepository)

	// Handler
	AuthHandler := handler.NewAuth(AuthService)
	UserHandler := handler.NewUser(UserService)

	route := route.RouteConfig{
		App:         app,
		AuthHandler: AuthHandler,
		UserHandler: UserHandler,
	}

	route.Setup()
	app.Run()
}
