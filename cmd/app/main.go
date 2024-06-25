package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/AkbarFikri/Aironment-BE/internal/app/handler"
	"github.com/AkbarFikri/Aironment-BE/internal/app/handler/route"
	"github.com/AkbarFikri/Aironment-BE/internal/app/repository"
	"github.com/AkbarFikri/Aironment-BE/internal/app/service"
	"github.com/AkbarFikri/Aironment-BE/internal/pkg/database/postgres"
	"github.com/AkbarFikri/Aironment-BE/internal/pkg/supabase"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("error loading env file!")
	}

	db := postgres.NewPostgres()
	app := gin.New()
	supabase := supabase.NewSupabaseClient()

	// Repository
	UserRepository := repository.NewUser(db)
	CommunityMemberRepository := repository.NewCommunityMember(db)
	CommunityRepository := repository.NewCommunity(db)
	InvoiceRepository := repository.NewInvoice(db)
	PostRepository := repository.NewPost(db)

	// Service
	AuthService := service.NewAuth(UserRepository)
	UserService := service.NewUser(UserRepository)
	PaymentService := service.NewPayment(InvoiceRepository, CommunityRepository, supabase)
	CommunityService := service.NewCommunity(CommunityRepository, CommunityMemberRepository, UserRepository, PostRepository)
	AqiService := service.NewAqi()

	// Handler
	AuthHandler := handler.NewAuth(AuthService)
	UserHandler := handler.NewUser(UserService)
	CommunityHandler := handler.NewCommunity(CommunityService, PaymentService)
	PaymentHandler := handler.NewPayment(CommunityService, PaymentService)
	AqiHandler := handler.NewAqi(AqiService)

	route := route.RouteConfig{
		App:              app,
		AuthHandler:      AuthHandler,
		UserHandler:      UserHandler,
		CommunityHandler: CommunityHandler,
		PaymentHandler:   PaymentHandler,
		AqiHandler:       AqiHandler,
	}

	route.Setup()
	app.Run()
}
