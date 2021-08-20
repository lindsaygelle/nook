package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	fursealNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Fur seal"}
)

var (
	fursealName = nook.Languages{
		language.AmericanEnglish: fursealNameAmericanEnglish}
)

var (
	Furseal = nook.Animal{
		Key:  nook.Key("Furseal"),
		Name: fursealName}
)
