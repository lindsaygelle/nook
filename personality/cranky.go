package personality

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	// crankyAmericanEnglishName represents the American English name for the Cranky personality.
	crankyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cranky"}
)

var (
	// crankyName contains the names of the Cranky personality in different languages.
	crankyName = nook.Languages{
		language.AmericanEnglish: crankyAmericanEnglishName}
)

var (
	// Cranky represents the Cranky personality in the Animal Crossing series.
	Cranky = nook.Personality{
		Name: crankyName}
)
