package routers

import (
	"lovehome/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	//api/v1.0/areas
	beego.Router("/api/v1.0/areas", &controllers.AreaController{},"get:GetArea")
	///api/v1.0/houses/index
	beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{},"get:HouseIndex")
	///api/v1.0/session
	beego.Router("/api/v1.0/session", &controllers.SessionController{},"get:GetSession")

}
