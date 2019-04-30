package controllers

import (
	"log"
	"github.com/astaxie/beego"
	"github.com/google/uuid"
	"github.com/fakegeass/notes/models"
)

type LoginController struct{
	beego.Controller
}

func (c *LoginController)Get(){
	c.TplName="login.tpl"
}

func (c *LoginController)Post(){
	user:=c.GetString("user")
	pass:=c.GetString("pass")
	if ok:=models.JudgePassword(user,pass);!ok{
		log.Printf("%v login fail!",user)
		c.Redirect("/login",403)
		return
	}
	temp:=uuid.New()
	c.SetSession("uid",temp)
	log.Println("set uid :",c.GetSession("uid"))
	log.Printf("%v login success!",user)
	c.Redirect("/",302)
}