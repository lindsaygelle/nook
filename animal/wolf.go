package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	wolfNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Wolf"}
)

var (
	wolfName = nook.Languages{
		language.AmericanEnglish: wolfNameAmericanEnglish}
)

var (
	Wolf = nook.Animal{
		Name: wolfName}
)
