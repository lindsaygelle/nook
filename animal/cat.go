package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	catNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cat"}
)

var (
	catName = nook.Languages{
		language.AmericanEnglish: catNameAmericanEnglish}
)

var (
	Cat = nook.Animal{
		Name: catName}
)
