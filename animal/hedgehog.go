package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	hedgehogNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hedgehog"}
)

var (
	hedgehogName = nook.Languages{
		language.AmericanEnglish: hedgehogNameAmericanEnglish}
)

var (
	Hedgehog = nook.Animal{
		Name: hedgehogName}
)
