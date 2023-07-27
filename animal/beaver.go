package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // beaver is the common reference for the Beaver animal type.
    beaver = "Beaver"
)

var (
	// beaverNameAmericanEnglish represents the American English name for the Beaver animal type.
	beaverNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    beaver,
	}
)

var (
	// beaverName contains the names of the Beaver animal type in different languages.
	beaverName = nook.Languages{
		language.AmericanEnglish: beaverNameAmericanEnglish,
	}
)

var (
	// Beaver represents an beaver animal in the Animal Crossing series.
	Beaver = nook.Animal{
		Key:  nook.Key(beaver),
		Name: beaverName,
	}
)
