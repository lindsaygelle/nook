package personality

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// bigSister is the common reference for the BigSister personality.
	bigSister = "BigSister"
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
		Key:  nook.Key(bigSister),
		Name: bigSisterName}
)
