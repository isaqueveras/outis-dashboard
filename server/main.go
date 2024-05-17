package main

import (
	"github.com/gin-gonic/gin"

	"github.com/isaqueveras/outis-dashboard/server/interfaces"
)

func main() {
	r := gin.Default()
	interfaces.Router(r.Group("v1"))
	r.Run(":8181")
}
