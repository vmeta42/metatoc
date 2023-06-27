package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"gitlab.dev.21vianet.com/liu.hao8/metatoc-service/controllers"
)

func init() {
	ns := beego.NewNamespace("/metatoc-service",
		beego.NSNamespace("/v1",
			beego.NSNamespace("/metadata",
				beego.NSInclude(&controllers.MetadataController{})),
			beego.NSNamespace("/blockchain",
				beego.NSInclude(&controllers.BlockchainController{}),
			),
		),
	)
	beego.AddNamespace(ns)
}
