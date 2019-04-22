package main

import (
	"log"
	_ "notes/routers"
	"github.com/astaxie/beego"
)

func main() {
	addFuncMap()
	log.Println("Welcome to my notes!")
	beego.Run()
}