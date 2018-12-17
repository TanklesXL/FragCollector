package manipulatefragranceitems

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

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
		currentCollection.MasterCollection[itemToAdd.Name] = itemToAdd
		updateCollection(currentCollection)
		fmt.Printf("%s has been added to your collection.\n", itemToAdd.Name)
		return true
	}
	fmt.Println("This fragrance is already in your collection")
	return false

}

func collectionContainsFragrance(collection FragranceCollection, name string) bool {
	_, exists := collection.MasterCollection[name]
	return exists
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
	sort.Slice(alphabeticalByBrand, func(i, j int) bool {
		return ((alphabeticalByBrand[i].House == alphabeticalByBrand[j].House && alphabeticalByBrand[i].Name < alphabeticalByBrand[j].Name) || alphabeticalByBrand[i].House < alphabeticalByBrand[j].House)
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

// ManualUpdate updates the rest of the json file when called, for example if someone were to manually add something to the file
func ManualUpdate() {
	collection := ReadInCollection(MASTER)
	updateCollection(collection)
}

// RemoveFromCollection takes the name of the fragrance to remove, removes it from the collection and then regenerated the json file
func RemoveFromCollection() {
	currentCollection := ReadInCollection(MASTER)
	if len(currentCollection.MasterCollection) == 0 {
		fmt.Println("\nYour collection is empty.")
		os.Exit(0)
	}
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

// ShowOptionsByBrandAndGetNumericInput displays the fragrances by brand and get the selection
func ShowOptionsByBrandAndGetNumericInput(collection FragranceCollection) int {
	max := len(collection.MasterCollection)
	var currentHouse string
	for i, f := range collection.FragrancesByHouse {
		if currentHouse == "" || f.House != currentHouse {
			currentHouse = f.House
			fmt.Printf("\n%s:\n", currentHouse)
		}
		fmt.Printf("\t%d-> %s\n", i+1, f.Name)
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
func updateCollection(collection FragranceCollection) {
	var newCollection FragranceCollection
	newCollection.MasterCollection = collection.MasterCollection
	newCollection.FragrancesByName = generateAlphabetical(newCollection)
	newCollection.FragrancesByHouse = generateAlphabeticalByBrand(newCollection)
	newCollection.Notes = generateByNote(newCollection)
	writeOutCollection(MASTER, newCollection)
}
