package routes

import (
	"github.com/gin-gonic/gin"

	"ambridge-backend/controllers"
	"ambridge-backend/middleware"
)

// SetupProjectRoutes configures the project routes
func SetupProjectRoutes(router *gin.Engine) {
	project := router.Group("/projects")
	{
		// Public routes
		project.GET("", controllers.GetAllProjects)
		project.GET("/:id", controllers.GetProject)

		// Protected routes (require authentication)
		authRequired := project.Group("/")
		authRequired.Use(middleware.AuthMiddleware())
		{
			authRequired.POST("", controllers.CreateProject)
			authRequired.PUT("/:id", controllers.UpdateProject)
			authRequired.DELETE("/:id", controllers.DeleteProject)
		}
	}
}
