package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	goatNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Goat"}
)

var (
	goatName = nook.Languages{
		language.AmericanEnglish: goatNameAmericanEnglish}
)

var (
	Goat = nook.Animal{
		Name: goatName}
)
