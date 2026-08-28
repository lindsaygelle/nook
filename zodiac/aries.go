package zodiac

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// aries is the common reference for the Aries zodiac sign.
	aries = "Aries"
)

var (
	// ariesNameAmericanEnglish represents the Aries zodiac sign's name in American English.
	ariesNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Aries"}
)

var (
	// ariesName contains the localized names of the Aries zodiac sign.
	ariesName = nook.Languages{
		language.AmericanEnglish: ariesNameAmericanEnglish}
)

var (
	// Aries represents the Aries zodiac sign.
	Aries = nook.ZodiacSign{
		Key:  nook.Key(aries),
		Name: ariesName}
)
