package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // alligator is the common reference for the Alligator animal type.
    alligator = "Alligator"
)

var (
	// alligatorNameAmericanEnglish represents the American English name for the Alligator animal type.
	alligatorNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    alligator,
	}
)

var (
	// alligatorName contains the names of the Alligator animal type in different languages.
	alligatorName = nook.Languages{
		language.AmericanEnglish: alligatorNameAmericanEnglish,
	}
)

var (
	// Alligator represents an alligator animal in the Animal Crossing series.
	Alligator = nook.Animal{
		Key:  nook.Key(alligator),
		Name: alligatorName,
	}
)
