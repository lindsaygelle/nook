package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	penguinNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Penguin"}
)

var (
	penguinName = nook.Languages{
		language.AmericanEnglish: penguinNameAmericanEnglish}
)

var (
	Penguin = nook.Animal{
		Name: penguinName}
)
