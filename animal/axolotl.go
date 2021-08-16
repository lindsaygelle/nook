package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	axolotlNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Axolotl"}
)

var (
	axolotlName = nook.Languages{
		language.AmericanEnglish: axolotlNameAmericanEnglish}
)

var (
	Axolotl = nook.Animal{
		Name: axolotlName}
)
