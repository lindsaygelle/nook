package boar

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	boarNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Boar"}
)

var (
	boarName = nook.Languages{
		language.AmericanEnglish: boarNameAmericanEnglish}
)

var (
	Boar = nook.Animal{
		Name: boarName}
)
