package zodiac

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// sagittarius is the common reference for the Sagittarius zodiac sign.
	sagittarius = "Sagittarius"
)

var (
	// sagittariusNameAmericanEnglish represents the Sagittarius zodiac sign's name in American English.
	sagittariusNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sagittarius"}
)

var (
	// sagittariusName contains the localized names of the Sagittarius zodiac sign.
	sagittariusName = nook.Languages{
		language.AmericanEnglish: sagittariusNameAmericanEnglish}
)

var (
	// Sagittarius represents the Sagittarius zodiac sign.
	Sagittarius = nook.ZodiacSign{
		Key:  nook.Key(sagittarius),
		Name: sagittariusName}
)
