package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	frogNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Frog"}
)

var (
	frogName = nook.Languages{
		language.AmericanEnglish: frogNameAmericanEnglish}
)

var (
	Frog = nook.Animal{
		Name: frogName}
)
