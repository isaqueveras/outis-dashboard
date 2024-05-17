package interfaces

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getRoutines(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": []map[string]string{
			{
				"uuid": "84cbfa32-cfa2-4f75-99d6-5423eca60985",
				"name": "Routine 01",
				"desc": "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
			},
			{
				"uuid": "73cb83e6-bcf2-47e4-b3e6-b672591e1736",
				"name": "Routine 02",
				"desc": "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
			},
			{
				"uuid": "634e4ef5-c99a-4e8d-ba3f-35386e42e8e2",
				"name": "Routine 03",
				"desc": "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
			},
		},
	})
}

func getWatcher(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"uuid":       ctx.Param("uuid"),
		"name":       "Minhas Principais Rotinas",
		"desc":       "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
		"created_at": time.Now(),
	})
}
