package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"lovehome/models"
	//_ "github.com/go-sql-driver/mysql"

	//"encoding/json"
	//"github.com/astaxie/beego/cache"
	//"time"
)

type AreaController struct {
	beego.Controller
}
func (this * AreaController)RetData(resp map[string]interface{}){
	this.Data["json"]=resp
	this.ServeJSON()
}
//type Area struct {//每个都要有一个判断所以用一个万能的map代替
//	errno int
//	errmsg string
//}
func (c *AreaController) GetArea() {
	beego.Info("ok+++++++++++++++++++++++++++++")
	resp := make(map[string]interface{})
	//resp["errno"] = models.RECODE_OK
	//resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	//defer c.RetData(resp)
	//从mysql中拿数据
	var area  []models.Area

	o := orm.NewOrm()
	num ,err :=o.QueryTable("area").All(&area)

	if err != nil {
		beego.Info("err++++++++++++++++++++++++++++")
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	if num ==0{
		beego.Info("num+++++++++++++++++++++++++++++")
		resp["errno"] = 4002
		resp["errmsg"] = "没有查到数据"
		return
	}
	resp["errno"]=0
	resp["errmsg"]="ok"
	resp["data"]=&area
	/*json_str,err:=	json.Marshal(resp)
	c.Ctx.WriteString(json_str)*/
	beego.Info("+========================",resp)
	c.RetData(resp)


}