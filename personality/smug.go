package personality

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	// smugAmericanEnglishName represents the American English name for the Smug personality.
	smugAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Smug"}
)

var (
	// smugName contains the names of the Smug personality in different languages.
	smugName = nook.Languages{
		language.AmericanEnglish: smugAmericanEnglishName}
)

var (
	// Smug represents the Smug personality in the Animal Crossing series.
	Smug = nook.Personality{
		Name: smugName}
)
