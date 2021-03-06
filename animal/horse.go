package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	horseNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Horse"}
)

var (
	horseName = nook.Languages{
		language.AmericanEnglish: horseNameAmericanEnglish}
)

var (
	Horse = nook.Animal{
		Key:  nook.Key("Horse"),
		Name: horseName}
)
