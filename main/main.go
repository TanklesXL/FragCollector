package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	_ = buildFragranceNoteStructure("http://www.basenotes.net/ID10211632.html")

}

func buildFragranceNoteStructure(url string) FragranceItem {
	var fragrance FragranceItem

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	//Get the name

	//Get the designer

	//Get the release year

	//Get the notes
	notesText := doc.Find(".notespyramid.notespyramidb").Text()

	if strings.Contains(notesText, "Top Notes") || strings.Contains(notesText, "Heart Notes") || strings.Contains(notesText, "Base Notes") {
		_, _ = handlePyramidStructure(notesText)
	} else {
		fragrance.FlatNotes = handleFlatStructure(notesText)
	}
	return fragrance
}

func handleFlatStructure(text string) []string {
	notes := strings.Split(text, ",")
	fmt.Println("FLAT")
	for i, n := range notes {
		n = strings.TrimSpace(n)
		fmt.Printf("%d -> %s\n", i, n)
	}
	return notes
}

func handlePyramidStructure(text string) ([]string, NotesPyramid) {
	var flatList []string
	var pyramid NotesPyramid
	text = strings.TrimSpace(text)
	topNotes := strings.Trim(strings.Split(text, "Heart Notes")[0], "Top Notes")
	topNotes = strings.TrimSpace(topNotes)

	heartNotes := strings.Split(strings.Split(text, "Heart Notes")[1], "Base notes")[0]
	heartNotes = strings.TrimSpace(heartNotes)

	baseNotes := strings.Split(text, "Base notes")[1]
	baseNotes = strings.TrimSpace(baseNotes)

	fmt.Println(topNotes)
	fmt.Println(heartNotes)
	fmt.Println(baseNotes)

	pyramid.TopNotes = strings.Split(topNotes, ",")
	pyramid.HeartNotes = strings.Split(heartNotes, ",")
	pyramid.BaseNotes = strings.Split(baseNotes, ",")

	return flatList, pyramid
}
