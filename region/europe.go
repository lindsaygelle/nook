package region

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// europe is the common reference for the Europe region.
	europe = "Europe"
)

var (
	// europeNameAmericanEnglish represents the Europe region's name in American English.
	europeNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Europe",
	}
)

var (
	// europeName contains the localized names of the Europe region.
	europeName = nook.Languages{
		language.AmericanEnglish: europeNameAmericanEnglish,
	}
)

var (
	// Europe represents the Europe region in the Animal Crossing series.
	Europe = nook.Region{
		Key:  nook.Key(europe),
		Name: europeName,
	}
)
