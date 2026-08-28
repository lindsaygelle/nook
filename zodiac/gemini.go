package zodiac

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// gemini is the common reference for the Gemini zodiac sign.
	gemini = "Gemini"
)

var (
	// geminiNameAmericanEnglish represents the Gemini zodiac sign's name in American English.
	geminiNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gemini"}
)

var (
	// geminiName contains the localized names of the Gemini zodiac sign.
	geminiName = nook.Languages{
		language.AmericanEnglish: geminiNameAmericanEnglish}
)

var (
	// Gemini represents the Gemini zodiac sign.
	Gemini = nook.ZodiacSign{
		Key:  nook.Key(gemini),
		Name: geminiName}
)
