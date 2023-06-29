package sonolusgo

import (
	"github.com/gin-gonic/gin"
)

func SonolusVersionHandler(ctx *gin.Context) {
	sonolusVersion := "0.7.0"
	ctx.Header("Sonolus-Version", sonolusVersion)
}
