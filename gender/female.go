package gender

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	// femaleAmericanEnglishName represents the American English name for the Female gender type.
	femaleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Female"}
)

var (
	// femaleName contains the names of the Female gender type in different languages.
	femaleName = nook.Languages{
		language.AmericanEnglish: femaleAmericanEnglishName}
)

var (
	// Female represents the Female gender type in the Animal Crossing series.
	Female = nook.Gender{
		Name: femaleName}
)
