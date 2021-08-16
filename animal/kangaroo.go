package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	kangarooNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kangaroo"}
)

var (
	kangarooName = nook.Languages{
		language.AmericanEnglish: kangarooNameAmericanEnglish}
)

var (
	Kangaroo = nook.Animal{
		Name: kangarooName}
)
