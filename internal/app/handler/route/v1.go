package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/handler"
	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/handler/middleware"
)

type RouteConfig struct {
	App         *gin.Engine
	AuthHandler handler.AuthHandler
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
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Backend ready!"})
	})

	v1 := c.App.Group("/api/v1")
	c.AuthRoute(v1)
}

func (c *RouteConfig) AuthRoute(r *gin.RouterGroup) {
	authEnds := r.Group("/auth")
	authEnds.POST("/register", c.AuthHandler.Register)
	authEnds.POST("/login", c.AuthHandler.Login)
}