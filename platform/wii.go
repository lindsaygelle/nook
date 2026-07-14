package platform

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// wii is the common reference for the Wii platform.
	wii = "Wii"
)

var (
	// wiiNameAmericanEnglish represents the Wii platform's name in American English.
	wiiNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Wii",
	}
)

var (
	// wiiName contains the localized names of the Wii platform.
	wiiName = nook.Languages{
		language.AmericanEnglish: wiiNameAmericanEnglish,
	}
)

var (
	// Wii represents the Wii platform in the Animal Crossing series.
	Wii = nook.Platform{
		Key:  nook.Key(wii),
		Name: wiiName,
	}
)
