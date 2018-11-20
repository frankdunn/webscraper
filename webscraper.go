package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func metaScrape() {
	f, err := os.Open("C:/Users/frank/Documents/test.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, _ := goquery.NewDocumentFromReader(f)
	//doc, err := goquery.NewDocument("C:/Users/frank/Documents/test.html")

	if err != nil {
		log.Fatal(err)
	}
	var pageTitle string
	pageTitle = doc.Find("title").Contents().Text()
	fmt.Printf("Page Title: '%s'\n", pageTitle)
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {

		newtext := s.Find("td").Text()
		//title := s.Find("i").Text()

		fmt.Printf("result : %d: %s \n", i, newtext)
		if strings.Contains(newtext, "HF2") {
			fmt.Println("***********found it *****************")
		}
	})

}

func main() {
	metaScrape()
}
