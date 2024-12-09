package main

import (
	"github.com/gin-gonic/gin"
	"qwen/internal/model"
)

func main() {
	model.InitMysql()

	gd := gin.Default()
	gd.Run(":8000")
}
