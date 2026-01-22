package api

import (
	pb "KamaPush/pb"
	"KamaPush/pkg/constants"
	"KamaPush/pkg/zlog"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register 注册
func Register(c *gin.Context) {
	var registerReq pb.RegisterRequest
	zlog.Info("开始登陆")
	if err := c.BindJSON(&registerReq); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	zlog.Info(fmt.Sprintf("the req is :%v", registerReq))
	rsp, err := grpcClient.Register(ctx, &registerReq)
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

// Login 登录
func Login(c *gin.Context) {
	var loginReq pb.LoginRequest
	if err := c.BindJSON(&loginReq); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	rsp, err := grpcClient.Login(ctx, &loginReq)
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

// SmsLogin 验证码登录
//func SmsLogin(c *gin.Context) {
//	var req request.SmsLoginRequest
//	if err := c.BindJSON(&req); err != nil {
//		zlog.Error(err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"code":    500,
//			"message": constants.SYSTEM_ERROR,
//		})
//		return
//	}
//	message, userInfo, ret := gorm.UserInfoService.SmsLogin(req)
//	JsonBack(c, message, ret, userInfo)
//}

// UpdateUserInfo 修改用户信息
//func UpdateUserInfo(c *gin.Context) {
//	var req request.UpdateUserInfoRequest
//	if err := c.BindJSON(&req); err != nil {
//		zlog.Error(err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"code":    500,
//			"message": constants.SYSTEM_ERROR,
//		})
//		return
//	}
//	message, ret := gorm.UserInfoService.UpdateUserInfo(req)
//	JsonBack(c, message, ret, nil)
//}

// GetUserInfoList 获取用户列表
//func GetUserInfoList(c *gin.Context) {
//	var req request.UserRequest
//	if err := c.BindJSON(&req); err != nil {
//		zlog.Error(err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"code":    500,
//			"message": constants.SYSTEM_ERROR,
//		})
//		return
//	}
//
//	message, userList, ret := gorm.UserInfoService.GetUserInfoList(req.UserId)
//	JsonBack(c, message, ret, userList)
//}

// AbleUsers 启用用户
//func AbleUsers(c *gin.Context) {
//	var req request.AbleUsersRequest
//	if err := c.BindJSON(&req); err != nil {
//		zlog.Error(err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"code":    500,
//			"message": constants.SYSTEM_ERROR,
//		})
//		return
//	}
//	message, ret := gorm.UserInfoService.AbleUsers(req.UuidList)
//	JsonBack(c, message, ret, nil)
//}
//
//// DisableUsers 禁用用户
//func DisableUsers(c *gin.Context) {
//	var req request.AbleUsersRequest
//	if err := c.BindJSON(&req); err != nil {
//		zlog.Error(err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"code":    500,
//			"message": constants.SYSTEM_ERROR,
//		})
//		return
//	}
//	message, ret := gorm.UserInfoService.DisableUsers(req.UuidList)
//	JsonBack(c, message, ret, nil)
//}

// GetUserInfo 获取用户信息
//func GetUserInfo(c *gin.Context) {
//	var req request.UserRequest
//	if err := c.BindJSON(&req); err != nil {
//		zlog.Error(err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"code":    500,
//			"message": constants.SYSTEM_ERROR,
//		})
//		return
//	}
//	message, userInfo, ret := gorm.UserInfoService.GetUserInfo(req.UserId)
//	JsonBack(c, message, ret, userInfo)
//}
//
//// DeleteUsers 删除用户
//func DeleteUsers(c *gin.Context) {
//	var req request.AbleUsersRequest
//	if err := c.BindJSON(&req); err != nil {
//		zlog.Error(err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"code":    500,
//			"message": constants.SYSTEM_ERROR,
//		})
//		return
//	}
//	message, ret := gorm.UserInfoService.DeleteUsers(req.UuidList)
//	JsonBack(c, message, ret, nil)
//}
//
//// SetAdmin 设置管理员
//func SetAdmin(c *gin.Context) {
//	var req request.AbleUsersRequest
//	if err := c.BindJSON(&req); err != nil {
//		zlog.Error(err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"code":    500,
//			"message": constants.SYSTEM_ERROR,
//		})
//		return
//	}
//	message, ret := gorm.UserInfoService.SetAdmin(req.UuidList, req.IsAdmin)
//	JsonBack(c, message, ret, nil)
//}
//
//// SendSmsCode 发送短信验证码
//func SendSmsCode(c *gin.Context) {
//	var req request.SendSmsCodeRequest
//	if err := c.BindJSON(&req); err != nil {
//		zlog.Error(err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"code":    500,
//			"message": constants.SYSTEM_ERROR,
//		})
//		return
//	}
//	message, ret := gorm.UserInfoService.SendSmsCode(req.Telephone)
//	JsonBack(c, message, ret, nil)
//}
