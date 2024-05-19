package main

import (
	"github.com/gin-gonic/gin"

	"github.com/isaqueveras/outis-dashboard/server/interfaces/watcher"
)

func main() {
	r := gin.Default()

	v1 := r.Group("v1")
	watcher.Router(v1.Group("watcher/:wid"))

	r.Run(":8181")
}
