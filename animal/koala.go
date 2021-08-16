package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	koalaNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Koala"}
)

var (
	koalaName = nook.Languages{
		language.AmericanEnglish: koalaNameAmericanEnglish}
)

var (
	Koala = nook.Animal{
		Name: koalaName}
)
