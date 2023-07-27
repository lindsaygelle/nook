package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // hamster is the common reference for the Hamster animal type.
    hamster = "Hamster"
)

var (
	// hamsterNameAmericanEnglish represents the American English name for the Hamster animal type.
	hamsterNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    hamster,
	}
)

var (
	// hamsterName contains the names of the Hamster animal type in different languages.
	hamsterName = nook.Languages{
		language.AmericanEnglish: hamsterNameAmericanEnglish,
	}
)

var (
	// Hamster represents an hamster animal in the Animal Crossing series.
	Hamster = nook.Animal{
		Key:  nook.Key(hamster),
		Name: hamsterName,
	}
)
