package nook

// Code is the unique identifier assigned to an Animal Crossing character.
// Codes are based on the data mined from the various Animal Crossing games.
type Code struct {
	Value string
}

// Ok returns a boolean indicating whether the Code information is complete.
func (v Code) Ok() bool {
	return v.Value != ""
}
