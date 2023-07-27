package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // otter is the common reference for the Otter animal type.
    otter = "Otter"
)

var (
	// otterNameAmericanEnglish represents the American English name for the Otter animal type.
	otterNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    otter,
	}
)

var (
	// otterName contains the names of the Otter animal type in different languages.
	otterName = nook.Languages{
		language.AmericanEnglish: otterNameAmericanEnglish,
	}
)

var (
	// Otter represents an otter animal in the Animal Crossing series.
	Otter = nook.Animal{
		Key:  nook.Key(otter),
		Name: otterName,
	}
)
