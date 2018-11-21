package manipulatefragranceitems

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	cmp "github.com/google/go-cmp/cmp"
)

// PATH is the location of the directory where the jsons are stored
const PATH string = "C:/Users/Robert/go/src/FragCollector/CollectionFiles/"

// MASTER is the master collecion filepath
const MASTER string = PATH + "Master.json"

// ALPHA is the filepath to the json where the collection is ordered alphabetically by name
const ALPHA string = PATH + "AlphabeticalName.json"

// BRAND is the filepath to the json where the collection is ordered alphabetically by fragrance house
const BRAND string = PATH + "AlphabeticalBrand.json"

// AddToCollection takes a url string and builds the corresponding fragrance item and adds it to the JSON
func AddToCollection(url string) bool {
	if _, err := os.Stat(MASTER); os.IsNotExist(err) {
		f, err := os.Create(MASTER)
		if err != nil {
			fmt.Println("UNABLE TO CREATE THE MASTER JSON FILE")
			os.Exit(0)
		}
		f.Close()
	}

	itemToAdd := BuildFragranceItem(url)

	currentCollection := readInCollection(MASTER)

	if !collectionContainsFragrance(currentCollection, itemToAdd) {
		currentCollection.Fragrances = append(currentCollection.Fragrances, itemToAdd)
		writeOutCollection(MASTER, currentCollection)
		Synchronise()
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

func readInCollection(filePath string) FragranceCollection {
	jsonFile, e := os.Open(filePath)
	if e != nil {
		fmt.Println("PROBLEM reading THE JSON FILE: " + filePath)
		os.Exit(0)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var currentCollection FragranceCollection
	json.Unmarshal(byteValue, &currentCollection)

	return currentCollection
}

func writeOutCollection(filePath string, currentCollection FragranceCollection) {
	json, _ := json.Marshal(currentCollection)
	err := ioutil.WriteFile(filePath, json, 0644)
	if err != nil {
		fmt.Println("PROBLEM WRITING TO THE JSON FILE")
		os.Exit(0)
	}
}

// Synchronise generates the JSON for the other views, built off of master.json
func Synchronise() {
	currentCollection := readInCollection(MASTER)
	generateAlphabetical(currentCollection)

	fmt.Println("***SYNCHRONISE EXECUTED***")

}
func generateAlphabetical(collection FragranceCollection) {
	//Alphabetical by name
	if _, err := os.Stat(ALPHA); os.IsNotExist(err) {
		f, err := os.Create(ALPHA)
		if err != nil {
			fmt.Println("Alphabetical by name: UNABLE TO CREATE THE JSON FILE")
			os.Exit(0)
		}
		defer f.Close()
	}

	// Sort by name
	sort.Slice(collection.Fragrances, func(i, j int) bool { return collection.Fragrances[i].Name < collection.Fragrances[j].Name })

	//Alphabetical by brand and by name
	writeOutCollection(ALPHA, collection)

	if _, err := os.Stat(BRAND); os.IsNotExist(err) {
		f, err := os.Create(BRAND)
		if err != nil {
			fmt.Println("Alphabetical by brand: UNABLE TO CREATE THE JSON FILE")
			os.Exit(0)
		}
		defer f.Close()
	}

	// Sort by brand
	sort.Slice(collection.Fragrances, func(i, j int) bool { return collection.Fragrances[i].House < collection.Fragrances[j].House })

	writeOutCollection(BRAND, collection)
}
