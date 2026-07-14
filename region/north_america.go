package region

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// northAmerica is the common reference for the North America region.
	northAmerica = "NorthAmerica"
)

var (
	// northAmericaNameAmericanEnglish represents the North America region's name in American English.
	northAmericaNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "North America",
	}
)

var (
	// northAmericaName contains the localized names of the North America region.
	northAmericaName = nook.Languages{
		language.AmericanEnglish: northAmericaNameAmericanEnglish,
	}
)

var (
	// NorthAmerica represents the North America region in the Animal Crossing series.
	NorthAmerica = nook.Region{
		Key:  nook.Key(northAmerica),
		Name: northAmericaName,
	}
)
