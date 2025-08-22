package routes

import (
	"github.com/gin-gonic/gin"

	"ambridge-backend/controllers"
	"ambridge-backend/middleware"
)

// SetupCrewRoutes configures the crew routes
func SetupCrewRoutes(router *gin.Engine) {
	crew := router.Group("/crews")
	{
		// Public routes - anyone can view crew members
		crew.GET("", controllers.GetAllCrewMembers)
		crew.GET("/:id", controllers.GetCrewMember)

		// Protected routes (require authentication)
		// Only admins can create, update, or delete crew members
		// The admin check is done in the controller
		authRequired := crew.Group("/")
		authRequired.Use(middleware.AuthMiddleware())
		{
			authRequired.POST("", controllers.CreateCrew)
			authRequired.PUT("/:id", controllers.UpdateCrewMember)
			authRequired.DELETE("/:id", controllers.DeleteCrewMember)
		}
	}
}
