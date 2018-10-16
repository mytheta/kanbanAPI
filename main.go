package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mytheta/kanbanAPI/controller"

)


func main() {

	r := gin.New()

	r.Use(gin.Logger())

	r.GET("/todo/:status", controller.Todo.FindByStatus)


	r.Run(":9000" )
}