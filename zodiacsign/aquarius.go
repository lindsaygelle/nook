package zodiacsign

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// aquarius is the common reference for the Aquarius zodiac sign.
	aquarius = "Aquarius"
)

var (
	// aquariusNameAmericanEnglish represents the Aquarius zodiac sign's name in American English.
	aquariusNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Aquarius"}
)

var (
	// aquariusName contains the localized names of the Aquarius zodiac sign.
	aquariusName = nook.Languages{
		language.AmericanEnglish: aquariusNameAmericanEnglish}
)

var (
	// Aquarius represents the Aquarius zodiac sign.
	Aquarius = nook.ZodiacSign{
		Key:  nook.Key(aquarius),
		Name: aquariusName}
)
