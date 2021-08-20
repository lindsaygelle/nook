package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	dodoNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dodo"}
)

var (
	dodoName = nook.Languages{
		language.AmericanEnglish: dodoNameAmericanEnglish}
)

var (
	Dodo = nook.Animal{
		Key:  nook.Key("Dodo"),
		Name: dodoName}
)
