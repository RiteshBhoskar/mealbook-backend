package routes

import (
	"github.com/RiteshBhoskar/mealbook-backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(r *gin.Engine) {
	// API v1 routes
	api := r.Group("/api/v1")
	{
		// User routes
		users := api.Group("/users")
		{
			users.GET("", handlers.ListUsers)
			users.POST("", handlers.CreateUser)
			users.GET("/:id", handlers.GetUser)
			users.PUT("/:id", handlers.UpdateUser)
			users.DELETE("/:id", handlers.DeleteUser)
		}
	}
}

