package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mytheta/kanbanAPI/controller"
)

func apiRouter(api *gin.RouterGroup) {
	api.GET("/todo/:status", controller.Todo.FindByStatus)

}

