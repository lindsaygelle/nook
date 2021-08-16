package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bullNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bull"}
)

var (
	bullName = nook.Languages{
		language.AmericanEnglish: bullNameAmericanEnglish}
)

var (
	Bull = nook.Animal{
		Name: bullName}
)
