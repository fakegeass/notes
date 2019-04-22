package controllers

import(
	"io/ioutil"
	"log"
	"github.com/astaxie/beego"
	"os"
)

type FileController struct{
	beego.Controller
}

func (this *FileController)Get(){
	this.get("")
}

func (this *FileController)get(filepath string){
	file,err:=ioutil.ReadDir("./file/"+filepath)
	if err!=nil{
		log.Println("Read fail,",err)
	}
	this.Data["file"]=file
	this.TplName="file.tpl"
}

func (this *FileController)Upload(){
	f, h, err := this.GetFile("upload")
    if err != nil {
		log.Println("getfile err ", err)
		this.Redirect("/file",200)
		return
    }
    defer f.Close()
	this.SaveToFile("upload", "./file/"+h.Filename)
	this.Redirect("/file",302)
}

//ChDir used to change the dict
//TODO:二级目录
func (this *FileController)ChDir(){
	file:="./file/"+this.Ctx.Input.Param(":dir")
	fileInfo,err:=os.Stat(file)
	if err!=nil{
		this.Redirect("/file",404)
		return
	}
	if fileInfo.IsDir()!=true{
		this.Ctx.Output.Download(file)
		return
	}
	log.Println("Change dir to ",file)
	this.get(file)
}