package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
	"github.com/wonderstone/awei-proxy"
)

func main() {

	// test getTodayDate
	month, day := getTodayDate()
	fmt.Println(month, day)

	c := colly.NewCollector()
	//
	// url = "https://agit.ai/12/a/raw/branch/master/2/2.16v2"
	// //  print all data from the url
	// c.OnResponse(func(r *colly.Response) {
	// 	fmt.Println(string(r.Body))

	// })

	// c.Visit(url)

	// new url1
	url1 := "https://agit.ai/12/a/src/branch/master/4"

	// get all tr from the html store its td's href in the slice
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		href := e.ChildAttr("td a", "href")
		fmt.Println(href)
	})

	// visit the url1
	c.Visit(url1)

}

// get today date
func getTodayDate() (month string, day string) {
	t := time.Now()
	return t.Format("1"), t.Format("2")
}
