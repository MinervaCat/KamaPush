package api

import (
	pb "KamaPush/pb"
	"KamaPush/pkg/zlog"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

var grpcClient pb.KamaChatClient
var ctx = context.Background()

func init() {
	conn, err := grpc.NewClient("101.200.184.252:9090", grpc.WithInsecure())
	if err != nil {
		zlog.Error(err.Error())
	}
	defer conn.Close()
	grpcClient = pb.NewKamaChatClient(conn)
}

func JsonBack(c *gin.Context, err error) {
	st, ok := status.FromError(err)
	if ok {
		// 处理 gRPC 错误
		switch st.Code() {
		case codes.NotFound:
			c.JSON(http.StatusOK, gin.H{
				"code":    404,
				"message": "查询对象不存在:" + st.Message(),
			})
		case codes.InvalidArgument:
			c.JSON(http.StatusOK, gin.H{
				"code":    400,
				"message": "参数错误:" + st.Message(),
			})
		case codes.Internal:
			c.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "服务器内部错误:" + st.Message(),
			})
		default:
			c.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "其他错误:" + st.Message(),
			})
		}
	} else {
		// 非 gRPC 错误（如网络错误）
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "网咯错误:" + st.Message(),
		})

	}
}
