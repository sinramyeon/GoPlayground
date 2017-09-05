package Selenium

import (
	"fmt"
	"log"
	"strings"

	"github.com/sclevine/agouti"
)

func SeleniumTest() {

	// 1. First, You should install agouti series.
	//For macOs
	/*$ brew install selenium-server-standalone
	$ brew insatll phantomjs
	$ brew install chromedriver
	$ go get github.com/sclevine/agouti
	$ go run main.go
	*/
	//For Windows
	/* you should puth chromdriver.... and so on to your %PATH%
	I used to use chrome, but selenium is best for firefox.
	But no IE please.
	*/
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalln("[ERROR]", err)
	}

	defer recover()
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("phantomjs"))
	if err != nil {
		log.Fatalln("[ERROR]", err)
	}

	if err := page.Navigate("http://ione.interpark.com/"); err != nil {
		log.Fatalln("[ERROR]", err)
	}

	ID := ""
	PW := ""
	page.FindByID("Username").SendKeys(ID)
	page.FindByID("Password").SendKeys(PW)

	page.FindByClass("loginSubmit").Click()

	if err := page.Navigate("http://ione.interpark.com/gw/app/bult/bbs00000.nsf/wviwnotice?ReadViewEntries&start=1&count=14&restricttocategory=03&page=1&&_=1504081645868"); err != nil {
		log.Fatalln("[ERROR]", err)
	}

	cookie, err := page.GetCookies()
	if err != nil {
		log.Fatalln("[ERROR]", err)
	}

	for _, v := range cookie {
		if strings.Contains(v.Name, "LtpaToken") {
			fmt.Println(v)
			fmt.Println(v.Value)
		}
	}

}
