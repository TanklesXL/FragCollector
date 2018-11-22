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

// MASTER is the master collecion filepath, when an item is added, master is regenerated in alphabetical order
const MASTER string = PATH + "Master.json"

// BRAND is the filepath to the json where the collection is ordered alphabetically by fragrance house
const BRAND string = PATH + "AlphabeticalBrand.json"

// NOTES is the filepath to the json where the collection is broken down by scent notes
const NOTES string = PATH + "NoteBreakdown.json"

// AddToCollection takes a url string and builds the corresponding fragrance item and adds it to the JSON
func AddToCollection(url string) bool {

	if _, err := os.Stat(PATH); os.IsNotExist(err) {
		err := os.Mkdir(PATH, os.FileMode(0522))
		if err != nil {
			fmt.Println("UNABLE TO CREATE THE DIRECTORY")
			os.Exit(0)
		}
	}

	if _, err := os.Stat(MASTER); os.IsNotExist(err) {
		f, err := os.Create(MASTER)
		if err != nil {
			fmt.Println("UNABLE TO CREATE THE MASTER JSON FILE")
			os.Exit(0)
		}
		f.Close()
	}

	itemToAdd := BuildFragranceItem(url)

	currentCollection := ReadInCollection(MASTER)

	if !collectionContainsFragrance(currentCollection, itemToAdd) {
		currentCollection.Fragrances = append(currentCollection.Fragrances, itemToAdd)
		fmt.Printf("%s has been added to your collection.\n", itemToAdd.Name)
		generateAlphabetical(currentCollection)
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

// ReadInCollection reads a file and outputs a FragranceCollection
func ReadInCollection(filePath string) FragranceCollection {

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if filePath != MASTER {
			Synchronise()
		} else {
			fmt.Println("\nPROBLEM reading THE JSON FILE: " + filePath)
			fmt.Println("Master.json should automatically be created when a fragrance is added to your collection, did something happen?")
			fmt.Println("Potential causes:\n1) No fragrance has been added to the collection\n2) Master.json has been moved or deleted")
			os.Exit(0)
		}
	}

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
		fmt.Println("PROBLEM WRITING TO THE JSON FILE: " + filePath)
		os.Exit(0)
	}
}

// Synchronise generates the JSON for the other views, built off of master.json
func Synchronise() {
	currentCollection := ReadInCollection(MASTER)
	generateAlphabeticalByBrand(currentCollection)
	generateByNote(currentCollection)
}
func generateAlphabetical(collection FragranceCollection) {
	//Alphabetical by name
	if _, err := os.Stat(MASTER); os.IsNotExist(err) {
		f, err := os.Create(MASTER)
		if err != nil {
			fmt.Println("UNABLE TO CREATE THE JSON FILE: " + MASTER)
			os.Exit(0)
		}
		defer f.Close()
	}

	// Sort by name
	sort.Slice(collection.Fragrances, func(i, j int) bool { return collection.Fragrances[i].Name < collection.Fragrances[j].Name })

	//Alphabetical by brand and by name
	writeOutCollection(MASTER, collection)

}
func generateAlphabeticalByBrand(collection FragranceCollection) {
	if _, err := os.Stat(BRAND); os.IsNotExist(err) {
		f, err := os.Create(BRAND)
		if err != nil {
			fmt.Println("UNABLE TO CREATE THE JSON FILE: " + BRAND)
			os.Exit(0)
		}
		defer f.Close()
	}

	// Sort by brand, ssuming Master is already sorted by name, this should have the brands sorted by name as well
	sort.Slice(collection.Fragrances, func(i, j int) bool { return collection.Fragrances[i].House < collection.Fragrances[j].House })

	writeOutCollection(BRAND, collection)
}

func generateByNote(collection FragranceCollection) {

	if _, err := os.Stat(NOTES); os.IsNotExist(err) {
		file, err := os.Create(NOTES)
		if err != nil {
			fmt.Println("UNABLE TO CREATE THE JSON FILE: " + NOTES)
			os.Exit(0)
		}
		defer file.Close()
	}
	noteMap := make(map[string][]string)
	for _, frag := range collection.Fragrances {
		for _, note := range frag.FlatNotes {
			noteMap[note] = append(noteMap[note], frag.Name)
		}
	}
	writeOutNotesMap(NOTES, noteMap)
}
func writeOutNotesMap(filePath string, notesMap map[string][]string) {
	json, _ := json.Marshal(notesMap)
	err := ioutil.WriteFile(filePath, json, 0644)
	if err != nil {
		fmt.Println("PROBLEM WRITING TO THE JSON FILE: " + filePath)
		os.Exit(0)
	}
}

// ReadInNotesMap reads a file and outputs a map[string][]string
func ReadInNotesMap(filePath string) map[string][]string {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		Synchronise()
	}

	jsonFile, e := os.Open(filePath)
	if e != nil {
		fmt.Println("PROBLEM reading THE JSON FILE: " + filePath)
		os.Exit(0)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var notesMap map[string][]string
	json.Unmarshal(byteValue, &notesMap)

	return notesMap
}
