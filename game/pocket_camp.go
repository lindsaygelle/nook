package game

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// pocketCamp is the common reference for Animal Crossing: Pocket Camp.
	pocketCamp = "PocketCamp"
)

var (
	// pocketCampNameAmericanEnglish represents Animal Crossing: Pocket Camp's name in American English.
	pocketCampNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Animal Crossing: Pocket Camp",
	}
)

var (
	// pocketCampName contains the localized names of Animal Crossing: Pocket Camp.
	pocketCampName = nook.Languages{
		language.AmericanEnglish: pocketCampNameAmericanEnglish,
	}
)

var (
	// PocketCamp represents Animal Crossing: Pocket Camp.
	PocketCamp = nook.Game{
		Key:          nook.Key(pocketCamp),
		Name:         pocketCampName,
		ReleaseOrder: 11,
	}
)
