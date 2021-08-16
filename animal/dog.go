package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	dogNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dog"}
)

var (
	dogName = nook.Languages{
		language.AmericanEnglish: dogNameAmericanEnglish}
)

var (
	Dog = nook.Animal{
		Name: dogName}
)
