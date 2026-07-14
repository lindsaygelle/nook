package platform

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// wiiU is the common reference for the Wii U platform.
	wiiU = "WiiU"
)

var (
	// wiiUNameAmericanEnglish represents the Wii U platform's name in American English.
	wiiUNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Wii U",
	}
)

var (
	// wiiUName contains the localized names of the Wii U platform.
	wiiUName = nook.Languages{
		language.AmericanEnglish: wiiUNameAmericanEnglish,
	}
)

var (
	// WiiU represents the Wii U platform in the Animal Crossing series.
	WiiU = nook.Platform{
		Key:  nook.Key(wiiU),
		Name: wiiUName,
	}
)
