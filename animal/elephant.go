package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // elephant is the common reference for the Elephant animal type.
    elephant = "Elephant"
)

var (
	// elephantNameAmericanEnglish represents the American English name for the Elephant animal type.
	elephantNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    elephant,
	}
)

var (
	// elephantName contains the names of the Elephant animal type in different languages.
	elephantName = nook.Languages{
		language.AmericanEnglish: elephantNameAmericanEnglish,
	}
)

var (
	// Elephant represents an elephant animal in the Animal Crossing series.
	Elephant = nook.Animal{
		Key:  nook.Key(elephant),
		Name: elephantName,
	}
)
