package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // wolf is the common reference for the Wolf animal type.
    wolf = "Wolf"
)

var (
	// wolfNameAmericanEnglish represents the American English name for the Wolf animal type.
	wolfNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    wolf,
	}
)

var (
	// wolfName contains the names of the Wolf animal type in different languages.
	wolfName = nook.Languages{
		language.AmericanEnglish: wolfNameAmericanEnglish,
	}
)

var (
	// Wolf represents an wolf animal in the Animal Crossing series.
	Wolf = nook.Animal{
		Key:  nook.Key(wolf),
		Name: wolfName,
	}
)
