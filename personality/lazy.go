package personality

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	// lazyAmericanEnglishName represents the American English name for the Lazy personality.
	lazyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lazy"}
)

var (
	// lazyName contains the names of the Lazy personality in different languages.
	lazyName = nook.Languages{
		language.AmericanEnglish: lazyAmericanEnglishName}
)

var (
	// Lazy represents the Lazy personality in the Animal Crossing series.
	Lazy = nook.Personality{
		Name: lazyName}
)
