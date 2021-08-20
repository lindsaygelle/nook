package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	frillneckedlizardNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Frill-necked Lizard"}
)

var (
	frillneckedlizardName = nook.Languages{
		language.AmericanEnglish: frillneckedlizardNameAmericanEnglish}
)

var (
	Frillneckedlizard = nook.Animal{
		Key:  nook.Key("Frillneckedlizard"),
		Name: frillneckedlizardName}
)
