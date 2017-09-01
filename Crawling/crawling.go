package Crawling

import (
	"github.com/PuerkitoBio/goquery"
)

// Find Today's packtpub Free book
func PacktFreeBook() string {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	// 1. Get url's html
	doc, _ := goquery.NewDocument("https://www.packtpub.com/packt/offers/free-learning")
	// 2. Find thing
	// I'm finding class name "dotd-title" with <h2></h2>
	freebook := doc.Find(".dotd-title").Find("h2").Text()

	return freebook
}

// Find Daily Github Go Opensource
func GoScrape() map[string]string {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	// 1. Get github's html
	doc, _ := goquery.NewDocument("https://github.com/trending/go?since=daily")

	githublist := make(map[string]string)

	// EachWithBreak can return true/false to stop searching
	// There were too many of repositories, so I stopped at count 5
	// I apolozise for this bad style of for loop

	var forLoop int = 0

	// 2. Find thing
	// I'm finding class name repo-list with <li></li>
	doc.Find(".repo-list li").EachWithBreak(func(i int, s *goquery.Selection) bool {

		if forLoop > 4 {
			return false
		} else {
			// 3. In <li>s, Let's find <h3></h3> and an <a href> in <h3>
			// If there is no urls in a href="", you can set default value
			// title is repository name and desc is description
			title := s.Find("h3 a").AttrOr("href", "default")
			desc := s.Find(".py-1 p").Text()

			githublist[title] = desc
			forLoop++
			return true
		}
	})

	return githublist
}

// Find Okky Tech writings
func OkkyScrape() map[string]string {

	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	// 1. Get OKKY.KR's..한국말써야징 오키 html을 갖고옵니다.

	// 크롬 개발자도구를 잘 이용하고 있습니다.
	// F12를 눌러서 조의를 표...아니 Elements를 보면 해당 페이지의 html을 볼 수 있습니다.
	doc, _ := goquery.NewDocument("https://okky.kr/")

	okkylist := make(map[string]string)

	// .Find(".class tag .class")
	// .Find("a").Text() a 태그 안에 든 거를 찾습니다.
	//여기서는 article-middle-block 클래스 안에 있는 걸 찾습니다. (여러개 있어서 Each로 돌림)
	doc.Find(".article-middle-block").Each(func(i int, s *goquery.Selection) {

		//<h5>안에 제목이 있으니 제목을 찾고
		title := s.Find("h5").Text()
		// url 전체가 아니라 뒷 부분만 있어서 요렇게 url을 만듭니다.
		url := "https://okky.kr" + s.Find("h5 a").AttrOr("href", "없음")

		okkylist[title] = url

	})
	// 이렇게 하면 글제목 인덱스 안 글주소 밸류가 든 맵을 만들 수 있습니다.
	return okkylist

}

// IT월드 뉴스 찾기
func NewsScrape() map[string]string {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	// 다양한 사이트를 크롤링하는걸 실습해보세요!
	doc, _ := goquery.NewDocument("http://www.itworld.co.kr/")

	newslist := make(map[string]string)

	var forLoop int = 0

	doc.Find(".cio_summary").EachWithBreak(func(i int, s *goquery.Selection) bool {

		if forLoop > 4 {
			return false
		} else {
			title := s.Find("ul li a").Text()
			url := s.Find("ul li a").AttrOr("href", "없음")

			newslist[title] = url
			forLoop++
			return true
		}
	})
	return newslist
}
