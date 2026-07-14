package region

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// australia is the common reference for the Australia region.
	australia = "Australia"
)

var (
	// australiaNameAmericanEnglish represents the Australia region's name in American English.
	australiaNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Australia",
	}
)

var (
	// australiaName contains the localized names of the Australia region.
	australiaName = nook.Languages{
		language.AmericanEnglish: australiaNameAmericanEnglish,
	}
)

var (
	// Australia represents the Australia region in the Animal Crossing series.
	Australia = nook.Region{
		Key:  nook.Key(australia),
		Name: australiaName,
	}
)
