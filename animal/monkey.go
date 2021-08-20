package animal

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
		Key:  nook.Key("Monkey"),
		Name: monkeyName}
)
