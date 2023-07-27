package gender

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	// maleAmericanEnglishName represents the American English name for the Male gender type.
	maleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Male"}
)

var (
	// maleName contains the names of the Male gender type in different languages.
	maleName = nook.Languages{
		language.AmericanEnglish: maleAmericanEnglishName}
)

var (
	// Male represents the Male gender type in the Animal Crossing series.
	Male = nook.Gender{
		Name: maleName}
)
