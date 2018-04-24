package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	"lovehome/models"
	//_ "github.com/go-sql-driver/mysql"

	//"encoding/json"
	//"github.com/astaxie/beego/cache"
	//"time"
)

type HouseIndexController struct {
	beego.Controller
}
func (this * HouseIndexController)RetData(resp map[string]interface{}){
	this.Data["json"]=resp
	this.ServeJSON()
}
//type Area struct {//每个都要有一个判断所以用一个万能的map代替
//	errno int
//	errmsg string
//}
func (c *HouseIndexController) HouseIndex() {
	resp := make(map[string]interface{})

	resp["errno"] =models.RECODE_DATAERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
	c.RetData(resp)

}
