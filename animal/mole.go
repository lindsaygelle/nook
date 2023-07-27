package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // mole is the common reference for the Mole animal type.
    mole = "Mole"
)

var (
	// moleNameAmericanEnglish represents the American English name for the Mole animal type.
	moleNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    mole,
	}
)

var (
	// moleName contains the names of the Mole animal type in different languages.
	moleName = nook.Languages{
		language.AmericanEnglish: moleNameAmericanEnglish,
	}
)

var (
	// Mole represents an mole animal in the Animal Crossing series.
	Mole = nook.Animal{
		Key:  nook.Key(mole),
		Name: moleName,
	}
)
