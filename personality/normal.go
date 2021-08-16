package personality

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	normalAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Normal"}
)

var (
	normalName = nook.Languages{
		language.AmericanEnglish: normalAmericanEnglishName}
)

var (
	Normal = nook.Personality{
		Name: normalName}
)
