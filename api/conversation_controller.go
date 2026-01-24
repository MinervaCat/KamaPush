package api

import (
	pb "KamaPush/pb"
	"KamaPush/pkg/constants"
	"KamaPush/pkg/zlog"
	"time"

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
	grpcRsp, err := grpcClient.GetConversationList(ctx, &req)
	if err != nil {
		zlog.Error(err.Error())
		JsonBack(c, err)
		return
	}
	var rsp []ConversationResponse
	for _, con := range grpcRsp.ConversationList {
		conversation := ConversationResponse{
			ConversationId: con.ConversationId,
			Avatar:         con.Avatar,
			Type:           int8(con.Type),
			Member:         con.Member,
			RecentMsgTime:  con.RecentMsgTime.AsTime(),
			LastReadSeq:    con.LastReadSeq,
			NotifyType:     int8(con.NotifyType),
			IsTop:          int8(con.IsTop),
			Name:           con.Name,
		}
		rsp = append(rsp, conversation)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": rsp,
	})
}

type ConversationResponse struct {
	ConversationId string    `json:"conversation_id"`
	Avatar         string    `json:"avatar"`
	Type           int8      `json:"type"`
	Member         int32     `json:"member"`
	RecentMsgTime  time.Time `json:"recent_msg_time"`
	LastReadSeq    int64     `json:"last_read_seq"`
	NotifyType     int8      ` json:"notify_type"`
	IsTop          int8      `json:"is_top"`
	Name           string    `json:"name"`
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
