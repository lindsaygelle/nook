package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	rhinocerosNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rhinoceros"}
)

var (
	rhinocerosName = nook.Languages{
		language.AmericanEnglish: rhinocerosNameAmericanEnglish}
)

var (
	Rhinoceros = nook.Animal{
		Name: rhinocerosName}
)
