package isauthed

import (
	routing "github.com/qiangxue/fasthttp-routing"
)

var conCRUD = newControllerCRUD()

func ExecuteCRUD() func(*routing.Context) error {
	return func(c *routing.Context) error {
		conCRUD.executeImpl(c)
		return nil
	}
}
