package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	turtleNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Turtle"}
)

var (
	turtleName = nook.Languages{
		language.AmericanEnglish: turtleNameAmericanEnglish}
)

var (
	Turtle = nook.Animal{
		Name: turtleName}
)
