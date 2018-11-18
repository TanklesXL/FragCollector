package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	getNotesPyramid("http://www.basenotes.net/ID10211632.html")

}
func getNotesPyramid(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	notesPyramid := doc.Find(".notespyramid.notespyramidb")

	notesPyramid.Each(func(index int, item *goquery.Selection) {
		if Contains(item.Text(), "Top Notes") {

		}
	})

	//case 1: notes pyramid is actually defined as a pyramid

	//case 2: pyramid is flat list
}
