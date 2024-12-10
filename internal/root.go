package internal

import (
	"github.com/gin-gonic/gin"
	"qwen/internal/handle"
)

func InitRooter(rooter *gin.Engine) {
	assistantRoter := rooter.Group("/assistant")
	assistantRoter.POST("", handle.CreateAssistant)
	assistantRoter.PUT("", handle.UpdateAssistant)
	assistantRoter.DELETE("", handle.DeleteAssistant)
	assistantRoter.GET("", handle.GetAssistant)
	assistantRoter.GET("list", handle.GetAssistantList)

	assistantRoter.Group("/message")
}
