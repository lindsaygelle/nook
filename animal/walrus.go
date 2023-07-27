package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // walrus is the common reference for the Walrus animal type.
    walrus = "Walrus"
)

var (
	// walrusNameAmericanEnglish represents the American English name for the Walrus animal type.
	walrusNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    walrus,
	}
)

var (
	// walrusName contains the names of the Walrus animal type in different languages.
	walrusName = nook.Languages{
		language.AmericanEnglish: walrusNameAmericanEnglish,
	}
)

var (
	// Walrus represents an walrus animal in the Animal Crossing series.
	Walrus = nook.Animal{
		Key:  nook.Key(walrus),
		Name: walrusName,
	}
)
