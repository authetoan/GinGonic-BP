package routers

import (
	"GoBP/controllers"
	"GoBP/middleware"
	"github.com/gin-gonic/gin"
)
func ApiRoutes(route *gin.Engine) {

	jwtMiddleware := middleware.JWTMiddleware()
	api := route.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", jwtMiddleware.LoginHandler)
		}
		api.GET("/ping-with-auth", jwtMiddleware.MiddlewareFunc(),controllers.PingWithAuth)
		api.GET("/ping", controllers.Ping)
		api.GET("/user", controllers.Example)

	}
}
