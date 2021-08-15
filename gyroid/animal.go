package gyroid

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	gyroidNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Gyroid"}
)

var (
	gyroidName = nook.Languages{
		language.AmericanEnglish: gyroidNameAmericanEnglish}
)

var (
	Gyroid = nook.Animal{
		Name: gyroidName}
)
