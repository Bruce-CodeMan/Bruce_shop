/**
 * @Author: Bruce
 * @Description: 描述
 * @Date: 2023-02-17 19:45
 */

package api

import (
	"Bruce_shop/api/user_web/forms"
	"Bruce_shop/api/user_web/global"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func GenerateSmsCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func SendSms(ctx *gin.Context) {
	smsForm := forms.SendSmsForm{}
	if err := ctx.ShouldBindJSON(&smsForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	smsCode := GenerateSmsCode(6)
	fmt.Println("发送的短信验证码为:", smsCode)
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	rdb.Set(context.Background(), smsForm.Mobile, smsCode, 120*time.Second)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":     "发送成功",
		"smsCode": smsCode,
	})
}
