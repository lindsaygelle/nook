package region

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// china is the common reference for the China region.
	china = "China"
)

var (
	// chinaNameAmericanEnglish represents the China region's name in American English.
	chinaNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "China",
	}
)

var (
	// chinaName contains the localized names of the China region.
	chinaName = nook.Languages{
		language.AmericanEnglish: chinaNameAmericanEnglish,
	}
)

var (
	// China represents the China region in the Animal Crossing series.
	China = nook.Region{
		Key:  nook.Key(china),
		Name: chinaName,
	}
)
