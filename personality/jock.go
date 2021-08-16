package personality

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	jockAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Jock"}
)

var (
	jockName = nook.Languages{
		language.AmericanEnglish: jockAmericanEnglishName}
)

var (
	Jock = nook.Personality{
		Name: jockName}
)
