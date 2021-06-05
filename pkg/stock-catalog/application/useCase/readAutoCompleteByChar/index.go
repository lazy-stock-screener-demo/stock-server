package readautocompletebychar

import (
	routing "github.com/qiangxue/fasthttp-routing"
)

var con = newController()

// Execute return a function to fasthttp router
func Execute() func(c *routing.Context) error {
	return func(c *routing.Context) error {
		con.executeImpl(c)
		return nil
	}
}
