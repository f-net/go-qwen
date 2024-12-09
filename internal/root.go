package internal

import (
	"github.com/gin-gonic/gin"
	"qwen/internal/handle"
)

func InitRooter(rooter *gin.Engine) {
	rooter.Group("/assistant")
	rooter.POST("", handle.CreateAssistant)
}
