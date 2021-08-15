package monkey

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	monkeyNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Monkey"}
)

var (
	monkeyName = nook.Languages{
		language.AmericanEnglish: monkeyNameAmericanEnglish}
)

var (
	Monkey = nook.Animal{
		Name: monkeyName}
)
