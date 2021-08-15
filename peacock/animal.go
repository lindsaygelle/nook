package peacock

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	peacockNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Peacock"}
)

var (
	peacockName = nook.Languages{
		language.AmericanEnglish: peacockNameAmericanEnglish}
)

var (
	Peacock = nook.Animal{
		Name: peacockName}
)
