package metric

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/outis-dashboard/server/application/metric"
)

func save(ctx *gin.Context) {
	in := &metric.Metric{}
	if err := ctx.BindJSON(in); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}

	go metric.Save(ctx.Copy(), in)

	ctx.JSON(http.StatusNoContent, nil)
}
