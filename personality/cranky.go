package personality

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	crankyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cranky"}
)

var (
	crankyName = nook.Languages{
		language.AmericanEnglish: crankyAmericanEnglishName}
)

var (
	Cranky = nook.Personality{
		Name: crankyName}
)
