package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // panther is the common reference for the Panther animal type.
    panther = "Panther"
)

var (
	// pantherNameAmericanEnglish represents the American English name for the Panther animal type.
	pantherNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    panther,
	}
)

var (
	// pantherName contains the names of the Panther animal type in different languages.
	pantherName = nook.Languages{
		language.AmericanEnglish: pantherNameAmericanEnglish,
	}
)

var (
	// Panther represents an panther animal in the Animal Crossing series.
	Panther = nook.Animal{
		Key:  nook.Key(panther),
		Name: pantherName,
	}
)
