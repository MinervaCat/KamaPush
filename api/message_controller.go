package api

import (
	pb "KamaPush/pb"
	"KamaPush/pkg/constants"
	"KamaPush/pkg/zlog"
	"github.com/gin-gonic/gin"
	"net/http"
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
	rsp, err := grpcClient.GetMessageBySeq(ctx, &req)
	if err != nil {
		zlog.Error(err.Error())
		JsonBack(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": rsp,
	})
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
