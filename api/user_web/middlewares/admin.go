/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/17 3:27 PM
 */
package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Bruce_shop/api/user_web/models"
)

func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*models.CustomClaims)

		if currentUser.AuthorityId != 2 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg": "无权限",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
