package metric

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/outis-dashboard/server/application/metric"
)

func event(ctx *gin.Context) {
	in := &metric.Metric{}
	if err := ctx.BindJSON(in); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	go metric.Event(ctx.Copy(), in)
	ctx.JSON(http.StatusNoContent, nil)
}

func setup(ctx *gin.Context) {
	in := &metric.SetupIn{}
	if err := ctx.BindJSON(in); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	go metric.Setup(ctx.Copy(), in)
	ctx.JSON(http.StatusNoContent, nil)
}
