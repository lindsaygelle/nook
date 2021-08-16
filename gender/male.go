package gender

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	maleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Male"}
)

var (
	maleName = nook.Languages{
		language.AmericanEnglish: maleAmericanEnglishName}
)

var (
	Male = nook.Gender{
		Name: maleName}
)
