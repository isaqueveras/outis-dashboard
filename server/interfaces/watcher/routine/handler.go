package routine

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/isaqueveras/outis-dashboard/server/application/watcher/routine"
)

func obtain(ctx *gin.Context) {
	data, err := routine.Obtain(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func indicators(ctx *gin.Context) {
	data, err := routine.Indicators(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func getLogs(ctx *gin.Context) {
	data, err := routine.GetLogs(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}
