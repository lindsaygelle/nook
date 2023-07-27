package personality

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	// snootyAmericanEnglishName represents the American English name for the Snooty personality.
	snootyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Snooty"}
)

var (
	// snootyName contains the names of the Snooty personality in different languages.
	snootyName = nook.Languages{
		language.AmericanEnglish: snootyAmericanEnglishName}
)

var (
	// Snooty represents the Snooty personality in the Animal Crossing series.
	Snooty = nook.Personality{
		Name: snootyName}
)
