package zodiac

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// pisces is the common reference for the Pisces zodiac sign.
	pisces = "Pisces"
)

var (
	// piscesNameAmericanEnglish represents the Pisces zodiac sign's name in American English.
	piscesNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pisces"}
)

var (
	// piscesName contains the localized names of the Pisces zodiac sign.
	piscesName = nook.Languages{
		language.AmericanEnglish: piscesNameAmericanEnglish}
)

var (
	// Pisces represents the Pisces zodiac sign.
	Pisces = nook.ZodiacSign{
		Key:  nook.Key(pisces),
		Name: piscesName}
)
