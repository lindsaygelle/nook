package fox

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	foxNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Fox"}
)

var (
	foxName = nook.Languages{
		language.AmericanEnglish: foxNameAmericanEnglish}
)

var (
	Fox = nook.Animal{
		Name: foxName}
)
