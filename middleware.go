package sonolusgo

import (
	"github.com/gin-gonic/gin"
)

var sonolusVersion = "0.8.5"

func SonolusVersionHandler(ctx *gin.Context) {
	ctx.Header("Sonolus-Version", sonolusVersion)
}
