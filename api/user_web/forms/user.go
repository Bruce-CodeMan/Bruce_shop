/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/16 1:44 PM
 */
package forms

type PasswordLoginForms struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=10"`
}
