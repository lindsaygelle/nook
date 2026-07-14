package region

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// worldwide is the common reference for the Worldwide region.
	worldwide = "Worldwide"
)

var (
	// worldwideNameAmericanEnglish represents the Worldwide region's name in American English.
	worldwideNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Worldwide",
	}
)

var (
	// worldwideName contains the localized names of the Worldwide region.
	worldwideName = nook.Languages{
		language.AmericanEnglish: worldwideNameAmericanEnglish,
	}
)

var (
	// Worldwide represents the Worldwide region in the Animal Crossing series.
	Worldwide = nook.Region{
		Key:  nook.Key(worldwide),
		Name: worldwideName,
	}
)
