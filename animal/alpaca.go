package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	alpacaNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Alpaca"}
)

var (
	alpacaName = nook.Languages{
		language.AmericanEnglish: alpacaNameAmericanEnglish}
)

var (
	Alpaca = nook.Animal{
		Key:  nook.Key("Alpaca"),
		Name: alpacaName}
)
