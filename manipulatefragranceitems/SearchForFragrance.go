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

	for canClickNext := true; canClickNext; {
		canClickNext = false
		doc.Find("h3").Each(func(i int, h3 *goquery.Selection) {
			h3.Find("a").Each(func(j int, a *goquery.Selection) {
				searchURL, _ := a.Attr("href")
				text := strings.Split(a.Text(), "by")
				name := strings.TrimSpace(text[0])
				info := strings.TrimSpace(text[1])

				result := newSearchResult(name, info, searchURL)

				if shrinkString(strings.ToLower(name)) == shrinkString(strings.ToLower(nameToSearch)) {
					fmt.Print("A match was found! Is this what you were looking for?  ->  ")
					fmt.Printf("%s by %s\n> ", result.Name, result.Info)
					scanner := bufio.NewScanner(os.Stdin)
					scanner.Scan()
					if scanner.Err() != nil {
						panic(scanner.Err())
					}

					if (strings.ToLower(scanner.Text()) == "yes" || strings.ToLower(scanner.Text()) == "y") && result != nil {
						matchResult = result
						if AddToCollection(matchResult.URL) {
							fmt.Printf("%s has been added to your collection.\n", matchResult.Name)
							return
						}
					}
				}
				searchResults = append(searchResults, *result)
			})
		})

		var link string
		doc.Find("a").Each(func(i int, sel *goquery.Selection) {
			if title, _ := sel.Attr("title"); title == "Next Page" {
				link, _ = sel.Attr("href")
				canClickNext = true
			}
		})
		if link != "" {
			doc, err = goquery.NewDocument(link)
			if err != nil {
				fmt.Println("PROBLEM GOING TO THE NEXT PAGE")
				os.Exit(0)
			}
		}
	}

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
	var houses []string
	houseFound := false
	selection := doc.Find("select").First()
	if attribute, _ := selection.Attr("id"); attribute == "house" {
		selection.Find("option").Each(func(j int, option *goquery.Selection) {
			if !houseFound {
				houseFromWeb := strings.Split(option.Text(), " (")[0]
				if shrinkString(strings.ToLower(houseFromWeb)) == shrinkString(strings.ToLower(house)) {
					id, _ := option.Attr("value")
					fmt.Println("\nSearching the house of " + houseFromWeb + "...")
					Search(houseBaseURL+id, name)
					houseFound = true
					return
				}
				if !houseFound {
					houses = append(houses, houseFromWeb)
				}
			}
		})
	}
	if !houseFound {
		fmt.Println("That house was not found")
	}
}

func newSearchResult(name, info, url string) *searchResult {
	result := new(searchResult)
	result.Name = name
	result.Info = info
	result.URL = url
	return result
}

func shrinkString(s string) string {
	return strings.Replace(s, " ", "", -1)
}
