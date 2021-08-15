package anteater

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	anteaterNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Anteater"}
)

var (
	anteaterName = nook.Languages{
		language.AmericanEnglish: anteaterNameAmericanEnglish}
)

var (
	Anteater = nook.Animal{
		Name: anteaterName}
)
