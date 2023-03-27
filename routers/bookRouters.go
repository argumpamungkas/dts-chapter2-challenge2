package routers

import (
	"chapter2-challenge-sesi-2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", controllers.CreateBooks)
	router.GET("/books", controllers.GetAllBooks)
	router.GET("/books/:bookID", controllers.GetBookById)
	router.POST("/books/:bookID", controllers.UpdateBooks)
	router.DELETE("/books/:bookID", controllers.DeleteBook)

	return router
}
