package octopus

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	octopusNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Octopus"}
)

var (
	octopusName = nook.Languages{
		language.AmericanEnglish: octopusNameAmericanEnglish}
)

var (
	Octopus = nook.Animal{
		Name: octopusName}
)
