package manipulatefragranceitems

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// FragranceItem type contains all notes in a pyramid (if applicable) and flat list of scent notes
type FragranceItem struct {
	Name        string
	Designer    string
	ReleaseYear string
	FlatNotes   []string
	Pyramid     NotesPyramid
}

// NotesPyramid contains the pyramid
type NotesPyramid struct {
	TopNotes   []string
	HeartNotes []string
	BaseNotes  []string
}

// BuildFragranceItem receives a URL (from basenotes fragrance directory) and returns a FragranceItem with the corresponding information
func BuildFragranceItem(url string) FragranceItem {
	var fragrance FragranceItem

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	header := doc.Find(".fragranceheading").Text()

	//Get the name
	fragrance.Name = strings.TrimSpace(strings.Split(header, "(")[0])

	//Get the designer
	fragrance.Designer = strings.TrimSpace(strings.TrimPrefix(strings.Split(header, ")")[1], " by "))

	//Get the release year
	fragrance.ReleaseYear = strings.TrimSpace(strings.Split(strings.Split(header, "(")[1], ")")[0])

	//Get the notes
	notesText := doc.Find(".notespyramid.notespyramidb").Text()

	if strings.Contains(notesText, "Top Notes") || strings.Contains(notesText, "Heart Notes") || strings.Contains(notesText, "Base Notes") {
		fragrance.FlatNotes, fragrance.Pyramid = handlePyramidStructure(notesText)
	} else {
		fragrance.FlatNotes = handleFlatStructure(notesText)
	}

	return fragrance
}

func handleFlatStructure(text string) []string {
	notes := strings.Split(text, ",")
	notes = trimSlices(notes)
	return notes
}

func handlePyramidStructure(text string) ([]string, NotesPyramid) {

	var flatList []string
	var pyramid NotesPyramid
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
	flatList = append(topNotesSlice, heartNotesSlice...)
	flatList = append(flatList, baseNotesSlice...)
	pyramid.TopNotes = topNotesSlice
	pyramid.HeartNotes = heartNotesSlice
	pyramid.BaseNotes = baseNotesSlice

	return flatList, pyramid
}
func trimSlices(slice []string) []string {
	var output []string
	for _, s := range slice {
		output = append(output, strings.TrimSpace(s))
	}
	return output
}
