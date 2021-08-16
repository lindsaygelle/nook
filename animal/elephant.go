package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	elephantNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Elephant"}
)

var (
	elephantName = nook.Languages{
		language.AmericanEnglish: elephantNameAmericanEnglish}
)

var (
	Elephant = nook.Animal{
		Name: elephantName}
)
