package personality

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bigSisterAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Big Sister"}
)

var (
	bigSisterName = nook.Languages{
		language.AmericanEnglish: bigSisterAmericanEnglishName}
)

var (
	BigSister = nook.Personality{
		Name: bigSisterName}
)
