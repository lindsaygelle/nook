package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // cow is the common reference for the Cow animal type.
    cow = "Cow"
)

var (
	// cowNameAmericanEnglish represents the American English name for the Cow animal type.
	cowNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    cow,
	}
)

var (
	// cowName contains the names of the Cow animal type in different languages.
	cowName = nook.Languages{
		language.AmericanEnglish: cowNameAmericanEnglish,
	}
)

var (
	// Cow represents an cow animal in the Animal Crossing series.
	Cow = nook.Animal{
		Key:  nook.Key(cow),
		Name: cowName,
	}
)
