package platform

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// nintendoDS is the common reference for the Nintendo DS platform.
	nintendoDS = "NintendoDS"
)

var (
	// nintendoDSNameAmericanEnglish represents the Nintendo DS platform's name in American English.
	nintendoDSNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nintendo DS",
	}
)

var (
	// nintendoDSName contains the localized names of the Nintendo DS platform.
	nintendoDSName = nook.Languages{
		language.AmericanEnglish: nintendoDSNameAmericanEnglish,
	}
)

var (
	// NintendoDS represents the Nintendo DS platform in the Animal Crossing series.
	NintendoDS = nook.Platform{
		Key:  nook.Key(nintendoDS),
		Name: nintendoDSName,
	}
)
