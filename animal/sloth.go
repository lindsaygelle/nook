package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // sloth is the common reference for the Sloth animal type.
    sloth = "Sloth"
)

var (
	// slothNameAmericanEnglish represents the American English name for the Sloth animal type.
	slothNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    sloth,
	}
)

var (
	// slothName contains the names of the Sloth animal type in different languages.
	slothName = nook.Languages{
		language.AmericanEnglish: slothNameAmericanEnglish,
	}
)

var (
	// Sloth represents an sloth animal in the Animal Crossing series.
	Sloth = nook.Animal{
		Key:  nook.Key(sloth),
		Name: slothName,
	}
)
