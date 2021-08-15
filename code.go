package nook

// Code is the unique identifier assigned to an Animal Crossing character.
// Codes are based on the data mined from the various Animal Crossing games.
type Code struct {
	Value string
}

func (v Code) Ok() bool {
	return v.Value != ""
}
