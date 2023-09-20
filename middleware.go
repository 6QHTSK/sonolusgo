package sonolusgo

import (
	"github.com/gin-gonic/gin"
)

var sonolusVersion = "0.7.3"

func OverWriteVersion(version string) {
	sonolusVersion = version
}

func SonolusVersionHandler(ctx *gin.Context) {
	ctx.Header("Sonolus-Version", sonolusVersion)
}
