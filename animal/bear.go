package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bearNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bear"}
)

var (
	bearName = nook.Languages{
		language.AmericanEnglish: bearNameAmericanEnglish}
)

var (
	Bear = nook.Animal{
		Key:  nook.Key("Bear"),
		Name: bearName}
)
