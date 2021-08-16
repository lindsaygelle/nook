package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	pelicanNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pelican"}
)

var (
	pelicanName = nook.Languages{
		language.AmericanEnglish: pelicanNameAmericanEnglish}
)

var (
	Pelican = nook.Animal{
		Name: pelicanName}
)
