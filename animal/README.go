package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // README is the common reference for the Readme animal type.
    README = "Readme"
)

var (
	// READMENameAmericanEnglish represents the American English name for the Readme animal type.
	READMENameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    README,
	}
)

var (
	// READMEName contains the names of the Readme animal type in different languages.
	READMEName = nook.Languages{
		language.AmericanEnglish: READMENameAmericanEnglish,
	}
)

var (
	// Readme represents an README animal in the Animal Crossing series.
	Readme = nook.Animal{
		Key:  nook.Key(README),
		Name: READMEName,
	}
)
