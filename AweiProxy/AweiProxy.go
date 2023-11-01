package AweiProxy

import (
	"net/url"
	"regexp"

	"github.com/gocolly/colly"
)

// define a variable to store the url and initialize it
var URL string = "https://agit.ai/12/a/src/branch/master/4"

// func to use regex to match month*day* and return the matched string
func IfMonthDayFileName(month string, date string, target string) bool {
	// regex to match month*day*
	regex := regexp.MustCompile(`(?m)` + month + `.*` + date + `.*` + `v` + `.*`)
	// return the matched string
	return regex.MatchString(target)
}

func IfMonthDayUpdated(month string, date string, target string) bool {
	// regex to match month*day*
	regex := regexp.MustCompile(`(?m)` + month + `.*` + date + `.*`)
	// return the matched string
	return regex.MatchString(target)
}

// Get file
func GetFile(outerurl string, month string, date string) (updated bool, res string) {
	c1l := colly.NewCollector()
	c2l := colly.NewCollector()
	c3l := colly.NewCollector()
	// get all tr from the html store whose td's href's value is regex month*day*
	c1l.OnHTML("tr", func(e *colly.HTMLElement) {
		// get href value from the first td
		href := e.ChildAttr("td:nth-child(1) a", "href")
		// get the first td text
		a := e.ChildText("td:nth-child(1)")
		if IfMonthDayUpdated(month, date, a) {
			updated = true
		}
		// if a is 4.23v2 then visit the href and return the Response
		if IfMonthDayFileName(month, date, a) {
			base, _ := url.Parse(outerurl)
			absolute := base.ResolveReference(&url.URL{Path: href}).String()
			c2l.Visit(absolute)
		}
	})
	// get c2 html div whose class is "ui buttons"
	c2l.OnHTML("div", func(e *colly.HTMLElement) {
		// only class is "ui buttons"
		if e.Attr("class") == "ui buttons" {
			// iter all the a tag
			e.ForEach("a", func(_ int, e *colly.HTMLElement) {
				// only class is "ui button" and text is "Raw"
				if e.Attr("class") == "ui button" && e.Text == "Raw" {
					href := e.Attr("href")
					base, _ := url.Parse(outerurl)
					tmpabs := base.ResolveReference(&url.URL{Path: href}).String()
					c3l.Visit(tmpabs)
				}
			})
		}
	})

	c3l.OnResponse(func(r *colly.Response) {
		res = string(r.Body)
	})
	// visit the url1
	c1l.Visit(URL)
	return updated, res
}
