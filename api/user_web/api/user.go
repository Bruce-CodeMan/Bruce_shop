/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/15 2:03 PM
 */

package api

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"Bruce_shop/api/user_web/proto"
)

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	// 将gRPC的code转换成Http的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误",
				})
			}
			return
		}
	}
}

func GetUserList(ctx *gin.Context) {
	ip := "127.0.0.1"
	port := 50051
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 [用户服务失效]",
			"msg", err.Error())
	}
	c := proto.NewUserClient(conn)
	resp, err := c.GetUserList(context.Background(), &proto.PageInfo{Pn: 0, PSize: 0})
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询 [用户列表信息失败]")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	for _, v := range resp.Data {
		data := make(map[string]interface{})
		data["id"] = v.Id
		data["name"] = v.NickName
		data["birthday"] = v.BirthDay
		data["gender"] = v.Gender
		data["mobile"] = v.Mobile
		result = append(result, data)
	}
	ctx.JSON(http.StatusOK, result)
}