package restfulapi

import (
	"net"
	"os"
	middleware "stock-contexts/pkg/shared/infra/http/middle"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

// App define struct
var App *AppStruct = nil

// AppStruct struct
type AppStruct struct {
	ln      net.Listener
	handler *routing.Router
}

// UseRouter define router
func (app *AppStruct) UseRouter(handlerFunc func(*routing.Router) *routing.Router) {
	app.handler = handlerFunc(app.getRouter())
}

func (app *AppStruct) getRouter() *routing.Router {
	return app.handler
}

// CreateServer create a server
func (app *AppStruct) CreateServer() *fasthttp.Server {
	// requestHandler := V1Router(app.handler)
	requestHandler := app.getRouter()
	return &fasthttp.Server{
		Handler: requestHandler.HandleRequest,
		Name:    "V1 Router",
	}
}

// GetListener get a Listener
func (app *AppStruct) GetListener() net.Listener {
	return app.ln
}

// NewApp func
func NewApp() *AppStruct {
	if App == nil {
		hostURL := os.Getenv("HOST_URL")
		ln, _ := net.Listen("tcp", hostURL)
		router := routing.New()
		router.Use(middleware.CORS())
		App = &AppStruct{
			ln:      ln,
			handler: router,
		}
		return App
	}
	return App
}
