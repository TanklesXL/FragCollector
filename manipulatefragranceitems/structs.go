package manipulatefragranceitems

// FragranceCollection is used to hold a collection of FragranceItems
type FragranceCollection struct {
	MasterCollection  map[string]FragranceItem
	FragrancesByName  []BasicInfo
	FragrancesByHouse []BasicInfo
	Notes             map[string][]BasicInfo
}

// FragranceItem type contains fragrance info and all notes in a pyramid (if applicable) and flat list of scent notes
type FragranceItem struct {
	BasicInfo BasicInfo
	FlatNotes []string
	Pyramid   notesPyramid
}

// BasicInfo of a fragrance Item will contain only it's name, house and release year
type BasicInfo struct {
	Name        string
	House       string
	ReleaseYear string
}

// notesPyramid contains the pyramid
type notesPyramid struct {
	TopNotes   []string
	HeartNotes []string
	BaseNotes  []string
}
