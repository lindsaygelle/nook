package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	walrusNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Walrus"}
)

var (
	walrusName = nook.Languages{
		language.AmericanEnglish: walrusNameAmericanEnglish}
)

var (
	Walrus = nook.Animal{
		Name: walrusName}
)
