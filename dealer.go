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

var websites = utils.GetConfig("urls")

var keywords = utils.GetConfig("keywords")

var emailConfig = utils.GetConfig("email")

var products = utils.GetConfig("products")

func wordExists(text string, searchtype string) bool {
	words := products.Get("products").MustArray()
	if (searchtype=="keywords"){
		words = keywords.Get("keywords").MustArray()
	}
	
	for _, word := range words {
		if strings.Contains(text, word.(string)) ||
			strings.Contains(strings.Title(text), word.(string)) ||
			strings.Contains(strings.ToLower(text), word.(string)) {
			return true
		}
	}
	return false
}

func displayDetails(single *goquery.Selection) {
	text := strings.TrimSpace(single.Text())
	href, _ := single.Attr("href")
	length := utf8.RuneCountInString(text)
	if ((length > 5) && wordExists(text, "keywords")) || ((length > 5) && wordExists(href, "keywords")) {
		if wordExists(text, "products"){
			fmt.Println("Link", single.Text(), "--->", href)
		}
	}

}

func fetchAndDisplay(url string, wg *sync.WaitGroup) {
	doc, err := goquery.NewDocument(url)
	utils.CheckError(err)

	sel := doc.Find("a")
	for i := range sel.Nodes {
		single := sel.Eq(i)
		displayDetails(single)
	}
	wg.Done()
}

func main() {
	urls := websites.Get("urls").MustArray()
	for _, url := range urls {
		wg.Add(1)
		go fetchAndDisplay(url.(string), &wg)
	}
	wg.Wait()
	
	// disabling email sending feature
	// utils.SendEmail("Message", emailConfig)
}
