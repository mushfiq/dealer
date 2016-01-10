package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

var wg sync.WaitGroup

var links = []string{
	"http://amazon.de",
	"http://cunda.de",
}

var keywords = []string{
	"sale",
	"angebote",
	"offer",
	"hife",
	"deals",
}

func keryWordExists(text string) bool {
	for _, keyword := range keywords{
		if strings.Contains(text, keyword) || strings.Contains(strings.Title(text), keyword) || strings.Contains(strings.ToLower(text), keyword) {
			return true
		} 
	}
	return false
}

func checkError(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}

func displayDetails(single *goquery.Selection) {
	text := strings.TrimSpace(single.Text())
	href, _ := single.Attr("href")
	length := utf8.RuneCountInString(text)
	if (length > 5) && keryWordExists(text) {
			fmt.Println("Link", single.Text(), "--->", href)
	}

}

func fetchAndDisplay(link string, wg *sync.WaitGroup) {
	fmt.Println(link)
	doc, err := goquery.NewDocument(link)
	checkError(err)

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
}
