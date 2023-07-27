package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // frog is the common reference for the Frog animal type.
    frog = "Frog"
)

var (
	// frogNameAmericanEnglish represents the American English name for the Frog animal type.
	frogNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    frog,
	}
)

var (
	// frogName contains the names of the Frog animal type in different languages.
	frogName = nook.Languages{
		language.AmericanEnglish: frogNameAmericanEnglish,
	}
)

var (
	// Frog represents an frog animal in the Animal Crossing series.
	Frog = nook.Animal{
		Key:  nook.Key(frog),
		Name: frogName,
	}
)
