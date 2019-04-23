package controllers

import (
	"os"
	"time"
	"github.com/astaxie/beego"
	"log"
	"github.com/fakegeass/notes/models"
	"encoding/json"
)

type NotesController struct{
	beego.Controller
}

func init(){
	log.SetOutput(os.Stdout)
}

func (this *NotesController)Get() {
	uuid:=this.GetString("UUID")
	if uuid!=""{if models.GetNotes(uuid)!=nil{
		log.Println("1.1")
		temp:=models.GetNotes(uuid)
		this.Data["UUID"]=temp.Uuid
		this.Data["Title"]=temp.Title
		this.Data["Val"]=temp.Data.Val
	}}

	//models.GetAll()
	temp :=make(map[string]*(models.Note))
	temp=models.GetAll()
	this.Data["Notes"]=temp
	this.TplName="notes.tpl"
	//result,_:=json.Marshal(temp)
	//this.Ctx.WriteString("Get all"+string(result))
}

func (this *NotesController)Post() {
	temp:=new(models.Note)
	uuid:=this.GetString("UUID")
	if uuid!=""{
		if ok:=(models.GetNotes(uuid));ok!=nil{
			temp=ok
		}
	}
	title:=this.GetString("Title")
	if title!=""{temp.Title=title}
	val:=this.GetString("Data.Val")
	temp.Date=time.Now()
	if val!=""{temp.Data.Val=val}
	if true{
	log.Printf("Get post from %v",uuid)
	log.Printf("Get post title is  %v",temp.Title)
	log.Printf("Get post Val is %v",temp.Data.Val)
}
	data,_:=json.Marshal(temp)
	err:=models.SetNotes(uuid,data)
	log.Printf("%#v\n",temp)
	if err!=nil{
		log.Printf("Post Error:%v;\nContent:%s",err,this.Ctx.Input.RequestBody)
	}
	log.Println(this.Ctx.Input.RequestBody)
	this.Redirect("/notes",302)
	//this.Get()
}

func (this *NotesController)Delete(){
	uuid:=this.GetString("UUID")
	err:=models.Delete(uuid)
	if err!=nil{
		log.Println("DELETE note fail for NoteID ",uuid,",reason:",err)
	}
	this.Redirect("/notes",302)
}