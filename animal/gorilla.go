package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	gorillaNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gorilla"}
)

var (
	gorillaName = nook.Languages{
		language.AmericanEnglish: gorillaNameAmericanEnglish}
)

var (
	Gorilla = nook.Animal{
		Name: gorillaName}
)
