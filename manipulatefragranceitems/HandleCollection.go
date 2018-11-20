package manipulatefragranceitems

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	cmp "github.com/google/go-cmp/cmp"
)

// AddToCollection takes a url string and builds the corresponding fragrance item and adds it to the JSON
func AddToCollection(url string) bool {
	if _, err := os.Stat("./Collection.json"); os.IsNotExist(err) {
		f, err := os.Create("./Collection.json")
		if err != nil {
			fmt.Println("UNABLE TO CREATE THE JSON FILE")
			os.Exit(0)
		}
		defer f.Close()
	}

	itemToAdd := BuildFragranceItem(url)

	currentCollection := readInCollection("./Collection.json")

	if !collectionContainsFragrance(currentCollection, itemToAdd) {
		currentCollection.Fragrances = append(currentCollection.Fragrances, itemToAdd)
		writeOutCollection("./Collection.json", currentCollection)
		return true
	}
	return false
}

func collectionContainsFragrance(collection FragranceCollection, fragrance FragranceItem) bool {
	for _, v := range collection.Fragrances {
		if cmp.Equal(v, fragrance) {
			fmt.Println("This fragrance is already in your collection")
			return true
		}
	}
	return false
}

func readInCollection(path string) FragranceCollection {
	jsonFile, e := os.Open(path)
	if e != nil {
		fmt.Println("PROBLEM OPENING THE JSON FILE")
		os.Exit(0)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var currentCollection FragranceCollection
	json.Unmarshal(byteValue, &currentCollection)

	return currentCollection
}

func writeOutCollection(path string, currentCollection FragranceCollection) {
	json, _ := json.Marshal(currentCollection)
	err := ioutil.WriteFile(path, json, 0644)
	if err != nil {
		fmt.Println("PROBLEM WRITING TO THE JSON FILE")
		os.Exit(0)
	}
}

// DisplayCollectionAlphabetical outputs the collection by fragrance in alphabetical order
func DisplayCollectionAlphabetical() {

	//Read the collection from the json file
	collection := readInCollection("./Collection.json")

	//Sort the collection

	// output the collection in alphabetical order
	fmt.Println()
	fmt.Println("Here is your collection in alphabetical order by name")
	fmt.Println("-----------------------------------------------------")
	for i, f := range collection.Fragrances {
		num := i + 1
		fmt.Printf("%d: %s by %s\n", num, f.Name, f.House)
	}
}
