package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	mouseNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mouse"}
)

var (
	mouseName = nook.Languages{
		language.AmericanEnglish: mouseNameAmericanEnglish}
)

var (
	Mouse = nook.Animal{
		Key:  nook.Key("Mouse"),
		Name: mouseName}
)
