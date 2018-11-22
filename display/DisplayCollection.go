package display

import "fmt"
import mfi "FragCollector/manipulatefragranceitems"

var readInCollection = mfi.ReadInCollection
var readInNotesMap = mfi.ReadInNotesMap

// DisplayCollectionAlphabetical outputs the collection in alphabetical order of fragrance names
func DisplayCollectionAlphabetical() {

	//Read the collection from the json file
	collection := readInCollection(mfi.MASTER)

	// output the collection in alphabetical order
	fmt.Println("\nHere is your collection in alphabetical order by name")
	fmt.Println("-------------------------------------------------------")
	for i, f := range collection.Fragrances {
		num := i + 1
		fmt.Printf("%d: %s by %s\n", num, f.Name, f.House)
	}
	fmt.Println("\n-------------------------------------------------------")
}

// DisplayCollectionAlphabeticalByBrand outputs the collection by fragrance in alphabetical order by fragrance houae and then by name
func DisplayCollectionAlphabeticalByBrand() {

	//Read the collection from the json file
	collection := readInCollection(mfi.BRAND)

	// output the collection in alphabetical order by brand and then by
	var currentHouse string

	fmt.Println("\nHere is your collection in alphabetical order by brand")
	fmt.Println("-------------------------------------------------------")

	numInGroup := 1
	for _, f := range collection.Fragrances {

		if currentHouse == "" || f.House != currentHouse {
			currentHouse = f.House
			numInGroup = 1
			fmt.Printf("\n%s:\n", currentHouse)
		}
		fmt.Printf("\t%d -> %s\n", numInGroup, f.Name)
		numInGroup++
	}
	fmt.Println("\n-------------------------------------------------------")
}

// DisplayCollectionNotes displays the collection broken down by its notes
func DisplayCollectionNotes() {
	notesMap := readInNotesMap(mfi.NOTES)
	fmt.Println("\nHere is your collection broken down by scent notes")
	fmt.Println("-------------------------------------------------------")
	for note, frags := range notesMap {
		fmt.Printf("\n%s:\n", note)
		for i, frag := range frags {
			fmt.Printf("\t%d -> %s\n", i+1, frag)
		}
	}

	fmt.Println("\n-------------------------------------------------------")
}
