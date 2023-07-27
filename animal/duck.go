package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // duck is the common reference for the Duck animal type.
    duck = "Duck"
)

var (
	// duckNameAmericanEnglish represents the American English name for the Duck animal type.
	duckNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    duck,
	}
)

var (
	// duckName contains the names of the Duck animal type in different languages.
	duckName = nook.Languages{
		language.AmericanEnglish: duckNameAmericanEnglish,
	}
)

var (
	// Duck represents an duck animal in the Animal Crossing series.
	Duck = nook.Animal{
		Key:  nook.Key(duck),
		Name: duckName,
	}
)
