package zodiac

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// scorpio is the common reference for the Scorpio zodiac sign.
	scorpio = "Scorpio"
)

var (
	// scorpioNameAmericanEnglish represents the Scorpio zodiac sign's name in American English.
	scorpioNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Scorpio"}
)

var (
	// scorpioName contains the localized names of the Scorpio zodiac sign.
	scorpioName = nook.Languages{
		language.AmericanEnglish: scorpioNameAmericanEnglish}
)

var (
	// Scorpio represents the Scorpio zodiac sign.
	Scorpio = nook.ZodiacSign{
		Key:  nook.Key(scorpio),
		Name: scorpioName}
)
