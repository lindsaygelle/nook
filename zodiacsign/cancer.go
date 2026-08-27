package zodiacsign

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// cancer is the common reference for the Cancer zodiac sign.
	cancer = "Cancer"
)

var (
	// cancerNameAmericanEnglish represents the Cancer zodiac sign's name in American English.
	cancerNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cancer"}
)

var (
	// cancerName contains the localized names of the Cancer zodiac sign.
	cancerName = nook.Languages{
		language.AmericanEnglish: cancerNameAmericanEnglish}
)

var (
	// Cancer represents the Cancer zodiac sign.
	Cancer = nook.ZodiacSign{
		Key:  nook.Key(cancer),
		Name: cancerName}
)
