package main

import (
	"dousheng/db"
	"dousheng/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	db.Init()
	router := routers.InitRouter()
	router.Run(":8080")
}
