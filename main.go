package main

import (
	"github.com/astaxie/beego"
	"monitor/controllers"
	"monitor/services"
)

func main() {
	//start moitor
	services.MonitorInit()

	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
