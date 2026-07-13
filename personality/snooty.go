package personality

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// snooty is the common reference for the Snooty personality.
	snooty = "Snooty"
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
		Key:  nook.Key(snooty),
		Name: snootyName}
)
