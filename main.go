package main

import (
	"log"
	_ "notes/routers"
	"github.com/astaxie/beego"
)


func main() {
	log.Println("Welcome to my notes!")
	beego.Run()
}

