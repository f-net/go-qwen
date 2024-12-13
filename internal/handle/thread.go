package handle

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"qwen/internal/config"
	"qwen/internal/logic"
	"qwen/internal/types"
)

func GetAssistantThreadList(c *gin.Context) {
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
	//SuccessResp(c, resp)
	data, _ := json.Marshal(resp)
	c.HTML(200, "assistant/list.html", map[string]interface{}{
		"title": "助手列表",
		"data":  string(data),
	})
}
