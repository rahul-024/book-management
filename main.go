package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rahul-024/BookManagement/controllers"
	"github.com/rahul-024/BookManagement/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.Run()
}
