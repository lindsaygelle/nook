package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // rabbit is the common reference for the Rabbit animal type.
    rabbit = "Rabbit"
)

var (
	// rabbitNameAmericanEnglish represents the American English name for the Rabbit animal type.
	rabbitNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    rabbit,
	}
)

var (
	// rabbitName contains the names of the Rabbit animal type in different languages.
	rabbitName = nook.Languages{
		language.AmericanEnglish: rabbitNameAmericanEnglish,
	}
)

var (
	// Rabbit represents an rabbit animal in the Animal Crossing series.
	Rabbit = nook.Animal{
		Key:  nook.Key(rabbit),
		Name: rabbitName,
	}
)
