package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	chickenNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Chicken"}
)

var (
	chickenName = nook.Languages{
		language.AmericanEnglish: chickenNameAmericanEnglish}
)

var (
	Chicken = nook.Animal{
		Key:  nook.Key("Chicken"),
		Name: chickenName}
)
