package manipulatefragranceitems

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"sort"
	"strconv"
)

// PATH is the location of the directory where the jsons are stored
var PATH string

// MASTER is the master collecion filepath, when an item is added, master is regenerated in alphabetical order
var MASTER string

// AddToCollection takes a url string and builds the corresponding fragrance item and adds it to the JSON
func AddToCollection(url string) bool {

	//ensure that the directory and master both exist, otherwise make them
	makeMaster()

	itemToAdd := BuildFragranceItem(url)

	currentCollection := ReadInCollection(MASTER)

	if !collectionContainsFragrance(currentCollection, itemToAdd.BasicInfo.Name) {
		if len(currentCollection.MasterCollection) == 0 {
			currentCollection.MasterCollection = make(map[string]FragranceItem)
		}
		currentCollection.MasterCollection[itemToAdd.BasicInfo.Name] = itemToAdd
		updateCollection(currentCollection)
		fmt.Printf("%s has been added to your collection.\n", itemToAdd.BasicInfo.Name)
		return true
	}
	fmt.Println("This fragrance is already in your collection")
	return false

}

// RemoveFromCollection takes the name of the fragrance to remove, removes it from the collection and then regenerated the json file
func RemoveFromCollection() {
	currentCollection := ReadInCollection(MASTER)
	fmt.Println("Please type in the number of the fragrance you'd like to remove:")
	inputAsInt := ShowOptionsAndGetNumericInput(currentCollection)
	inputIndex := inputAsInt - 1
	keyToRemove := currentCollection.FragrancesByName[inputIndex].Name
	if collectionContainsFragrance(currentCollection, keyToRemove) {
		delete(currentCollection.MasterCollection, keyToRemove)
		updateCollection(currentCollection)
		fmt.Printf("%s has been removed from your collection.\n", keyToRemove)
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

// ShowOptionsAndGetNumericInput displays the fragrances in the collection and gets the users selection
func ShowOptionsAndGetNumericInput(collection FragranceCollection) int {
	max := len(collection.MasterCollection)
	for i, v := range collection.FragrancesByName {
		fmt.Printf("%d -> %s by %s\n", i+1, v.Name, v.House)
	}
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() != nil {
		fmt.Println("INVALID INPUT")
		os.Exit(0)
	}
	inputAsInt, err := strconv.Atoi(scanner.Text())
	if err != nil || inputAsInt <= 0 || inputAsInt > max {
		fmt.Println("INVALID INPUT")
		os.Exit(0)
	}
	return inputAsInt
}

// ManualUpdate updates the rest of the json file when called, for example if someone were to manually add something to the file
func ManualUpdate() {
	collection := ReadInCollection(MASTER)
	updateCollection(collection)
}

func updateCollection(collection FragranceCollection) {
	var newCollection FragranceCollection
	newCollection.MasterCollection = collection.MasterCollection
	newCollection.FragrancesByName = generateAlphabetical(newCollection)
	newCollection.FragrancesByHouse = generateAlphabeticalByBrand(newCollection)
	newCollection.Notes = generateByNote(newCollection)
	writeOutCollection(MASTER, newCollection)
}

// SetPath determines what the path should be based on the OS
func SetPath() {
	if runtime.GOOS == "windows" {
		PATH = "C:\\Users\\Robert\\Documents\\FragCollector"
		MASTER = PATH + "\\Master.json"
	} else {
		PATH = "$HOME/FragCollector"
		MASTER = PATH + "/Master.json"
	}
}

func makeMaster() {
	if _, err := os.Stat(PATH); os.IsNotExist(err) {
		err := os.Mkdir(PATH, os.FileMode(0522))
		if err != nil {
			fmt.Println("UNABLE TO CREATE THE DIRECTORY: " + err.Error())
			os.Exit(0)
		}
	}
	if _, err := os.Stat(MASTER); os.IsNotExist(err) {
		f, err := os.Create(MASTER)
		if err != nil {
			fmt.Println("UNABLE TO CREATE THE MASTER JSON FILE: " + err.Error())
			os.Exit(0)
		}
		f.Close()
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
	var alphabetical []BasicInfo
	for _, v := range collection.MasterCollection {
		alphabetical = append(alphabetical, v.BasicInfo)
	}
	// Sort by name
	sort.Slice(alphabetical, func(i, j int) bool { return alphabetical[i].Name < alphabetical[j].Name })
	return alphabetical
}

func generateAlphabeticalByBrand(collection FragranceCollection) []BasicInfo {
	var alphabeticalByBrand []BasicInfo
	for _, v := range collection.MasterCollection {
		alphabeticalByBrand = append(alphabeticalByBrand, v.BasicInfo)
	}
	sort.Slice(alphabeticalByBrand, func(i, j int) bool { return alphabeticalByBrand[i].Name < alphabeticalByBrand[j].Name })
	sort.Slice(alphabeticalByBrand, func(i, j int) bool { return alphabeticalByBrand[i].House < alphabeticalByBrand[j].House })
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
