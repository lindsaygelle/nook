package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // kangaroo is the common reference for the Kangaroo animal type.
    kangaroo = "Kangaroo"
)

var (
	// kangarooNameAmericanEnglish represents the American English name for the Kangaroo animal type.
	kangarooNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    kangaroo,
	}
)

var (
	// kangarooName contains the names of the Kangaroo animal type in different languages.
	kangarooName = nook.Languages{
		language.AmericanEnglish: kangarooNameAmericanEnglish,
	}
)

var (
	// Kangaroo represents an kangaroo animal in the Animal Crossing series.
	Kangaroo = nook.Animal{
		Key:  nook.Key(kangaroo),
		Name: kangarooName,
	}
)
