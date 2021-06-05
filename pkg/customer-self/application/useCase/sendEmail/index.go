package sendEmail

import (
	routing "github.com/qiangxue/fasthttp-routing"
)

var con = newController()

// Execute return a function to fasthttp router
func ExecuteCRUD() func(c *routing.Context) error {
	return func(c *routing.Context) error {
		con.executeImpl(c)
		return nil
	}
}
