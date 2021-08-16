package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	turkeyNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Turkey"}
)

var (
	turkeyName = nook.Languages{
		language.AmericanEnglish: turkeyNameAmericanEnglish}
)

var (
	Turkey = nook.Animal{
		Name: turkeyName}
)
