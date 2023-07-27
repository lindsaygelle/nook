package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // tortoise is the common reference for the Tortoise animal type.
    tortoise = "Tortoise"
)

var (
	// tortoiseNameAmericanEnglish represents the American English name for the Tortoise animal type.
	tortoiseNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    tortoise,
	}
)

var (
	// tortoiseName contains the names of the Tortoise animal type in different languages.
	tortoiseName = nook.Languages{
		language.AmericanEnglish: tortoiseNameAmericanEnglish,
	}
)

var (
	// Tortoise represents an tortoise animal in the Animal Crossing series.
	Tortoise = nook.Animal{
		Key:  nook.Key(tortoise),
		Name: tortoiseName,
	}
)
