package main

import (
	"github.com/astaxie/beego"
	"go-web-monitor/controllers"
	"go-web-monitor/services"
)

func main() {
	//start moitor
	services.MonitorInit()

	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
