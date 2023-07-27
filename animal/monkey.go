package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // monkey is the common reference for the Monkey animal type.
    monkey = "Monkey"
)

var (
	// monkeyNameAmericanEnglish represents the American English name for the Monkey animal type.
	monkeyNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    monkey,
	}
)

var (
	// monkeyName contains the names of the Monkey animal type in different languages.
	monkeyName = nook.Languages{
		language.AmericanEnglish: monkeyNameAmericanEnglish,
	}
)

var (
	// Monkey represents an monkey animal in the Animal Crossing series.
	Monkey = nook.Animal{
		Key:  nook.Key(monkey),
		Name: monkeyName,
	}
)
