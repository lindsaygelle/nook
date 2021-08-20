package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	lionNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lion"}
)

var (
	lionName = nook.Languages{
		language.AmericanEnglish: lionNameAmericanEnglish}
)

var (
	Lion = nook.Animal{
		Key:  nook.Key("Lion"),
		Name: lionName}
)
