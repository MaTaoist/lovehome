package controllers

import "github.com/astaxie/beego"


type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	this.Ctx.WriteString("getinfo responce succeess")
}

func(this*UserController) GetInfo(){
	id := this.Ctx.Input.Param(":id")
	this.Ctx.WriteString("getInfo data ,id = "+id)
}

func(this*UserController) GetNum(){
	//ext := this.Ctx.Input.Param(":ext")
	//path:=this.Ctx.Input.Param(":path")
	splate:=this.Ctx.Input.Param(":splat")
	this.Ctx.WriteString("getNum data ,username = "+splate)
}

func (this* UserController)PostData(){
	this.Ctx.WriteString("this is post func")
}
