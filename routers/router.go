package routers

import (
	"github.com/fakegeass/notes/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/notes", &controllers.NotesController{})
	beego.Router("/notes/delete", &controllers.NotesController{},"get:Delete")
	beego.Router("/file",&controllers.FileController{})
	beego.Router("/file/?:dir",&controllers.FileController{},"*:ChDir")
	beego.Router("/upload",&controllers.FileController{},"*:Upload")
}
