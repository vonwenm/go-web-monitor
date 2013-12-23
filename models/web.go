package models

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"strconv"
)

type Web struct {
	Url   string
	Admin Admin
}

func (w *Web) Monitor() {
	req := httplib.Get(w.Url)
	str, err := req.Response()
	if err != nil {
		fmt.Println(err)
		context := "Dear Administrator，\n  Can't connent to the server(" + w.Url + "),please check it!\n  " + "Error message:" + err.Error()
		sendMail(w.Admin.Mail, "Server Error:Can't connect error", context)
	} else {
		fmt.Println(str.Status)
		if str.StatusCode != 200 {
			context := "Dear Administrator,\n  The error code is " + strconv.Itoa(str.StatusCode) + "\n  The error message is:" + str.Status
			sendMail(w.Admin.Mail, "Server Error: Response errer", context)
		}
	}
}

func sendMail(mail, subject, context string) {
	err := SendMail(AdminMail, AdminPass, AdminHost, mail, subject, context, "text")

	if err != nil {
		fmt.Println(AdminMail)
		fmt.Println(AdminPass)
		fmt.Println(AdminHost)
		fmt.Println("Send Mail Failed:", err)
	} else {
		fmt.Println(mail)
		fmt.Println("Send Mail Successfully")
	}
}
