package sloth

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	slothNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sloth"}
)

var (
	slothName = nook.Languages{
		language.AmericanEnglish: slothNameAmericanEnglish}
)

var (
	Sloth = nook.Animal{
		Name: slothName}
)
