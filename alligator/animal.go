package alligator

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	alligatorNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Alligator"}
)

var (
	alligatorName = nook.Languages{
		language.AmericanEnglish: alligatorNameAmericanEnglish}
)

var (
	Alligator = nook.Animal{
		Name: alligatorName}
)
