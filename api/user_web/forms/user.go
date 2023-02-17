/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/16 1:44 PM
 */

package forms

type PasswordLoginForms struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password  string `form:"password" json:"password" binding:"required,min=3,max=10"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}

type RegisterForms struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=10"`
	Code     string `form:"code" json:"code" binding:"required,min=6,max=6"`
}
