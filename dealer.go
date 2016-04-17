package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode/utf8"
	"github.com/PuerkitoBio/goquery"
	"github.com/mushfiq/dealer/utils"
)

var wg sync.WaitGroup

var links = []string{
	"http://amazon.de",
	"http://cunda.de",
	"http://www.hm.com/de/",
}

var keywords = utils.GetConfig("keywords")

var emailConfig = utils.GetConfig("email")

func keyWordExists(text string) bool {
	keywords := keywords.Get("keywords").MustArray()
	for _, keyword := range keywords {
		if strings.Contains(text, keyword.(string)) ||
			strings.Contains(strings.Title(text), keyword.(string)) ||
			strings.Contains(strings.ToLower(text), keyword.(string)) {
			return true
		}
	}
	return false
}

func displayDetails(single *goquery.Selection) {
	text := strings.TrimSpace(single.Text())
	href, _ := single.Attr("href")
	length := utf8.RuneCountInString(text)
	if ((length > 5) && keyWordExists(text)) || ((length > 5) && keyWordExists(href)) {
		fmt.Println("Link", single.Text(), "--->", href)
	}

}

func fetchAndDisplay(link string, wg *sync.WaitGroup) {
	fmt.Println(link)
	doc, err := goquery.NewDocument(link)
	utils.CheckError(err)

	sel := doc.Find("a")
	for i := range sel.Nodes {
		single := sel.Eq(i)
		displayDetails(single)
	}
	wg.Done()
}

func main() {
	for _, link := range links {
		wg.Add(1)
		go fetchAndDisplay(link, &wg)
	}
	wg.Wait()
	
	// disabling email sending feature
	// utils.SendEmail("Message", emailConfig)
}
