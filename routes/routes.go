package routes

import (
	"sb-go-quiz/controllers"
	"sb-go-quiz/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/users/login", controllers.LoginUser)

		// Protected routes
		auth := api.Group("/")
		auth.Use(middlewares.JWTAuthMiddleware())
		{
			// Category endpoints
			auth.GET("/categories", controllers.GetAllCategories)
			auth.POST("/categories", controllers.CreateCategory)
			auth.GET("/categories/:id", controllers.GetCategoryByID)
			auth.DELETE("/categories/:id", controllers.DeleteCategory)
			auth.GET("/categories/:id/books", controllers.GetBooksByCategory)

			// Book endpoints
			auth.GET("/books", controllers.GetAllBooks)
			auth.POST("/books", controllers.CreateBook)
			auth.GET("/books/:id", controllers.GetBookByID)
			auth.DELETE("/books/:id", controllers.DeleteBook)
		}
	}

	return r
}
