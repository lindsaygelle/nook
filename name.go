package nook

import "golang.org/x/text/language"

// Name is the name of a value in a specific language.
type Name struct {
	Language language.Tag
	Value    string
}

// Ok returns a boolean indicating whether the Name information is complete.
func (v Name) Ok() bool {
	return v.Value != ""
}
