/**
 * @Author: Bruce
 * @Description: 描述
 * @Date: 2023-02-17 21:06
 */

package forms

type SendSmsForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Type   int    `form:"type" json:"type" binding:"required,oneof=1 2"`
}
