package main

import (
	"dome/dome01"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println("1")
	url := "http://www.xbiquge.la/100000/15409/"
	res, e := http.Get(url) //根据url获取该网页的内容  res
	if e != nil {
		log.Println("错误！")
	}
	defer res.Body.Close()
	fmt.Println("2")
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	slice := make([]dome01.Chapter, 3)
	// Find the review items
	doc.Find("#list > dl > dd").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()           //#list > dl > dd > a
		title, _ := s.Find("a").Attr("href") //#list > dl > dd > a
		zjURL := "http://www.xbiquge.la" + title
		//fmt.Printf("Review %d: %s - %s\n", i, band, zjURL)
		var a dome01.Chapter
		a.ID = i
		a.ChapterName = band
		a.ChapterURL = zjURL
		slice = append(slice, a)
	})
	fmt.Printf("%v", slice)
}
