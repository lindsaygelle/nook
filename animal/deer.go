package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	deerNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Deer"}
)

var (
	deerName = nook.Languages{
		language.AmericanEnglish: deerNameAmericanEnglish}
)

var (
	Deer = nook.Animal{
		Name: deerName}
)
