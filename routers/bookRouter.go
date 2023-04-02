package routers

import (
	"tugas8/controllers"
	"tugas8/repository"
	"github.com/gin-gonic/gin"
)

func SetupRouter(repo *repository.BookRepository) *gin.Engine {
	router := gin.Default()

	bookController := controllers.NewBookController(repo)

	router.GET("/books", bookController.FindAll)
	router.GET("/books/:id", bookController.FindByID)
	router.POST("/books", bookController.Create)
	router.PUT("/books/:id", bookController.Update)
	router.DELETE("/books/:id", bookController.Delete)

	return router
}