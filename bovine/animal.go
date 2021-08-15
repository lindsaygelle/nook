package bovine

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bovineNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bovine"}
)

var (
	bovineName = nook.Languages{
		language.AmericanEnglish: bovineNameAmericanEnglish}
)

var (
	Bovine = nook.Animal{
		Name: bovineName}
)
