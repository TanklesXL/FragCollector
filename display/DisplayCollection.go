package display

import (
	mfi "FragCollector/manipulatefragranceitems"
	"fmt"
	"sort"
)

var readInCollection = mfi.ReadInCollection

// CollectionAlphabetical outputs the collection in alphabetical order of fragrance names
func CollectionAlphabetical() {

	//Read the collection from the json file
	collection := readInCollection(mfi.MASTER)

	// output the collection in alphabetical order
	fmt.Println("\nHere is your collection in alphabetical order by name")
	fmt.Println("-------------------------------------------------------")
	for i, f := range collection.FragrancesByName {
		num := i + 1
		fmt.Printf("%d: %s by %s\n", num, f.Name, f.House)
	}
	fmt.Println("-------------------------------------------------------")
}

// CollectionAlphabeticalByBrand outputs the collection by fragrance in alphabetical order by fragrance houae and then by name
func CollectionAlphabeticalByBrand() {

	//Read the collection from the json file
	collection := readInCollection(mfi.MASTER)

	// output the collection in alphabetical order by brand and then by
	var currentHouse string

	fmt.Println("\nHere is your collection in alphabetical order by brand")
	fmt.Println("-------------------------------------------------------")

	for _, f := range collection.FragrancesByName {

		if currentHouse == "" || f.House != currentHouse {
			currentHouse = f.House
			fmt.Printf("\n%s:\n", currentHouse)
		}
		fmt.Printf("\t%s\n", f.Name)
	}
	fmt.Println("-------------------------------------------------------")
}

// CollectionNotes displays the collection broken down by its notes
func CollectionNotes() {
	notesMap := readInCollection(mfi.MASTER).Notes
	fmt.Println("\nHere is your collection broken down by scent notes")
	fmt.Println("-------------------------------------------------------")
	var noteList []string
	for note := range notesMap {
		noteList = append(noteList, note)
	}

	sort.Slice(noteList, func(i, j int) bool { return noteList[i] < noteList[j] })

	for _, note := range noteList {
		frags := notesMap[note]
		fmt.Printf("\n%s:\n", note)
		for _, frag := range frags {
			fmt.Printf("\t%s by %s \n", frag.Name, frag.House)
		}
	}
	fmt.Println("-------------------------------------------------------")
}
