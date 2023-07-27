package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // sheep is the common reference for the Sheep animal type.
    sheep = "Sheep"
)

var (
	// sheepNameAmericanEnglish represents the American English name for the Sheep animal type.
	sheepNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    sheep,
	}
)

var (
	// sheepName contains the names of the Sheep animal type in different languages.
	sheepName = nook.Languages{
		language.AmericanEnglish: sheepNameAmericanEnglish,
	}
)

var (
	// Sheep represents an sheep animal in the Animal Crossing series.
	Sheep = nook.Animal{
		Key:  nook.Key(sheep),
		Name: sheepName,
	}
)
