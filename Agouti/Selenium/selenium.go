package Selenium

import (
	"fmt"
	"log"
	"strings"

	"github.com/sclevine/agouti"
)

func SeleniumLoginTest() {

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

	//Get chrome driver(...or your browser driver)
	driver := agouti.ChromeDriver()
	//Start browser
	if err := driver.Start(); err != nil {
		log.Fatalln("[ERROR]", err)
	}
	defer driver.Stop()

	//Make new page with browser
	page, err := driver.NewPage(agouti.Browser("phantomjs"))
	if err != nil {
		log.Fatalln("[ERROR]", err)
	}

	//Go into some page
	if err := page.Navigate("somewhere"); err != nil {
		log.Fatalln("[ERROR]", err)
	}

	// Test Login
	ID := "your ID"
	PW := "your PW"
	page.FindByID("Some #id").SendKeys(ID)
	page.FindByID("Some #pw").SendKeys(PW)

	page.FindByClass(".loginbutton").Click()

	if err := page.Navigate("somewhere"); err != nil {
		log.Fatalln("[ERROR]", err)
	}

	// You can get page cookies to get some token or information
	cookie, err := page.GetCookies()
	if err != nil {
		log.Fatalln("[ERROR]", err)
	}

	for _, v := range cookie {
		if strings.Contains(v.Name, "value you want to find") {
			fmt.Println(v.Value)
		}
	}

	// Or just see html
	html, err := page.HTML()
	if err != nil {
		log.Fatalln("[ERROR]", err)
	}
	fmt.Println(html)

	// Or took screenshot
	page.Screenshot("/save.jpg")

	// Or move your mouse to make something special!
	err = page.MoveMouseBy(32, 45)
	if err != nil {
		log.Fatalln("[ERROR]", err)
	}
}
