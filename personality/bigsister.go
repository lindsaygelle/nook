package personality

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	// bigSisterAmericanEnglishName represents the American English name for the BigSister personality.
	bigSisterAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "BigSister"}
)

var (
	// bigSisterName contains the names of the BigSister personality in different languages.
	bigSisterName = nook.Languages{
		language.AmericanEnglish: bigSisterAmericanEnglishName}
)

var (
	// BigSister represents the BigSister personality in the Animal Crossing series.
	BigSister = nook.Personality{
		Name: bigSisterName}
)
