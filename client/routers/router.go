package routers

import (
	"iCenter-client/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// APIs := beego.NewNamespace("/api",
	// 	beego.NSNamespace("/pods",
	// 		beego.NSInclude(
	// 			&controllers.Pod{},	
	// 		),
	// 	),
	// )
	// beego.AddNamespaces(APIs)
	beego.Router("/pods",&controllers.PodController{})
	beego.Run()
}
