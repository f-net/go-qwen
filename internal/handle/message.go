package handle

import (
	"github.com/gin-gonic/gin"
	"qwen/internal/types"
)

func SendMessageOnAssistant(c *gin.Context) {
	var req = &types.Message{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		ErrorResp(c, err)
	}
	//logic.NewAssistantLogic(config.GetDB()).Create()
}
