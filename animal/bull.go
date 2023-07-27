package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // bull is the common reference for the Bull animal type.
    bull = "Bull"
)

var (
	// bullNameAmericanEnglish represents the American English name for the Bull animal type.
	bullNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    bull,
	}
)

var (
	// bullName contains the names of the Bull animal type in different languages.
	bullName = nook.Languages{
		language.AmericanEnglish: bullNameAmericanEnglish,
	}
)

var (
	// Bull represents an bull animal in the Animal Crossing series.
	Bull = nook.Animal{
		Key:  nook.Key(bull),
		Name: bullName,
	}
)
