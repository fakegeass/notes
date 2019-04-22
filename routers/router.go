package routers

import (
	"notes/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/notes", &controllers.NotesController{})
	beego.Router("/notes/delete", &controllers.NotesController{},"get:Delete")
	beego.Router("/download",&controllers.FileController{})
	beego.Router("/download/file?name=*",&controllers.FileController{},"*:Download")
	beego.Router("/download/?:dir",&controllers.FileController{},"*:ChDir")
}
