package gamecategory

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// spinoff is the common reference for the Spin-off game category.
	spinoff = "Spinoff"
)

var (
	// spinoffNameAmericanEnglish represents the Spin-off game category's name in American English.
	spinoffNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Spin-off",
	}
)

var (
	// spinoffName contains the localized names of the Spin-off game category.
	spinoffName = nook.Languages{
		language.AmericanEnglish: spinoffNameAmericanEnglish,
	}
)

var (
	// Spinoff represents the Spin-off game category in the Animal Crossing series.
	Spinoff = nook.GameCategory{
		Key:  nook.Key(spinoff),
		Name: spinoffName,
	}
)
