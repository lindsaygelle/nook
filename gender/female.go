package gender

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// female is the common reference for the Female gender type.
	female = "Female"
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
		Key:  nook.Key(female),
		Name: femaleName}
)
