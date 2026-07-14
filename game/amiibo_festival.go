package game

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// amiiboFestival is the common reference for Animal Crossing: amiibo Festival.
	amiiboFestival = "AmiiboFestival"
)

var (
	// amiiboFestivalNameAmericanEnglish represents Animal Crossing: amiibo Festival's name in American English.
	amiiboFestivalNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Animal Crossing: amiibo Festival",
	}
)

var (
	// amiiboFestivalName contains the localized names of Animal Crossing: amiibo Festival.
	amiiboFestivalName = nook.Languages{
		language.AmericanEnglish: amiiboFestivalNameAmericanEnglish,
	}
)

var (
	// AmiiboFestival represents Animal Crossing: amiibo Festival.
	AmiiboFestival = nook.Game{
		Key:          nook.Key(amiiboFestival),
		Name:         amiiboFestivalName,
		ReleaseOrder: 10,
	}
)
