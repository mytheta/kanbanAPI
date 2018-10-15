package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mytheta/kanbanAPI/service"
	"net/http"
)

var Todo = todoimpl{}

type todoimpl struct {}

// @Summary Show a todo by staus
// @Description get string by status
// @Success 200 {object} model.Todo
// @Router /todo/{status} [get]
func(t *todoimpl) FindByStatus(c *gin.Context) {
	status := c.Param("status")

	todos, err := service.Todo.FindByStatus(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, todos)
}