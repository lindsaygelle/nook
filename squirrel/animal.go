package squirrel

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	squirrelNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Squirrel"}
)

var (
	squirrelName = nook.Languages{
		language.AmericanEnglish: squirrelNameAmericanEnglish}
)

var (
	Squirrel = nook.Animal{
		Name: squirrelName}
)
