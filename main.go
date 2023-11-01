package main

import (
	"fmt"
	"time"

	"github.com/wonderstone/colly-proxy/AweiProxy"
)

func main() {

	// test getTodayDate
	month, date := getTodayDate()
	updated, result := AweiProxy.GetFile(AweiProxy.URL, month, date)
	fmt.Println(updated, result)
}

// get today date
func getTodayDate() (month string, day string) {
	t := time.Now()
	return t.Format("1"), t.Format("2")
}
