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
	jsonFile, e := os.Open("./Collection.json")
	if e != nil {
		panic(e)
	}
	itemToAdd := BuildFragranceItem(url)

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var currentCollection FragranceCollection
	json.Unmarshal(byteValue, &currentCollection)

	if !collectionContainsFragrance(currentCollection, itemToAdd) {

		var collection FragranceCollection

		collection.Fragrances = append(currentCollection.Fragrances, itemToAdd)

		json, _ := json.Marshal(collection)
		err := ioutil.WriteFile("./Collection.json", json, 0644)
		if err != nil {
			panic(err)
		}
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
