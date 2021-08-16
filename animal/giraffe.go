package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	giraffeNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Giraffe"}
)

var (
	giraffeName = nook.Languages{
		language.AmericanEnglish: giraffeNameAmericanEnglish}
)

var (
	Giraffe = nook.Animal{
		Name: giraffeName}
)
