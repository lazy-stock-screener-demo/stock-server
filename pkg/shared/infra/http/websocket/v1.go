package websocketapi

import (
	stockcatalogws "stock-contexts/pkg/stock-catalog/infra/http/websocket"

	"github.com/fasthttp/websocket"
	routing "github.com/qiangxue/fasthttp-routing"
)

var upgrader = websocket.FastHTTPUpgrader{}

// V1Router define websocket
func V1Router(router *routing.Router) *routing.Router {
	hub := stockcatalogws.NewHub()
	router.Any("/ws", func(ctx *routing.Context) error {
		upgrader.Upgrade(ctx.RequestCtx, func(conn *websocket.Conn) {
			client := &stockcatalogws.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
			client.Hub.Register <- client

			go client.WritePump()
			client.ReadPump()
		})
		return nil
	})
	return router
}
