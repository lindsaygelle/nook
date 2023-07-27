package personality

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	// jockAmericanEnglishName represents the American English name for the Jock personality.
	jockAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jock"}
)

var (
	// jockName contains the names of the Jock personality in different languages.
	jockName = nook.Languages{
		language.AmericanEnglish: jockAmericanEnglishName}
)

var (
	// Jock represents the Jock personality in the Animal Crossing series.
	Jock = nook.Personality{
		Name: jockName}
)
