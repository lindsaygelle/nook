package nook

// Code represents the unique identifier assigned to an Animal Crossing character.
// These codes are based on data mined from various Animal Crossing games.
type Code struct {
	// Value holds the actual unique identifier code.
	Value string
}

// Ok returns a boolean indicating whether the Code information is complete and valid.
// It checks if the Value field has been set, meaning the code is not an empty string.
func (c Code) Ok() bool {
	return c.Value != ""
}
