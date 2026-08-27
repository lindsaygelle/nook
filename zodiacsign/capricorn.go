package zodiacsign

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// capricorn is the common reference for the Capricorn zodiac sign.
	capricorn = "Capricorn"
)

var (
	// capricornNameAmericanEnglish represents the Capricorn zodiac sign's name in American English.
	capricornNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Capricorn"}
)

var (
	// capricornName contains the localized names of the Capricorn zodiac sign.
	capricornName = nook.Languages{
		language.AmericanEnglish: capricornNameAmericanEnglish}
)

var (
	// Capricorn represents the Capricorn zodiac sign.
	Capricorn = nook.ZodiacSign{
		Key:  nook.Key(capricorn),
		Name: capricornName}
)
