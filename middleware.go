package sonolusgo

import (
	"github.com/gin-gonic/gin"
)

var sonolusVersion = "0.8.4"

func SonolusVersionHandler(ctx *gin.Context) {
	ctx.Header("Sonolus-Version", sonolusVersion)
}
