package internal

import (
	"github.com/gin-gonic/gin"
	"qwen/internal/handle"
)

func InitRooter(rooter *gin.Engine) {
	rooter.GET("", handle.GetAssistantList)
	assistantRoter := rooter.Group("/assistant")
	assistantRoter.POST("", handle.CreateAssistant)

	assistantRoter.POST("update", handle.UpdateAssistant)
	assistantRoter.POST("delete", handle.DeleteAssistant)
	assistantRoter.GET("", handle.GetAssistant)
	assistantRoter.GET("list", handle.GetAssistantList)

	threadRooter := assistantRoter.Group("/thread")
	threadRooter.GET("list", handle.GetAssistantThreadList)

	messageRooter := threadRooter.Group("/message")
	messageRooter.GET("list", handle.GetMessageOnAssistantList)
	messageRooter.POST("send", handle.SendMessageOnAssistant)

}
