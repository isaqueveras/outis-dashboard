package metric

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/outis-dashboard/server/application/metric"
)

func save(ctx *gin.Context) {
	metric.Save(ctx.Copy())
	ctx.JSON(http.StatusNoContent, nil)
}
