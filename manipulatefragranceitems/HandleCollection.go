package manipulatefragranceitems

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// PATH is the location of the directory where the jsons are stored
const PATH string = "C:/Users/Robert/go/src/FragCollector/"

// MASTER is the master collecion filepath, when an item is added, master is regenerated in alphabetical order
const MASTER string = PATH + "CollectionFile/Master.json"

/*
EXPORTED FUNCTIONS
These are used by either a command or a display function
*/

// AddToCollection takes a url string and builds the corresponding fragrance item and adds it to the JSON
func AddToCollection(url string) bool {

	//ensure that the directory and master both exist, otherwise make them
	makeDir()
	makeMaster()

	itemToAdd := BuildFragranceItem(url)

	currentCollection := ReadInCollection(MASTER)

	if !collectionContainsFragrance(currentCollection, itemToAdd.BasicInfo.Name) {
		if len(currentCollection.MasterCollection) == 0 {
			currentCollection.MasterCollection = make(map[string]FragranceItem)
		}
		currentCollection.MasterCollection[itemToAdd.BasicInfo.Name] = itemToAdd
		fmt.Printf("%s has been added to your collection.\n", itemToAdd.BasicInfo.Name)
		currentCollection.FragrancesByName = generateAlphabetical(currentCollection)
		currentCollection.FragrancesByHouse = generateAlphabeticalByBrand(currentCollection)
		currentCollection.Notes = generateByNote(currentCollection)
		writeOutCollection(MASTER, currentCollection)
		return true
	} else {
		fmt.Println("This fragrance is already in your collection")
		return false
	}
}

// RemoveFromCollection takes the name of the fragrance to remove, removes it from the collection and then regenerated the json file
func RemoveFromCollection(name string) {
	currentCollection := ReadInCollection(MASTER)

	if collectionContainsFragrance(currentCollection, name) {
		if len(currentCollection.MasterCollection) == 0 {
			currentCollection.MasterCollection = make(map[string]FragranceItem)
		}
		delete(currentCollection.MasterCollection, name)
		fmt.Printf("%s has been removed from your collection.\n", name)
		currentCollection.FragrancesByName = generateAlphabetical(currentCollection)
		currentCollection.FragrancesByHouse = generateAlphabeticalByBrand(currentCollection)
		currentCollection.Notes = generateByNote(currentCollection)
		writeOutCollection(MASTER, currentCollection)
	} else {
		fmt.Println("This fragrance is not in your collection")
	}

}

// ReadInCollection reads a file and outputs a FragranceCollection
func ReadInCollection(filePath string) FragranceCollection {

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("\nPROBLEM reading THE JSON FILE: " + filePath)
		fmt.Println("Master.json should automatically be created when a fragrance is added to your collection, did something happen?")
		fmt.Println("Potential causes:\n1) No fragrance has been added to the collection\n2) Master.json has been moved or deleted")
		os.Exit(0)

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

func makeDir() {
	if _, err := os.Stat(PATH); os.IsNotExist(err) {
		err := os.Mkdir(PATH, os.FileMode(0522))
		if err != nil {
			fmt.Println("UNABLE TO CREATE THE DIRECTORY")
			os.Exit(0)
		}
	}
}

func makeMaster() {
	if _, err := os.Stat(MASTER); os.IsNotExist(err) {
		f, err := os.Create(MASTER)
		if err != nil {
			fmt.Println("UNABLE TO CREATE THE MASTER JSON FILE")
			os.Exit(0)
		}
		defer f.Close()
		writeOutCollection(MASTER, *new(FragranceCollection))
	}

}

func collectionContainsFragrance(collection FragranceCollection, name string) bool {
	_, exists := collection.MasterCollection[name]
	return exists
}

func writeOutCollection(filePath string, currentCollection FragranceCollection) {
	json, _ := json.Marshal(currentCollection)
	err := ioutil.WriteFile(filePath, json, 0644)
	if err != nil {
		fmt.Println("PROBLEM WRITING TO THE JSON FILE: " + filePath)
		os.Exit(0)
	}
}

func generateAlphabetical(collection FragranceCollection) []BasicInfo {

	var alphabeticalSlice []BasicInfo

	for _, value := range collection.MasterCollection {
		alphabeticalSlice = append(alphabeticalSlice, value.BasicInfo)
	}
	// Sort by name
	sort.Slice(alphabeticalSlice, func(i, j int) bool { return alphabeticalSlice[i].Name < alphabeticalSlice[j].Name })

	return alphabeticalSlice
}

func generateAlphabeticalByBrand(collection FragranceCollection) []BasicInfo {
	alphabeticalByBrand := collection.FragrancesByName
	// Sort the list of items already sorted by name by their fragrance house, this results in the brands being sorted by name as well
	sort.Slice(alphabeticalByBrand, func(i, j int) bool {
		return alphabeticalByBrand[i].House < alphabeticalByBrand[j].House
	})

	return alphabeticalByBrand
}

func generateByNote(collection FragranceCollection) map[string][]BasicInfo {

	noteMap := make(map[string][]BasicInfo)
	for _, frag := range collection.MasterCollection {
		for _, note := range frag.FlatNotes {
			noteMap[note] = append(noteMap[note], frag.BasicInfo)
		}
	}
	return noteMap
}
