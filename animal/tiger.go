package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	tigerNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tiger"}
)

var (
	tigerName = nook.Languages{
		language.AmericanEnglish: tigerNameAmericanEnglish}
)

var (
	Tiger = nook.Animal{
		Name: tigerName}
)
