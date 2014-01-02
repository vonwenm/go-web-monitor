package services

import (
	"fmt"
	"go-web-monitor/models"
	"io/ioutil"
	"strings"
	"time"
)

var webs []models.Web

var adminMail string
var adminPwd string
var adminMailHost string

/**
 *init function.
 *read uri from local file,
 *and start the monitor server.
 */
func MonitorInit() {
	readSystemAdminMailConfFormFile()
	readUriFromFile()
	startTiker()
}

/**
 *read uri from local file
 */
func readUriFromFile() {
	b, err := ioutil.ReadFile("conf/urls.conf")
	if err != nil {
		panic(err)
	}

	urls := string(b)
	var urlArray = strings.Split(urls, "\n")
	fmt.Println(len(urlArray))
	webs = make([]models.Web, len(urlArray))
	for i, url := range urlArray {
		fmt.Println(url)
		eArray := strings.Split(url, "#")

		web := new(models.Web)
		admin := new(models.Admin)
		admin.Mail = eArray[1]
		web.Url = eArray[0]
		web.Admin = *admin

		systemAdmin := new(models.SystemAdmin)
		systemAdmin.SystemMail = adminMail
		systemAdmin.SystemPwd = adminPwd
		systemAdmin.SystemHost = adminMailHost
		web.SystemAdmin = *systemAdmin
		webs[i] = *web
	}
}

/**
 *read system mail form local conf file
 */
func readSystemAdminMailConfFormFile() {
	b, err := ioutil.ReadFile("conf/admin.conf")
	if err != nil {
		panic(err)
	}

	adminConfs := string(b)
	var confArray = strings.Split(adminConfs, "\n")
	fmt.Println(len(confArray))
	for _, array := range confArray {
		cArray := strings.Split(array, "#")

		switch cArray[0] {
		case "adminMail":
			adminMail = cArray[1]
		case "adminPwd":
			adminPwd = cArray[1]
		case "mailHost":
			adminMailHost = cArray[1]
		}
	}
}

/**
 * init a tiker to start monitor every 1 minutes
 */
func startTiker() {
	startMonitor()
	ticker := time.NewTicker(60 * time.Second)
	quit := make(chan struct{})
	var i int
	go func() {
		for {
			select {
			case <-ticker.C:
				i++
				fmt.Println("Tick at", i)
				startMonitor()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func startMonitor() {
	for _, web := range webs {
		web.Monitor()
	}
}
