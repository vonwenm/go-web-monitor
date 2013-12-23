package services

import (
	"fmt"
	"io/ioutil"
	"monitor/models"
	"strings"
	"time"
)

var webs []models.Web

/**
 *init function.
 *read uri from local file,
 *and start the monitor server.
 */
func MonitorInit() {
	readUriFromFile()
	startTiker()
}

/**
 *read uri from local file
 */
func readUriFromFile() {
	b, err := ioutil.ReadFile("views/urls.txt")
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
		webs[i] = *web
	}
}

/**
 * init a tiker to start monitor every 1 minutes
 */
func startTiker() {
	startMonitor()
	ticker := time.NewTicker(20 * 60 * time.Second)
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
