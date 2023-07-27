package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // furseal is the common reference for the Furseal animal type.
    furseal = "Furseal"
)

var (
	// fursealNameAmericanEnglish represents the American English name for the Furseal animal type.
	fursealNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    furseal,
	}
)

var (
	// fursealName contains the names of the Furseal animal type in different languages.
	fursealName = nook.Languages{
		language.AmericanEnglish: fursealNameAmericanEnglish,
	}
)

var (
	// Furseal represents an furseal animal in the Animal Crossing series.
	Furseal = nook.Animal{
		Key:  nook.Key(furseal),
		Name: fursealName,
	}
)
