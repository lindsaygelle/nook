package platform

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// nintendoSwitch is the common reference for the Nintendo Switch platform.
	nintendoSwitch = "NintendoSwitch"
)

var (
	// nintendoSwitchNameAmericanEnglish represents the Nintendo Switch platform's name in American English.
	nintendoSwitchNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nintendo Switch",
	}
)

var (
	// nintendoSwitchName contains the localized names of the Nintendo Switch platform.
	nintendoSwitchName = nook.Languages{
		language.AmericanEnglish: nintendoSwitchNameAmericanEnglish,
	}
)

var (
	// NintendoSwitch represents the Nintendo Switch platform in the Animal Crossing series.
	NintendoSwitch = nook.Platform{
		Key:  nook.Key(nintendoSwitch),
		Name: nintendoSwitchName,
	}
)
