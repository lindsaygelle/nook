package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // skunk is the common reference for the Skunk animal type.
    skunk = "Skunk"
)

var (
	// skunkNameAmericanEnglish represents the American English name for the Skunk animal type.
	skunkNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    skunk,
	}
)

var (
	// skunkName contains the names of the Skunk animal type in different languages.
	skunkName = nook.Languages{
		language.AmericanEnglish: skunkNameAmericanEnglish,
	}
)

var (
	// Skunk represents an skunk animal in the Animal Crossing series.
	Skunk = nook.Animal{
		Key:  nook.Key(skunk),
		Name: skunkName,
	}
)
