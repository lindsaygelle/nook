package game

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// newHorizons is the common reference for Animal Crossing: New Horizons.
	newHorizons = "NewHorizons"
)

var (
	// newHorizonsNameAmericanEnglish represents Animal Crossing: New Horizons's name in American English.
	newHorizonsNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Animal Crossing: New Horizons",
	}
)

var (
	// newHorizonsName contains the localized names of Animal Crossing: New Horizons.
	newHorizonsName = nook.Languages{
		language.AmericanEnglish: newHorizonsNameAmericanEnglish,
	}
)

var (
	// NewHorizons represents Animal Crossing: New Horizons.
	NewHorizons = nook.Game{
		Key:          nook.Key(newHorizons),
		Name:         newHorizonsName,
		ReleaseOrder: 12,
	}
)
