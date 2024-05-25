package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/outis-dashboard/server/config"
	"github.com/isaqueveras/outis-dashboard/server/database"
	"github.com/isaqueveras/outis-dashboard/server/delivery"
)

func main() {
	r := gin.Default()
	config.Load()

	database.OpenConnections(&config.Get().Database)
	defer database.CloseConnections()

	delivery.Router(r.Group("v1"))
	r.Run(":8181")
}
