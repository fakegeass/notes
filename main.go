package main

import (
	"log"
	_ "github.com/fakegeass/notes/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/google/uuid"
)

var FilterUser = func(ctx *context.Context) {
	sessionID,ok := ctx.Input.Session("uid").(uuid.UUID)
  if !ok&& ctx.Request.RequestURI != "/login" {
		log.Printf("Session is illegal(%v).",sessionID)
		ctx.Redirect(302, "/login")
	}
}

func main() {
	addFuncMap()
	beego.InsertFilter("/*",beego.BeforeRouter,FilterUser)
	beego.BConfig.WebConfig.Session.SessionOn = true
	log.Println("Welcome to my notes!")
	beego.Run()
}