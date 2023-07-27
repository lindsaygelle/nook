package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // seagull is the common reference for the Seagull animal type.
    seagull = "Seagull"
)

var (
	// seagullNameAmericanEnglish represents the American English name for the Seagull animal type.
	seagullNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    seagull,
	}
)

var (
	// seagullName contains the names of the Seagull animal type in different languages.
	seagullName = nook.Languages{
		language.AmericanEnglish: seagullNameAmericanEnglish,
	}
)

var (
	// Seagull represents an seagull animal in the Animal Crossing series.
	Seagull = nook.Animal{
		Key:  nook.Key(seagull),
		Name: seagullName,
	}
)
