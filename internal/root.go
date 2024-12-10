package internal

import (
	"github.com/gin-gonic/gin"
	"qwen/internal/handle"
)

func InitRooter(rooter *gin.Engine) {
	assistantRoter := rooter.Group("/assistant")
	assistantRoter.POST("", handle.CreateAssistant)
}
