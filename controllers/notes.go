package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"notes/models"
	//"encoding/json"
)

type NotesController struct{
	beego.Controller
}

func (this *NotesController)Get() {
	log.Print()

	models.GetAll()
	temp :=make(map[string]*(models.Note))
	temp=models.GetAll()
	this.Data["Notes"]=temp
	this.TplName="notes.tpl"
	//result,_:=json.Marshal(temp)
	//this.Ctx.WriteString("Get all"+string(result))
}

func (this *NotesController)Post() {
	uuid:=this.GetString("Uuid")
	log.Printf("Get post from %v",uuid)
	err:=models.SetNotes(uuid,this.Ctx.Input.RequestBody)
	log.Printf("Post Error:%v",err)
	log.Println(this.Ctx.Input.RequestBody)
	this.Ctx.WriteString("Set!")
}