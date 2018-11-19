package manipulatefragranceitems

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Search will do a search on BaseNotes for the fragrance name inputted, if no exact match is found, a list of options will be given
func Search(name string) {
	searchString := strings.Replace(name, " ", "+", -1)
	url := "http://www.basenotes.net/fragrancedirectory/?search=" + searchString

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	items:=doc.Find("h3").Each(func(index int, s *goquery.Selection){fmt.println(s.Text())})
	
		
	}
}
