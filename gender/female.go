package gender

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	femaleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Female"}
)

var (
	femaleName = nook.Languages{
		language.AmericanEnglish: femaleAmericanEnglishName}
)

var (
	Female = nook.Gender{
		Name: femaleName}
)
