package hippo

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	hippoNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hippo"}
)

var (
	hippoName = nook.Languages{
		language.AmericanEnglish: hippoNameAmericanEnglish}
)

var (
	Hippo = nook.Animal{
		Name: hippoName}
)
