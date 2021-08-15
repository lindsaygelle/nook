package owl

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	owlNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Owl"}
)

var (
	owlName = nook.Languages{
		language.AmericanEnglish: owlNameAmericanEnglish}
)

var (
	Owl = nook.Animal{
		Name: owlName}
)
