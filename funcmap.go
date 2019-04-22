package main

import(
	"time"
	"os"
	"strconv"
	"github.com/astaxie/beego"
)

func addFuncMap(){
	beego.AddFuncMap("format", format)
	beego.AddFuncMap("isFile", isFile)
}

func format(in time.Time)string{
	return in.Format("2006-01-02 15:04:05")
}

func isFile(file os.FileInfo)string{
	if file.IsDir()==true{
		return "-"
	}else{
		return strconv.Itoa(int(file.Size()))
	}
}