package display

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	mfi "github.com/TanklesXL/FragCollector/manipulatefragranceitems"
)

var readInCollection = mfi.ReadInCollection

// CollectionAlphabetical outputs the collection in alphabetical order of fragrance names
func CollectionAlphabetical() {

	//Read the collection from the json file
	collection := readInCollection(mfi.MASTER)
	fmt.Println("\nHere is your collection in alphabetical order by name")
	fmt.Println("-------------------------------------------------------")
	for i, f := range collection.FragrancesByName {
		num := i + 1
		fmt.Printf("%d: %s by %s\n", num, f.Name, f.House)
	}
	fmt.Println("-------------------------------------------------------")
}

// CollectionAlphabeticalByBrand outputs the collection by fragrance in alphabetical order by fragrance houae and then by name
func CollectionAlphabeticalByBrand() {

	//Read the collection from the json file
	collection := readInCollection(mfi.MASTER)

	// output the collection in alphabetical order by brand and then by
	var currentHouse string

	fmt.Println("\nHere is your collection in alphabetical order by brand")
	fmt.Println("-------------------------------------------------------")

	for _, f := range collection.FragrancesByHouse {

		if currentHouse == "" || f.House != currentHouse {
			currentHouse = f.House
			fmt.Printf("\n%s:\n", currentHouse)
		}
		fmt.Printf("\t%s\n", f.Name)
	}
	fmt.Println("-------------------------------------------------------")
}

// CollectionNotes displays the collection broken down by its notes
func CollectionNotes() {
	notesMap := readInCollection(mfi.MASTER).Notes
	fmt.Println("\nHere is your collection broken down by scent notes")
	fmt.Println("-------------------------------------------------------")
	var noteList []string
	for note := range notesMap {
		noteList = append(noteList, note)
	}

	sort.Slice(noteList, func(i, j int) bool { return noteList[i] < noteList[j] })

	for _, note := range noteList {
		frags := notesMap[note]
		fmt.Printf("\n%s:\n", note)
		for _, frag := range frags {
			fmt.Printf("\t%s by %s \n", frag.Name, frag.House)
		}
	}
	fmt.Println("-------------------------------------------------------")
}

// SingleNote displays the fragrances listed for a single chosen scent note
func SingleNote() {
	notesMap := readInCollection(mfi.MASTER).Notes
	var noteList []string
	for note := range notesMap {
		noteList = append(noteList, note)
	}

	max := len(noteList)

	sort.Slice(noteList, func(i, j int) bool { return noteList[i] < noteList[j] })
	fmt.Println("\nPlease type in the number corresponding to your selection:")
	for i, n := range noteList {
		fmt.Printf("%d\t-> %s (%d)\n", i+1, n, len(notesMap[n]))
	}
	fmt.Printf("> ")

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

	index := inputAsInt - 1

	fmt.Println("\nFragrances containing " + noteList[index] + ":")
	for _, frag := range notesMap[noteList[index]] {
		fmt.Println(frag.Name + " by " + frag.House)
	}

	fmt.Println("-------------------------------------------------------")
}

// FragranceInfo outputs for the info for a specific fragrance
func FragranceInfo() {
	collection := readInCollection(mfi.MASTER)
	fmt.Println("\nWhich fragrance's info would you like")
	fmt.Println("-------------------------------------------------------")
	index := mfi.ShowOptionsByBrandAndGetNumericInput(collection) - 1
	frag := collection.MasterCollection[collection.FragrancesByHouse[index].Name]
	fmt.Printf("%s selected!\n", frag.BasicInfo.Name)
	printFragInfo(frag)
}

// FragranceListInfo outputs the info for a specific fragrance, chosen from a flat list
func FragranceListInfo() {
	collection := readInCollection(mfi.MASTER)
	fmt.Println("\nWhich fragrance's info would you like")
	fmt.Println("-------------------------------------------------------")
	index := mfi.ShowOptionsAndGetNumericInput(collection) - 1
	frag := collection.MasterCollection[collection.FragrancesByName[index].Name]
	fmt.Printf("%s selected!\n", frag.BasicInfo.Name)
	printFragInfo(frag)
}

func printFragInfo(frag mfi.FragranceItem) {
	fmt.Println("\n-------------------FRAGRANCE INFO-------------------")
	fmt.Printf("Name: %s\n", frag.BasicInfo.Name)
	fmt.Printf("Fragrance House: %s\n", frag.BasicInfo.House)
	fmt.Printf("Release Year: %s\n", frag.BasicInfo.ReleaseYear)
	fmt.Printf("\nAbout %s: \n\t%s\n", frag.BasicInfo.Name, frag.Blurb)

	if len(frag.Pyramid.TopNotes) != 0 {
		fmt.Println("Top Notes:")
		for _, n := range frag.Pyramid.TopNotes {
			fmt.Printf("\t%s\n", n)
		}
	}
	if len(frag.Pyramid.HeartNotes) != 0 {
		fmt.Println("\nHeart Notes:")
		for _, n := range frag.Pyramid.HeartNotes {
			fmt.Printf("\t%s\n", n)
		}
	}
	if len(frag.Pyramid.BaseNotes) != 0 {
		fmt.Println("\nBase Notes:")
		for _, n := range frag.Pyramid.BaseNotes {
			fmt.Printf("\t%s\n", n)
		}
	}
	if len(frag.Pyramid.TopNotes) == 0 && len(frag.Pyramid.HeartNotes) == 0 && len(frag.Pyramid.BaseNotes) == 0 {
		fmt.Println("\nScent Notes")
		for _, n := range frag.FlatNotes {
			fmt.Printf("\t%s\n", n)
		}
	}
	fmt.Printf("\nURL: %s\n", frag.Link)
}
