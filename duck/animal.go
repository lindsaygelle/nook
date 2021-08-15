package duck

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	duckNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Duck"}
)

var (
	duckName = nook.Languages{
		language.AmericanEnglish: duckNameAmericanEnglish}
)

var (
	Duck = nook.Animal{
		Name: duckName}
)
