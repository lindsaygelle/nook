package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // goat is the common reference for the Goat animal type.
    goat = "Goat"
)

var (
	// goatNameAmericanEnglish represents the American English name for the Goat animal type.
	goatNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    goat,
	}
)

var (
	// goatName contains the names of the Goat animal type in different languages.
	goatName = nook.Languages{
		language.AmericanEnglish: goatNameAmericanEnglish,
	}
)

var (
	// Goat represents an goat animal in the Animal Crossing series.
	Goat = nook.Animal{
		Key:  nook.Key(goat),
		Name: goatName,
	}
)
