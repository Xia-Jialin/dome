package dome01

import (
	"errors"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

//Chapter 章节
type Chapter struct {
	ID          int    `json:"id,omitempty"`
	ChapterName string `json:"chapter_name,omitempty"`
	ChapterURL  string `json:"chapter_url,omitempty"`
}

//GetChapterContent 获取章节内容
func (c Chapter) GetChapterContent() (string, error) {
	if c.ChapterURL == "" {
		log.Println("ChapterURL: nil")
		return "", errors.New("GetChapterContent: ChapterURL is empty")
	}
	res, e := http.Get(c.ChapterURL) //根据url获取该网页的内容  res
	if e != nil {
		log.Println("错误！")
		return "", errors.New("GetChapterContent: Content acquisition failed")
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// var str string
	// str = doc.Find("#content").Text()
	return doc.Find("#content").Text(), nil
}
