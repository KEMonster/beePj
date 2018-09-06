package routers

import (
	"beePj/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1", //创建域v1
		beego.NSNamespace("/order", //创建子域order
			beego.NSRouter("/addorder",&controllers.OrderController{},"post:AddOrder"), //注册路由 /v1/order/addorder 为post 请求
			beego.NSRouter("/searchorder",&controllers.OrderController{},"get:SearchOrder"), //注册路由 /v1/order/searchorder 为get 请求
			beego.NSRouter("/updateorder",&controllers.OrderController{},"get:UpdateOrder"),
			beego.NSRouter("/deleteorder",&controllers.OrderController{},"delete:DeleteOrder"), //注册路由 /v1/order/deleteorder 为delete 请求
		),
			
	)
	beego.AddNamespace(ns) //启用域
}