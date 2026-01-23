package api

import (
	pb "KamaPush/pb"
	"KamaPush/pkg/constants"
	"KamaPush/pkg/zlog"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// GetMessageList 获取聊天记录
//func GetMessageList(c *gin.Context) {
//	var req request.GetMessageListRequest
//	if err := c.BindJSON(&req); err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code":    500,
//			"message": constants.SYSTEM_ERROR,
//		})
//		return
//	}
//	//todo
//	message, rsp, ret := gorm.MessageService.GetMessageList(req.UserOneId, req.UserTwoId)
//	JsonBack(c, message, ret, rsp)
//}

func GetMessageBySeq(c *gin.Context) {
	var req pb.GetMessageBySeqRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	grpcRsp, err := grpcClient.GetMessageBySeq(ctx, &req)
	if err != nil {
		zlog.Error(err.Error())
		JsonBack(c, err)
		return
	}
	var rsp []MessageRespond
	for _, msg := range grpcRsp.MessageList {
		message := MessageRespond{
			MsgId:          msg.MsgId,
			ConversationId: msg.ConversationId,
			Seq:            msg.Seq,
			SendId:         msg.SendId,
			Type:           int8(msg.Type),
			Content:        msg.Content,
			Status:         int8(msg.Status),
			SendTime:       msg.SendTime.AsTime(),
		}
		rsp = append(rsp, message)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": rsp,
	})
}

type MessageRespond struct {
	MsgId          int64     `json:"msg_id"`
	ConversationId string    `json:"conversation_id"`
	Seq            int64     `json:"seq"`
	SendId         int64     `json:"send_id"`
	Type           int8      `json:"type"`
	Content        string    `json:"content"`
	Status         int8      `json:"status"`
	SendTime       time.Time `json:"send_time"`
}

//
//// UploadAvatar 上传头像
//func UploadAvatar(c *gin.Context) {
//	message, ret := gorm.MessageService.UploadAvatar(c)
//	JsonBack(c, message, ret, nil)
//}
//
//// UploadFile 上传头像
//func UploadFile(c *gin.Context) {
//	message, ret := gorm.MessageService.UploadFile(c)
//	JsonBack(c, message, ret, nil)
//}
