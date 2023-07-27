package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // gyroid is the common reference for the Gyroid animal type.
    gyroid = "Gyroid"
)

var (
	// gyroidNameAmericanEnglish represents the American English name for the Gyroid animal type.
	gyroidNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    gyroid,
	}
)

var (
	// gyroidName contains the names of the Gyroid animal type in different languages.
	gyroidName = nook.Languages{
		language.AmericanEnglish: gyroidNameAmericanEnglish,
	}
)

var (
	// Gyroid represents an gyroid animal in the Animal Crossing series.
	Gyroid = nook.Animal{
		Key:  nook.Key(gyroid),
		Name: gyroidName,
	}
)
