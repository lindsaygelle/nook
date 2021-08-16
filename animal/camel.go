package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	camelNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Camel"}
)

var (
	camelName = nook.Languages{
		language.AmericanEnglish: camelNameAmericanEnglish}
)

var (
	Camel = nook.Animal{
		Name: camelName}
)
