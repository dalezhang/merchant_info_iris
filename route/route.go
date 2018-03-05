package route

import (
	"github.com/dalezhang/merchant_info_iris/config"
	"github.com/dalezhang/merchant_info_iris/controller/merchant"
	"github.com/dalezhang/merchant_info_iris/controller/sub_merchant"
	"github.com/kataras/iris"
)

func Route(app *iris.Application) {
	apiPrefix := config.ServerConfig.APIPrefix

	routers := app.Party(apiPrefix)
	{
		routers.Get("/merchants/:id", merchant.Show)
		routers.Get("/merchants", merchant.Index)

		routers.Post("/sub_merchants", sub_merchant.Create)
	}
}
