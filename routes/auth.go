package routes

import (
	"github.com/gin-gonic/gin"

	"ambridge-backend/controllers"
	"ambridge-backend/middleware"
)

// SetupAuthRoutes configures the authentication routes
func SetupAuthRoutes(router *gin.Engine) {
	auth := router.Group("/api/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.POST("/refresh-token", controllers.RefreshToken)

		// Protected routes
		authRequired := auth.Group("/")
		authRequired.Use(middleware.AuthMiddleware())
		{
			authRequired.POST("/logout", controllers.Logout)
		}
	}
}
