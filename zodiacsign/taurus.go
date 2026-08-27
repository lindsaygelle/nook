package zodiacsign

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// taurus is the common reference for the Taurus zodiac sign.
	taurus = "Taurus"
)

var (
	// taurusNameAmericanEnglish represents the Taurus zodiac sign's name in American English.
	taurusNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Taurus"}
)

var (
	// taurusName contains the localized names of the Taurus zodiac sign.
	taurusName = nook.Languages{
		language.AmericanEnglish: taurusNameAmericanEnglish}
)

var (
	// Taurus represents the Taurus zodiac sign.
	Taurus = nook.ZodiacSign{
		Key:  nook.Key(taurus),
		Name: taurusName}
)
