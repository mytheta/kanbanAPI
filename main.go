package main

import (
	"github.com/mytheta/kanbanAPI/router"
)

func main() {

	r := router.GetRouter()

	r.Run(":9000" )
}