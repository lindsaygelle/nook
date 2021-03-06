package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	cowNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cow"}
)

var (
	cowName = nook.Languages{
		language.AmericanEnglish: cowNameAmericanEnglish}
)

var (
	Cow = nook.Animal{
		Key:  nook.Key("Cow"),
		Name: cowName}
)
