package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	chameleonNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chameleon"}
)

var (
	chameleonName = nook.Languages{
		language.AmericanEnglish: chameleonNameAmericanEnglish}
)

var (
	Chameleon = nook.Animal{
		Name: chameleonName}
)
