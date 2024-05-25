package routine

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("obtain", obtain)
	r.GET("indicators", indicators)
	r.GET("logs", getLogs)
}
