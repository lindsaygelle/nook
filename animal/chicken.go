package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // chicken is the common reference for the Chicken animal type.
    chicken = "Chicken"
)

var (
	// chickenNameAmericanEnglish represents the American English name for the Chicken animal type.
	chickenNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    chicken,
	}
)

var (
	// chickenName contains the names of the Chicken animal type in different languages.
	chickenName = nook.Languages{
		language.AmericanEnglish: chickenNameAmericanEnglish,
	}
)

var (
	// Chicken represents an chicken animal in the Animal Crossing series.
	Chicken = nook.Animal{
		Key:  nook.Key(chicken),
		Name: chickenName,
	}
)
