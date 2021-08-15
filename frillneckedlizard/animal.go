package frillneckedlizard

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	frillneckedlizardNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Frillneckedlizard"}
)

var (
	frillneckedlizardName = nook.Languages{
		language.AmericanEnglish: frillneckedlizardNameAmericanEnglish}
)

var (
	Frillneckedlizard = nook.Animal{
		Name: frillneckedlizardName}
)
