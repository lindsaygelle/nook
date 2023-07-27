package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // squirrel is the common reference for the Squirrel animal type.
    squirrel = "Squirrel"
)

var (
	// squirrelNameAmericanEnglish represents the American English name for the Squirrel animal type.
	squirrelNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    squirrel,
	}
)

var (
	// squirrelName contains the names of the Squirrel animal type in different languages.
	squirrelName = nook.Languages{
		language.AmericanEnglish: squirrelNameAmericanEnglish,
	}
)

var (
	// Squirrel represents an squirrel animal in the Animal Crossing series.
	Squirrel = nook.Animal{
		Key:  nook.Key(squirrel),
		Name: squirrelName,
	}
)
