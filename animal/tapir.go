package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // tapir is the common reference for the Tapir animal type.
    tapir = "Tapir"
)

var (
	// tapirNameAmericanEnglish represents the American English name for the Tapir animal type.
	tapirNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    tapir,
	}
)

var (
	// tapirName contains the names of the Tapir animal type in different languages.
	tapirName = nook.Languages{
		language.AmericanEnglish: tapirNameAmericanEnglish,
	}
)

var (
	// Tapir represents an tapir animal in the Animal Crossing series.
	Tapir = nook.Animal{
		Key:  nook.Key(tapir),
		Name: tapirName,
	}
)
