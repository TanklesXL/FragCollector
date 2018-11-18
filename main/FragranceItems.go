package main

// FragranceItem type contains all notes in a pyramid (if applicable) and flat list of scent notes
type FragranceItem struct {
	Title       string
	Designer    string
	ReleaseYear int
	FlatNotes   []string
	Pyramid     NotesPyramid
}

// NotesPyramid contains the pyramid
type NotesPyramid struct {
	TopNotes   []string
	HeartNotes []string
	BaseNotes  []string
}
