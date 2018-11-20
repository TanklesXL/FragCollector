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
			panic(err)
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
		panic(e)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var currentCollection FragranceCollection
	json.Unmarshal(byteValue, &currentCollection)

	return currentCollection
}

func writeOutCollection(path string ,currentCollection FragranceCollection) {
	json, _ := json.Marshal(currentCollection)
	err := ioutil.WriteFile(path, json, 0644)
	if err != nil {
		panic(err)
	}
	
}
