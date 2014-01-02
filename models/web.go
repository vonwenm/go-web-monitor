package models

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"strconv"
)

type Web struct {
	Url         string
	Admin       Admin
	SystemAdmin SystemAdmin
}

type SystemAdmin struct {
	SystemMail string
	SystemPwd  string
	SystemHost string
}

func (w *Web) Monitor() {
	req := httplib.Get(w.Url)
	str, err := req.Response()
	if err != nil {
		fmt.Println(err)
		context := "Dear Administrator，\n  Can't connent to the server(" + w.Url + "),please check it!\n  " + "Error message:" + err.Error()
		sendMail(w.SystemAdmin, w.Admin.Mail, "Server Error:Can't connect error", context)
	} else {
		fmt.Println(str.Status)
		if str.StatusCode != 200 {
			context := "Dear Administrator,\n  The error code is " + strconv.Itoa(str.StatusCode) + "\n  The error message is:" + str.Status
			sendMail(w.SystemAdmin, w.Admin.Mail, "Server Error: Response errer", context)
		}
	}
}

func sendMail(sytemAdmin SystemAdmin, mail, subject, context string) {
	err := SendMail(sytemAdmin.SystemMail, sytemAdmin.SystemPwd, sytemAdmin.SystemHost, mail, subject, context, "text")

	if err != nil {
		fmt.Println(sytemAdmin.SystemMail)
		fmt.Println(sytemAdmin.SystemPwd)
		fmt.Println(sytemAdmin.SystemHost)
		fmt.Println("mail:" + mail)
		fmt.Println("Send Mail Failed:", err)
	} else {
		fmt.Println(mail)
		fmt.Println("Send Mail Successfully")
	}
}
