/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/17 2:08 PM
 */
package models

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	Id          uint
	NickName    string
	AuthorityId uint
	jwt.StandardClaims
}
