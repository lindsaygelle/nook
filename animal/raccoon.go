package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // raccoon is the common reference for the Raccoon animal type.
    raccoon = "Raccoon"
)

var (
	// raccoonNameAmericanEnglish represents the American English name for the Raccoon animal type.
	raccoonNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    raccoon,
	}
)

var (
	// raccoonName contains the names of the Raccoon animal type in different languages.
	raccoonName = nook.Languages{
		language.AmericanEnglish: raccoonNameAmericanEnglish,
	}
)

var (
	// Raccoon represents an raccoon animal in the Animal Crossing series.
	Raccoon = nook.Animal{
		Key:  nook.Key(raccoon),
		Name: raccoonName,
	}
)
