package main

import (
	sonolus "github.com/6qhtsk/sonolusgo"
	"github.com/gin-gonic/gin"
)

func main() {
	sonolusConfig := sonolus.DefaultConfig()
	router := gin.Default()
	sonolusConfig.LoadHandlers(router)
	err := router.Run()
	if err != nil {
		panic(err)
	}
	return
}
