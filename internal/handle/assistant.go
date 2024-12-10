package handle

import (
	"context"
	"github.com/gin-gonic/gin"
	"qwen/internal/config"
	"qwen/internal/logic"
	"qwen/internal/types"
)

func CreateAssistant(c *gin.Context) {
	var req = &types.CreateAssistantReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		ErrorResp(c, err)
		return
	}
	err = logic.NewAssistantLogic(config.GetDB()).Create(context.Background(), req)
	if err != nil {
		ErrorResp(c, err)
		return
	}
	SuccessResp(c, "")

}

func UpdateAssistant(c *gin.Context) {
	var req = &types.UpdateAssistantReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		ErrorResp(c, err)
		return
	}
	err = logic.NewAssistantLogic(config.GetDB()).Save(context.Background(), req)
	if err != nil {
		ErrorResp(c, err)
		return
	}
	SuccessResp(c, "")
}

func GetAssistant(c *gin.Context) {
	var req = &types.GetId{}
	err := c.ShouldBind(req)
	if err != nil {
		ErrorResp(c, err)
		return
	}
	resp, err := logic.NewAssistantLogic(config.GetDB()).First(context.Background(), req.Id)
	if err != nil {
		ErrorResp(c, err)
		return
	}
	SuccessResp(c, resp)
}
func GetAssistantList(c *gin.Context) {
	var req = &types.ListAssistantReq{}
	err := c.ShouldBind(req)
	if err != nil {
		ErrorResp(c, err)
		return
	}

	req.InitPageSize()

	resp, err := logic.NewAssistantLogic(config.GetDB()).List(context.Background(), req)
	if err != nil {
		ErrorResp(c, err)
		return
	}
	SuccessResp(c, resp)
}
