package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/handler"
	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/handler/middleware"

)

type RouteConfig struct {
	App              *gin.Engine
	AuthHandler      handler.AuthHandler
	UserHandler      handler.UserHandler
	CommunityHandler handler.CommunityHandler
	PaymentHandler   handler.PaymentHandler
	AqiHandler       handler.AqiHandler
	PostHandler      handler.PostHandler
}

func (c *RouteConfig) Setup() {
	c.ServeRoute()
}

func (c *RouteConfig) ServeRoute() {
	c.App.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	c.App.Use(gin.Logger())
	c.App.Use(gin.Recovery())
	c.App.Use(middleware.CORSMiddleware())

	c.App.GET("/healt", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Backend ready! Testing 3"})
	})

	v1 := c.App.Group("/api/v1")
	c.AuthRoute(v1)
	c.UserRoute(v1)
	c.CommunityRoute(v1)
	c.PaymentRoute(v1)
	c.AQIRoute(v1)
}

func (c *RouteConfig) AuthRoute(r *gin.RouterGroup) {
	authEnds := r.Group("/auth")
	authEnds.POST("/register", c.AuthHandler.Register)
	authEnds.POST("/login", c.AuthHandler.Login)
}

func (c *RouteConfig) UserRoute(r *gin.RouterGroup) {
	userEnds := r.Group("/user")
	userEnds.GET("/current",middleware.JwtUser(), c.UserHandler.CurrentUser)
	userEnds.GET("/airqualitys", c.UserHandler.GetAirqualityPoints)
}

func (c *RouteConfig) CommunityRoute(r *gin.RouterGroup) {
	commEnds := r.Group("/community")
	commEnds.Use(middleware.JwtUser())
	commEnds.GET("", c.CommunityHandler.GetCommunities)
	commEnds.GET("/:id", c.CommunityHandler.GetCommunityDetails)
	commEnds.POST("", c.CommunityHandler.CreateCommunity)
	commEnds.POST("/join", c.CommunityHandler.JoinCommunity)
}

func (c *RouteConfig) PaymentRoute(r *gin.RouterGroup) {
	paymentEnds := r.Group("/payment")
	paymentEnds.POST("/verify", c.PaymentHandler.Verify)
}

func (c *RouteConfig) AQIRoute(r *gin.RouterGroup) {
	AqiEnds := r.Group("/airquality")
	AqiEnds.GET("", c.AqiHandler.GetCurrentPosition)
}

func (c *RouteConfig) PostRoute(r *gin.RouterGroup) {
	postEnds := r.Group("/post")
	postEnds.GET("/:id", c.PostHandler.GetData)
}
