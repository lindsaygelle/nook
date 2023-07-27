package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // deer is the common reference for the Deer animal type.
    deer = "Deer"
)

var (
	// deerNameAmericanEnglish represents the American English name for the Deer animal type.
	deerNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    deer,
	}
)

var (
	// deerName contains the names of the Deer animal type in different languages.
	deerName = nook.Languages{
		language.AmericanEnglish: deerNameAmericanEnglish,
	}
)

var (
	// Deer represents an deer animal in the Animal Crossing series.
	Deer = nook.Animal{
		Key:  nook.Key(deer),
		Name: deerName,
	}
)
