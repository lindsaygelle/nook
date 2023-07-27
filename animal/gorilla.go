package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // gorilla is the common reference for the Gorilla animal type.
    gorilla = "Gorilla"
)

var (
	// gorillaNameAmericanEnglish represents the American English name for the Gorilla animal type.
	gorillaNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    gorilla,
	}
)

var (
	// gorillaName contains the names of the Gorilla animal type in different languages.
	gorillaName = nook.Languages{
		language.AmericanEnglish: gorillaNameAmericanEnglish,
	}
)

var (
	// Gorilla represents an gorilla animal in the Animal Crossing series.
	Gorilla = nook.Animal{
		Key:  nook.Key(gorilla),
		Name: gorillaName,
	}
)
