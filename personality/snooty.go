package personality

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	snootyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Snooty"}
)

var (
	snootyName = nook.Languages{
		language.AmericanEnglish: snootyAmericanEnglishName}
)

var (
	Snooty = nook.Personality{
		Name: snootyName}
)
