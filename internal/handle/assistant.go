package handle

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"qwen/internal/config"
	"qwen/internal/logic"
	"qwen/internal/types"
)

func CreateAssistant(c *gin.Context) {
	var req = &types.Assistant{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = logic.NewAssistantLogic(config.GetDB()).Create(context.Background(), req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "")

}
