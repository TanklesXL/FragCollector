package manipulatefragranceitems

import "fmt"

// DisplayCollectionAlphabetical outputs the collection in alphabetical order of fragrance names
func DisplayCollectionAlphabetical() {

	//Read the collection from the json file
	collection := readInCollection(ALPHA)

	// output the collection in alphabetical order
	fmt.Println("\nHere is your collection in alphabetical order by name")
	fmt.Println("-----------------------------------------------------")
	for i, f := range collection.Fragrances {
		num := i + 1
		fmt.Printf("%d: %s by %s\n", num, f.Name, f.House)
	}
}

// DisplayCollectionAlphabeticalByBrand outputs the collection by fragrance in alphabetical order by fragrance houae and then by name
func DisplayCollectionAlphabeticalByBrand() {

	//Read the collection from the json file
	collection := readInCollection(BRAND)

	// output the collection in alphabetical order by brand and then by
	var currentHouse string

	fmt.Println("\nHere is your collection in alphabetical order by brand")
	fmt.Println("-----------------------------------------------------")

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
}
