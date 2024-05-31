package report

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("/routine/:rid/get_histograms", getHistograms)
}
