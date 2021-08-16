package personality

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	peppyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Peppy"}
)

var (
	peppyName = nook.Languages{
		language.AmericanEnglish: peppyAmericanEnglishName}
)

var (
	Peppy = nook.Personality{
		Name: peppyName}
)
