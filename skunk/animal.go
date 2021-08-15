package skunk

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	skunkNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Skunk"}
)

var (
	skunkName = nook.Languages{
		language.AmericanEnglish: skunkNameAmericanEnglish}
)

var (
	Skunk = nook.Animal{
		Name: skunkName}
)
