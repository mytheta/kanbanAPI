package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mytheta/kanbanAPI/controller"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "./docs"

)


// @title Swagger Example API
// @version 1.0
// @description This is a kanbanAPI



// @host localhost:9000
// @BasePath /
func main() {

	r := gin.New()

	r.Use(gin.Logger())

	// use ginSwagger middleware to
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/todo/:status", controller.Todo.FindByStatus)


	r.Run(":9000" )
}