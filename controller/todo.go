package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mytheta/kanbanAPI/service"
	"net/http"
)

var Todo = todoimpl{}

type todoimpl struct {}

func(t *todoimpl) FindByStatus(c *gin.Context) {
	status := c.Param("status")

	todos, err := service.Todo.FindByStatus(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, todos)
}