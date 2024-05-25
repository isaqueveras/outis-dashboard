package metric

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("event", event)
	r.POST("setup", setup)
}
