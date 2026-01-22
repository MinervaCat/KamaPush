package api

import (
	pb "KamaPush/pb"
	"KamaPush/pkg/constants"
	"KamaPush/pkg/zlog"

	"github.com/gin-gonic/gin"

	"net/http"
)

func GetConversationList(c *gin.Context) {
	var req pb.UserIdRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	conversationList, err := grpcClient.GetConversationList(ctx, &req)
	if err != nil {
		zlog.Error(err.Error())
		JsonBack(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": conversationList,
	})
}

// DeleteSession 删除会话
//func DeleteConversation(c *gin.Context) {
//	var req request.UserConversationRequest
//	if err := c.BindJSON(&req); err != nil {
//		zlog.Error(err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"code":    500,
//			"message": constants.SYSTEM_ERROR,
//		})
//		return
//	}
//	message, ret := gorm.ConversationService.DeleteConversation(req.UserId, req.ConversationId)
//	JsonBack(c, message, ret, nil)
//}
//
//// CheckOpenSessionAllowed 检查是否可以打开会话
//func CheckOpenConversationAllowed(c *gin.Context) {
//	var req request.UserFriendRequest
//	if err := c.BindJSON(&req); err != nil {
//		zlog.Error(err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"code":    500,
//			"message": constants.SYSTEM_ERROR,
//		})
//		return
//	}
//	message, res, ret := gorm.ConversationService.CheckOpenConversationAllowed(req.UserId, req.FriendId)
//	JsonBack(c, message, ret, res)
//}
