package personality

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	// peppyAmericanEnglishName represents the American English name for the Peppy personality.
	peppyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Peppy"}
)

var (
	// peppyName contains the names of the Peppy personality in different languages.
	peppyName = nook.Languages{
		language.AmericanEnglish: peppyAmericanEnglishName}
)

var (
	// Peppy represents the Peppy personality in the Animal Crossing series.
	Peppy = nook.Personality{
		Name: peppyName}
)
