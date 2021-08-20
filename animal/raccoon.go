package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	raccoonNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Raccoon"}
)

var (
	raccoonName = nook.Languages{
		language.AmericanEnglish: raccoonNameAmericanEnglish}
)

var (
	Raccoon = nook.Animal{
		Key:  nook.Key("Raccoon"),
		Name: raccoonName}
)
