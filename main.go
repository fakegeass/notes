package main

import (
	"log"
	_ "github.com/fakegeass/notes/routers"
	"github.com/astaxie/beego"
)

func main() {
	addFuncMap()
	log.Println("Welcome to my notes!")
	beego.Run()
}