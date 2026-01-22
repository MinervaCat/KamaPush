package api

import (
	pb "KamaPush/pb"
	"KamaPush/pkg/constants"
	"KamaPush/pkg/zlog"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserList 获取联系人列表
func GetFriendList(c *gin.Context) {
	var req pb.UserIdRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
	}
	rsp, err := grpcClient.GetFriendList(ctx, &req)
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

// LoadMyJoinedGroup 获取我加入的群聊
//func LoadMyJoinedGroup(c *gin.Context) {
//	var loadMyJoinedGroupReq request.OwnlistRequest
//	if err := c.BindJSON(&loadMyJoinedGroupReq); err != nil {
//		zlog.Error(err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"code":    500,
//			"message": constants.SYSTEM_ERROR,
//		})
//		return
//	}
//	message, groupList, ret := gorm.UserContactService.LoadMyJoinedGroup(loadMyJoinedGroupReq.OwnerId)
//	JsonBack(c, message, ret, groupList)
//}

// GetContactInfo 获取联系人信息
//func GetFriendInfo(c *gin.Context) {
//	var getFriendInfoReq request.UserRequest
//	if err := c.BindJSON(&getFriendInfoReq); err != nil {
//		zlog.Error(err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"code":    500,
//			"message": constants.SYSTEM_ERROR,
//		})
//		return
//	}
//	log.Println(getFriendInfoReq)
//	message, contactInfo, ret := gorm.UserContactService.GetFriendInfo(getFriendInfoReq.UserId)
//	JsonBack(c, message, ret, contactInfo)
//}

// DeleteContact 删除联系人
//func DeleteContact(c *gin.Context) {
//	var deleteContactReq request.DeleteContactRequest
//	if err := c.BindJSON(&deleteContactReq); err != nil {
//		zlog.Error(err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"code":    500,
//			"message": constants.SYSTEM_ERROR,
//		})
//		return
//	}
//	message, ret := gorm.UserContactService.DeleteContact(deleteContactReq.OwnerId, deleteContactReq.ContactId)
//	JsonBack(c, message, ret, nil)
//}

// ApplyContact 申请添加联系人
func ApplyFriend(c *gin.Context) {
	var req pb.ApplyFriendRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	rsp, err := grpcClient.ApplyFriend(ctx, &req)
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

// GetNewContactList 获取新的联系人申请列表
func GetNewApplyList(c *gin.Context) {
	var req pb.UserIdRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	rsp, err := grpcClient.GetApplyList(ctx, &req)
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

//func RespondApply(c *gin.Context) {
//	var req pb.RespondToApply
//	if err := c.BindJSON(&req); err != nil {
//		zlog.Error(err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"code":    500,
//			"message": constants.SYSTEM_ERROR,
//		})
//		return
//	}
//	rsp, err := grpcClient.RespondToApply(ctx, &req)
//	if err != nil {
//		zlog.Error(err.Error())
//		JsonBack(c, err)
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"data": rsp,
//	})
//}

// PassContactApply 通过联系人申请
func PassRelationApply(c *gin.Context) {
	var req pb.RespondToApply
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	rsp, err := grpcClient.RespondToApply(ctx, &req)
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
//// RefuseContactApply 拒绝联系人申请
//func RefuseRelationApply(c *gin.Context) {
//	var passContactApplyReq request.UserFriendRequest
//	if err := c.BindJSON(&passContactApplyReq); err != nil {
//		zlog.Error(err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"code":    500,
//			"message": constants.SYSTEM_ERROR,
//		})
//		return
//	}
//	message, ret := gorm.UserContactService.RefuseRelationApply(passContactApplyReq.UserId, passContactApplyReq.FriendId)
//	JsonBack(c, message, ret, nil)
//}

/*
// BlackContact 拉黑联系人
func BlackApply(c *gin.Context) {
	var req request.BlackContactRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.UserContactService.BlackContact(req.OwnerId, req.ContactId)
	JsonBack(c, message, ret, nil)
}

// CancelBlackContact 解除拉黑联系人
func CancelBlackContact(c *gin.Context) {
	var req request.BlackContactRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.UserContactService.CancelBlackContact(req.OwnerId, req.ContactId)
	JsonBack(c, message, ret, nil)
}

// GetAddGroupList 获取新的群聊申请列表
func GetAddGroupList(c *gin.Context) {
	var req request.AddGroupListRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, data, ret := gorm.UserContactService.GetAddGroupList(req.GroupId)
	JsonBack(c, message, ret, data)
}
*/
//
//// BlackApply 拉黑申请
//func BlackApply(c *gin.Context) {
//	var req request.UserFriendRequest
//	if err := c.BindJSON(&req); err != nil {
//		zlog.Error(err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"code":    500,
//			"message": constants.SYSTEM_ERROR,
//		})
//		return
//	}
//	message, ret := gorm.UserContactService.BlackApply(req.UserId, req.FriendId)
//	JsonBack(c, message, ret, nil)
//}
