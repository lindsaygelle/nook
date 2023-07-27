package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // rhinoceros is the common reference for the Rhinoceros animal type.
    rhinoceros = "Rhinoceros"
)

var (
	// rhinocerosNameAmericanEnglish represents the American English name for the Rhinoceros animal type.
	rhinocerosNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    rhinoceros,
	}
)

var (
	// rhinocerosName contains the names of the Rhinoceros animal type in different languages.
	rhinocerosName = nook.Languages{
		language.AmericanEnglish: rhinocerosNameAmericanEnglish,
	}
)

var (
	// Rhinoceros represents an rhinoceros animal in the Animal Crossing series.
	Rhinoceros = nook.Animal{
		Key:  nook.Key(rhinoceros),
		Name: rhinocerosName,
	}
)
