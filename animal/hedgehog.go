package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // hedgehog is the common reference for the Hedgehog animal type.
    hedgehog = "Hedgehog"
)

var (
	// hedgehogNameAmericanEnglish represents the American English name for the Hedgehog animal type.
	hedgehogNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    hedgehog,
	}
)

var (
	// hedgehogName contains the names of the Hedgehog animal type in different languages.
	hedgehogName = nook.Languages{
		language.AmericanEnglish: hedgehogNameAmericanEnglish,
	}
)

var (
	// Hedgehog represents an hedgehog animal in the Animal Crossing series.
	Hedgehog = nook.Animal{
		Key:  nook.Key(hedgehog),
		Name: hedgehogName,
	}
)
