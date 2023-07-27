package nook

import "golang.org/x/text/language"

// Name represents the name of a value in a specific language.
type Name struct {
	// Language specifies the language tag of the name.
	Language language.Tag

	// Value holds the actual name value in the specified language.
	Value string
}

// Ok returns a boolean indicating whether the Name information is complete and valid.
// It checks if the Value field has been set, meaning the name is not an empty string.
func (n Name) Ok() bool {
	return n.Value != ""
}
