package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // anteater is the common reference for the Anteater animal type.
    anteater = "Anteater"
)

var (
	// anteaterNameAmericanEnglish represents the American English name for the Anteater animal type.
	anteaterNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    anteater,
	}
)

var (
	// anteaterName contains the names of the Anteater animal type in different languages.
	anteaterName = nook.Languages{
		language.AmericanEnglish: anteaterNameAmericanEnglish,
	}
)

var (
	// Anteater represents an anteater animal in the Animal Crossing series.
	Anteater = nook.Animal{
		Key:  nook.Key(anteater),
		Name: anteaterName,
	}
)
