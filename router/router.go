package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {


	r := gin.New()

	r.Use(gin.Logger())

	r.GET("/alive", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})



	api := r.Group("")
	apiRouter(api)

	return r

}
