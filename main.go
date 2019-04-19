package main

import (
	"log"
	_ "notes/routers"
	"github.com/astaxie/beego"
	"time"
)


func main() {
	log.Println("Welcome to my notes!")
	beego.AddFuncMap("format", format)
	beego.Run()
}

func format(in time.Time)string{
	return in.Format("2006-01-02 15:04:05")
}

