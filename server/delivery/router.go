package delivery

import (
	"github.com/gin-gonic/gin"

	"github.com/isaqueveras/outis-dashboard/server/delivery/metric"
	"github.com/isaqueveras/outis-dashboard/server/delivery/report"
	"github.com/isaqueveras/outis-dashboard/server/delivery/watcher"
)

func Router(r *gin.RouterGroup) {
	metric.Router(r.Group("metric"))
	watcher.Router(r.Group("watcher/:wid"))
	report.Router(r.Group("report"))
}
