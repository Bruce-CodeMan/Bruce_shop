/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/15 2:00 PM
 */

package main

import "Bruce_shop/api/user_web/inintialize"

func main() {
	// Step 1 ,initialize the router
	Router := inintialize.Routers()
	Router.Run()
}
