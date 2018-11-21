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
const PATH string = "./CollectionFiles/"

// MASTER is the master collecion filepath
const MASTER string = PATH + "Master.json"

// ALBPHA_N is the filepath to the json where the collection is ordered alphabetically by name
const ALPHA string = PATH + "Alphabetical.json"

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
		fmt.Println("PROBLEM OPENING THE JSON FILE")
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

// DisplayCollectionAlphabetical outputs the collection by fragrance in alphabetical order
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

// Synchronise generates the JSON for the other views, built off of master.json
func Synchronise() {
	currentCollection := readInCollection(MASTER)
	generateAlphabeticalByName(currentCollection)

	fmt.Println("***SYNCHRONISE EXECUTED***")

}
func generateAlphabeticalByName(collection FragranceCollection) {
	if _, err := os.Stat(ALPHA); os.IsNotExist(err) {
		f, err := os.Create(ALPHA)
		if err != nil {
			fmt.Println("Alphabetical by name: UNABLE TO CREATE THE JSON FILE")
			os.Exit(0)
		}
		defer f.Close()
	}

	//Sort the collection
	sort.Slice(collection.Fragrances, func(i, j int) bool { return collection.Fragrances[i].Name < collection.Fragrances[j].Name })

	writeOutCollection(PATH+"Alphabetical.json", collection)
}
