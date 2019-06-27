package sale_stock

import (
	"github.com/hexya-addons/web/controllers"
	"github.com/hexya-erp/hexya/src/server"
)

const MODULE_NAME string = "sale_stock"

func init() {
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PreInit:  func() {},
		PostInit: func() {},
	})
	controllers.BackendJS = append(controllers.BackendJS,
		"/static/sale_stock/src/js/tour.js",
	)

}
