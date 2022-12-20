package main

import (
	"os"

	"github.com/bangsyir/go-jwt/controllers"
	"github.com/bangsyir/go-jwt/initializers"
	"github.com/bangsyir/go-jwt/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte(os.Getenv("SECRET")))
	r.Use(sessions.Sessions(("mysession"), store))

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("logout", middleware.RequireAuth, controllers.Logout)

	r.Run()
}
