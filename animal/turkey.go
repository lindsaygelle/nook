package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // turkey is the common reference for the Turkey animal type.
    turkey = "Turkey"
)

var (
	// turkeyNameAmericanEnglish represents the American English name for the Turkey animal type.
	turkeyNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    turkey,
	}
)

var (
	// turkeyName contains the names of the Turkey animal type in different languages.
	turkeyName = nook.Languages{
		language.AmericanEnglish: turkeyNameAmericanEnglish,
	}
)

var (
	// Turkey represents an turkey animal in the Animal Crossing series.
	Turkey = nook.Animal{
		Key:  nook.Key(turkey),
		Name: turkeyName,
	}
)
