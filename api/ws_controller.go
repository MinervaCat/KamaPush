package api

import (
	"github.com/gin-gonic/gin"

	"KamaPush/internal/service/push"
	"KamaPush/pkg/constants"
	"KamaPush/pkg/zlog"
	"net/http"
	"strconv"
)

// WsLogin wss登录 Get
func WsLogin(c *gin.Context) {
	//todo
	clientId := c.Query("user_id")
	if clientId == "" {
		zlog.Error("userId获取失败")
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "clientId获取失败",
		})
		return
	}
	userId, err := strconv.ParseInt(clientId, 10, 64)
	if err != nil {
		zlog.Error("userId转换失败")
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "clientId获取失败",
		})
		return
	}
	zlog.Info("userId获取成功")
	push.NewClientInit(c, userId)
}

// WsLogout wss登出
func WsLogout(c *gin.Context) {
	var req UserRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := push.ClientLogout(req.UserId)
	if ret == -1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": message,
		})
	}

}

type UserRequest struct {
	UserId int64 `json:"user_id"`
}
