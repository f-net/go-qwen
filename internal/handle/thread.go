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
	var req = &types.ListAssistantThreadReq{}
	err := c.ShouldBind(req)
	if err != nil {
		ErrorResp(c, err)
		return
	}

	//req.InitPageSize()

	resp, err := logic.NewAssistantThreadLogic(config.GetDB()).List(context.Background(), req)
	if err != nil {
		ErrorResp(c, err)
		return
	}
	//SuccessResp(c, resp)
	data, _ := json.Marshal(resp)
	c.HTML(200, "assistant/chat/list.html", map[string]interface{}{
		"title": "助手会话",
		"data":  string(data),
	})
}
