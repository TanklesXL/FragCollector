package display

import "fmt"
import mfi "FragCollector/manipulatefragranceitems"

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
	fmt.Println("\n-------------------------------------------------------")
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
		fmt.Printf("\t-> %s\n", f.Name)
	}
	fmt.Println("\n-------------------------------------------------------")
}

// CollectionNotes displays the collection broken down by its notes
func CollectionNotes() {
	notesMap := readInCollection(mfi.MASTER).Notes
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
