package zodiacsign

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// virgo is the common reference for the Virgo zodiac sign.
	virgo = "Virgo"
)

var (
	// virgoNameAmericanEnglish represents the Virgo zodiac sign's name in American English.
	virgoNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Virgo"}
)

var (
	// virgoName contains the localized names of the Virgo zodiac sign.
	virgoName = nook.Languages{
		language.AmericanEnglish: virgoNameAmericanEnglish}
)

var (
	// Virgo represents the Virgo zodiac sign.
	Virgo = nook.ZodiacSign{
		Key:  nook.Key(virgo),
		Name: virgoName}
)
