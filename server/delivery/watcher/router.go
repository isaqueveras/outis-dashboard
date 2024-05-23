package watcher

import (
	"github.com/gin-gonic/gin"

	"github.com/isaqueveras/outis-dashboard/server/delivery/watcher/routine"
)

func Router(r *gin.RouterGroup) {
	r.GET("obtain", getWatcher)
	r.GET("routines", getRoutines)

	routine.Router(r.Group("routine/:rid"))
}
