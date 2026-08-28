package zodiac

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// leo is the common reference for the Leo zodiac sign.
	leo = "Leo"
)

var (
	// leoNameAmericanEnglish represents the Leo zodiac sign's name in American English.
	leoNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Leo"}
)

var (
	// leoName contains the localized names of the Leo zodiac sign.
	leoName = nook.Languages{
		language.AmericanEnglish: leoNameAmericanEnglish}
)

var (
	// Leo represents the Leo zodiac sign.
	Leo = nook.ZodiacSign{
		Key:  nook.Key(leo),
		Name: leoName}
)
