package main

import (
	"github.com/gin-gonic/gin"
	"qwen/internal"
	"qwen/internal/config"
)

func main() {
	config.InitConfig()
	config.InitMysql()
	config.InitOpenaiClient()

	r := gin.Default()
	internal.InitRooter(r)

	r.Run(":8000")
}
