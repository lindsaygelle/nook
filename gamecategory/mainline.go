package gamecategory

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// mainline is the common reference for the Mainline game category.
	mainline = "Mainline"
)

var (
	// mainlineNameAmericanEnglish represents the Mainline game category's name in American English.
	mainlineNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mainline",
	}
)

var (
	// mainlineName contains the localized names of the Mainline game category.
	mainlineName = nook.Languages{
		language.AmericanEnglish: mainlineNameAmericanEnglish,
	}
)

var (
	// Mainline represents the Mainline game category in the Animal Crossing series.
	Mainline = nook.GameCategory{
		Key:  nook.Key(mainline),
		Name: mainlineName,
	}
)
