package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	pantherNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Panther"}
)

var (
	pantherName = nook.Languages{
		language.AmericanEnglish: pantherNameAmericanEnglish}
)

var (
	Panther = nook.Animal{
		Key:  nook.Key("Panther"),
		Name: pantherName}
)
