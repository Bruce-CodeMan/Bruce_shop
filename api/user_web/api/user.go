/**
 * @Author: Bruce
 * @Description: some user functions
 * @Date: 2023/2/15 2:03 PM
 */

package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"Bruce_shop/api/user_web/forms"
	"Bruce_shop/api/user_web/global"
	"Bruce_shop/api/user_web/global/response"
	"Bruce_shop/api/user_web/middlewares"
	"Bruce_shop/api/user_web/models"
	"Bruce_shop/api/user_web/proto"
)

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}

// HandleGrpcErrorToHttp convert the gRPC response code to the Http response
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

// GetUserList get the user list info by pn/pSize in the browser
func GetUserList(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户ID: %d", currentUser.Id)

	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("pSize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)
	resp, err := global.UserSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询 [用户列表信息失败]")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	for _, v := range resp.Data {
		user := response.UserResponse{
			Id:       v.Id,
			NickName: v.NickName,
			Birthday: time.Unix(int64(v.BirthDay), 0).Format("2006-01-02"),
			Gender:   v.Gender,
			Mobile:   v.Mobile,
		}
		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)
}

// PasswordLogin User login the project by password/nickname
func PasswordLogin(ctx *gin.Context) {

	passwordLoginForm := forms.PasswordLoginForms{}
	if err := ctx.ShouldBindJSON(&passwordLoginForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	// true 代表每次证明之后都关闭
	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, true) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"captcha": "验证码错误",
		})
		return
	}
	// login logic
	if resp, err := global.UserSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	}); err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "用户不存在",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]string{
					"msg": "登录失败",
				})
			}
			return
		}
	} else {
		// check the user's password
		if passwordResp, passwordErr := global.UserSrvClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
			Password:         passwordLoginForm.Password,
			EncryptoPassword: resp.Password,
		}); passwordErr != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"password": "登录失败",
			})
		} else {
			if passwordResp.Success {
				// Generate token
				j := middlewares.NewJWT()
				claims := models.CustomClaims{
					Id:          uint(resp.Id),
					NickName:    resp.NickName,
					AuthorityId: uint(resp.Role),
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(),               // 签名的生效时间
						ExpiresAt: time.Now().Unix() + 60*60*24*30, // 过期时间
						Issuer:    "Bruce",                         // 签发机构
					},
				}
				token, err := j.CreateToken(claims)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"msg": "生成token失败",
					})
				}

				ctx.JSON(http.StatusOK, gin.H{
					"id":        resp.Id,
					"nickName":  resp.NickName,
					"token":     token,
					"expiredAt": (time.Now().Unix() + 60*60*24*60) * 1000,
				})
			} else {
				ctx.JSON(http.StatusBadRequest, map[string]string{
					"msg": "登录失败",
				})
			}
		}
	}

}

// Register User can register the system by mobile
func Register(ctx *gin.Context) {
	// User Register
	registerForms := forms.RegisterForms{}
	if err := ctx.ShouldBindJSON(&registerForms); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	// Code from redis database to check whether is correct or not
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	value, err := rdb.Get(context.Background(), registerForms.Mobile).Result()
	if err == redis.Nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "验证码错误",
		})
		return
	} else {
		if value != registerForms.Code {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": "验证码错误",
			})
		}
	}

	resp, err := global.UserSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		NickName: registerForms.Mobile,
		Mobile:   registerForms.Mobile,
		Password: registerForms.Password,
	})
	if err != nil {
		zap.S().Errorf("[Register]新建用户失败 %s", err.Error())
		HandleValidatorError(ctx, err)
		return
	}

	// Create the JWT
	j := middlewares.NewJWT()
	claims := models.CustomClaims{
		Id:          uint(resp.Id),
		NickName:    resp.NickName,
		AuthorityId: uint(resp.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*60*24*30,
			Issuer:    "Bruce",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":    "创建用户成功",
		"id":     resp.Id,
		"mobile": resp.Mobile,
		"token":  token,
	})

}
