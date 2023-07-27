package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // dodo is the common reference for the Dodo animal type.
    dodo = "Dodo"
)

var (
	// dodoNameAmericanEnglish represents the American English name for the Dodo animal type.
	dodoNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    dodo,
	}
)

var (
	// dodoName contains the names of the Dodo animal type in different languages.
	dodoName = nook.Languages{
		language.AmericanEnglish: dodoNameAmericanEnglish,
	}
)

var (
	// Dodo represents an dodo animal in the Animal Crossing series.
	Dodo = nook.Animal{
		Key:  nook.Key(dodo),
		Name: dodoName,
	}
)
