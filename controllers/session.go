package controllers

import (
	"github.com/astaxie/beego"
	"lovehome/models"
	//"encoding/json"
	//"github.com/astaxie/beego/orm"
)

type SessionController struct {
	beego.Controller
}
func (this * SessionController)RetData(resp map[string]interface{}){
	this.Data["json"]=resp
	this.ServeJSON()
}
//type Area struct {//每个都要有一个判断所以用一个万能的map代替
//	errno int
//	errmsg string
//}
func (c *SessionController) GetSession() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	resp["errno"] =models.RECODE_DATAERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
	c.RetData(resp)
	user :=models.User{}
	name :=c.GetSession()

	if name != nil {
		user.Name = name.(string)
		resp["errno"] =models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		resp["data"] = user

	}

}
//func (this*SessionController)DeleteSessionData(){
//	resp := make(map[string]interface{})
//	defer this.RetData(resp)
//	this.DelSession("name")
//
//	resp["errno"] =models.RECODE_OK
//	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
//
//
//}

