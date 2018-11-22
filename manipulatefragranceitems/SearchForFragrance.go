package manipulatefragranceitems

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type searchResult struct {
	Name string
	Info string
	URL  string
}

//SearchByName will do a search on BaseNotes for the fragrance name inputted, if no exact match is found, a list of options will be given
func SearchByName(name string) {
	searchString := strings.Replace(name, " ", "+", -1)
	url := "http://www.basenotes.net/fragrancedirectory/?search=" + searchString
	Search(url, name)
}

// Search goes through the search page on basenotes and tries to find a match, otherwise it gives the available options
func Search(url, nameToSearch string) {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("PROBLEM SEARCHING BASENOTES (are you connected to the internet?")
		os.Exit(0)
	}
	var searchResults []searchResult
	var matchResult *searchResult
	var added bool
	doc.Find("h3").Each(func(i int, h3 *goquery.Selection) {
		h3.Find("a").Each(func(j int, a *goquery.Selection) {
			searchURL, _ := a.Attr("href")
			text := strings.Split(a.Text(), "by")
			name := strings.TrimSpace(text[0])
			info := strings.TrimSpace(text[1])

			result := newSearchResult(name, info, searchURL)

			if strings.ToLower(name) == strings.ToLower(nameToSearch) {
				fmt.Print("A match was found! Is this what you were looking for?  ->  ")
				fmt.Printf("%s by %s\n> ", result.Name, result.Info)
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				if scanner.Err() != nil {
					panic(scanner.Err())
				}

				if (scanner.Text() == "yes" || scanner.Text() == "y") && result != nil && !added {
					matchResult = result
					added = true
					if AddToCollection(matchResult.URL) {
						fmt.Printf("%s has been added to your collection.\n", matchResult.Name)

					}
					return
				}
			}
			searchResults = append(searchResults, *result)
		})
	})

	if matchResult == nil {
		fmt.Println("I'm sorry, no match was found. Here are your options:")
		for _, r := range searchResults {
			fmt.Printf("%s	->	%s\n", r.Name, r.Info)
		}
	}
}

// SearchByHouse searches basenotes by fragrance house first and then by
func SearchByHouse(house, name string) {
	houseBaseURL := "http://www.basenotes.net/fragrancedirectory/?house="
	doc, err := goquery.NewDocument("http://www.basenotes.net/fragrancedirectory/")
	if err != nil {
		fmt.Println("PROBLEM SEARCHING BASENOTES (are you connected to the internet?")
		os.Exit(0)
	}
	doc.Find("select").Each(func(i int, selection *goquery.Selection) {
		if attribute, _ := selection.Attr("id"); attribute == "house" {
			selection.Find("option").Each(func(j int, option *goquery.Selection) {
				houseFromWeb := strings.Split(option.Text(), " (")[0]
				if strings.ToLower(houseFromWeb) == strings.ToLower(house) {
					id, _ := option.Attr("value")
					fmt.Println("\nSearching the house of " + houseFromWeb)
					Search(houseBaseURL+id, name)
				}
			})
		}
	})
}

func newSearchResult(name, info, url string) *searchResult {
	result := new(searchResult)
	result.Name = name
	result.Info = info
	result.URL = url
	return result
}
