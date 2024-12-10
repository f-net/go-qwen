package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorResp(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": 500,
		"msg":  err.Error(),
		"data": nil,
	})
}

func SuccessResp(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}
