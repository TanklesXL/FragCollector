package manipulatefragranceitems

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// BuildFragranceItem receives a URL (from basenotes fragrance directory) and returns a FragranceItem with the corresponding information
func BuildFragranceItem(url string) FragranceItem {
	var fragrance FragranceItem

	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("PROBLEM ACCESSING THE BASENOTES PAGE: " + url)
		os.Exit(0)
	}

	header := doc.Find(".fragranceheading").Text()

	fragrance.BasicInfo.Name, fragrance.BasicInfo.House, fragrance.BasicInfo.ReleaseYear = getInfoFromHeader(header)

	//Get the notes
	notesText := doc.Find(".notespyramid.notespyramidb").Text()

	if strings.Contains(notesText, "Top Notes") || strings.Contains(notesText, "Heart Notes") || strings.Contains(notesText, "Base Notes") {
		fragrance.FlatNotes, fragrance.Pyramid = handlePyramidStructure(notesText)
	} else {
		fragrance.FlatNotes = handleFlatStructure(notesText)
	}

	return fragrance
}

func getInfoFromHeader(header string) (string, string, string) {

	var name, house, releaseYear string

	if strings.Contains(header, ") (") {
		name = strings.TrimSpace(strings.Split(header, ") (")[0] + ")")
		house = strings.TrimSpace(strings.TrimPrefix(strings.Split(header, ")")[2], " by "))
		releaseYear = strings.TrimSpace(strings.Split(strings.Split(header, ") (")[1], ")")[0])
	} else {
		name = strings.TrimSpace(strings.Split(header, "(")[0])
		house = strings.TrimSpace(strings.TrimPrefix(strings.Split(header, ")")[1], " by "))
		releaseYear = strings.TrimSpace(strings.Split(strings.Split(header, "(")[1], ")")[0])
	}
	return name, house, releaseYear
}

func handleFlatStructure(text string) []string {
	notes := strings.Split(text, ",")
	notes = trimSlices(notes)
	return notes
}

func handlePyramidStructure(text string) ([]string, notesPyramid) {

	text = strings.TrimSpace(text)
	topNotes := strings.TrimPrefix(strings.Split(text, "Heart Notes")[0], "Top Notes")
	topNotes = strings.TrimSpace(topNotes)

	heartNotes := strings.Split(strings.Split(text, "Heart Notes")[1], "Base notes")[0]
	heartNotes = strings.TrimSpace(heartNotes)

	baseNotes := strings.Split(text, "Base notes")[1]
	baseNotes = strings.TrimSpace(baseNotes)

	topNotesSlice := strings.Split(topNotes, ",")
	topNotesSlice = trimSlices(topNotesSlice)

	heartNotesSlice := strings.Split(heartNotes, ",")
	heartNotesSlice = trimSlices(heartNotesSlice)

	baseNotesSlice := strings.Split(baseNotes, ",")
	baseNotesSlice = trimSlices(baseNotesSlice)

	//create returned items
	flatList := append(topNotesSlice, heartNotesSlice...)
	flatList = append(flatList, baseNotesSlice...)
	pyramid := notesPyramid{topNotesSlice, heartNotesSlice, baseNotesSlice}
	return flatList, pyramid
}
func trimSlices(slice []string) []string {
	var output []string
	for _, s := range slice {
		output = append(output, strings.TrimSpace(s))
	}
	return output
}
