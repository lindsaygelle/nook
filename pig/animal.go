package pig

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	pigNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pig"}
)

var (
	pigName = nook.Languages{
		language.AmericanEnglish: pigNameAmericanEnglish}
)

var (
	Pig = nook.Animal{
		Name: pigName}
)
