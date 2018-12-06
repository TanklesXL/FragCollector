package manipulatefragranceitems

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"runtime"
)

// PATH is the location of the directory where the jsons are stored
var PATH string

// MASTER is the master collecion filepath, when an item is added, master is regenerated in alphabetical order
var MASTER string

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

// SetPath determines what the path should be based on the OS
func SetPath() {
	userInfo, err := user.Current()
	if err != nil {
		fmt.Println("Unable to get your user info")
	}
	if runtime.GOOS == "windows" {
		PATH = userInfo.HomeDir + "\\Documents\\FragCollector"
		MASTER = PATH + "\\Master.json"
	} else {
		PATH = userInfo.HomeDir + "/FragCollector"
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

func writeOutCollection(filePath string, currentCollection FragranceCollection) {
	json, _ := json.Marshal(currentCollection)
	err := ioutil.WriteFile(filePath, json, 0644)
	if err != nil {
		fmt.Println("PROBLEM WRITING TO THE JSON FILE: " + filePath)
		os.Exit(0)
	}
}
