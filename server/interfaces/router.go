package interfaces

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("watcher/:uuid/obtain", getWatcher)
	r.GET("watcher/:uuid/routines", getRoutines)
}
