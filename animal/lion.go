package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // lion is the common reference for the Lion animal type.
    lion = "Lion"
)

var (
	// lionNameAmericanEnglish represents the American English name for the Lion animal type.
	lionNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    lion,
	}
)

var (
	// lionName contains the names of the Lion animal type in different languages.
	lionName = nook.Languages{
		language.AmericanEnglish: lionNameAmericanEnglish,
	}
)

var (
	// Lion represents an lion animal in the Animal Crossing series.
	Lion = nook.Animal{
		Key:  nook.Key(lion),
		Name: lionName,
	}
)
