package personality

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	smugAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Smug"}
)

var (
	smugName = nook.Languages{
		language.AmericanEnglish: smugAmericanEnglishName}
)

var (
	Smug = nook.Personality{
		Name: smugName}
)
