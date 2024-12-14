package handle

import (
	"context"
	"github.com/gin-gonic/gin"
	"qwen/internal/config"
	"qwen/internal/logic"
	"qwen/internal/types"
)

func SendMessageOnAssistant(c *gin.Context) {
	var req = &types.CreateMessageReq{}
	var resp = &types.CreateMessageResp{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		ErrorResp(c, err)
	}
	tx := config.GetDB()
	tx = tx.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
			SuccessResp(c, resp)
			return
		}
		tx.Rollback()
		ErrorResp(c, err)
	}()
	resp.Answer, resp.ThreadId, err = logic.NewMessageLogic(tx).SendMessage(context.Background(), req)
}

func GetMessageOnAssistantList(c *gin.Context) {
	var req = &types.GetMessageListReq{}
	err := c.ShouldBind(req)
	if err != nil {
		ErrorResp(c, err)
	}
	list, err := logic.NewMessageLogic(config.GetDB()).List(context.Background(), req)
	if err != nil {
		ErrorResp(c, err)
	}
	SuccessResp(c, list)
}
