package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // peacock is the common reference for the Peacock animal type.
    peacock = "Peacock"
)

var (
	// peacockNameAmericanEnglish represents the American English name for the Peacock animal type.
	peacockNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    peacock,
	}
)

var (
	// peacockName contains the names of the Peacock animal type in different languages.
	peacockName = nook.Languages{
		language.AmericanEnglish: peacockNameAmericanEnglish,
	}
)

var (
	// Peacock represents an peacock animal in the Animal Crossing series.
	Peacock = nook.Animal{
		Key:  nook.Key(peacock),
		Name: peacockName,
	}
)
