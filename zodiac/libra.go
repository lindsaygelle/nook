package zodiac

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// libra is the common reference for the Libra zodiac sign.
	libra = "Libra"
)

var (
	// libraNameAmericanEnglish represents the Libra zodiac sign's name in American English.
	libraNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Libra"}
)

var (
	// libraName contains the localized names of the Libra zodiac sign.
	libraName = nook.Languages{
		language.AmericanEnglish: libraNameAmericanEnglish}
)

var (
	// Libra represents the Libra zodiac sign.
	Libra = nook.ZodiacSign{
		Key:  nook.Key(libra),
		Name: libraName}
)
