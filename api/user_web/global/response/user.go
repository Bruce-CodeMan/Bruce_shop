/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/15 3:38 PM
 */
package response

type UserResponse struct {
	Id       int32  `json:"id"`
	Mobile   string `json:"mobile"`
	NickName string `json:"nickname"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
}
