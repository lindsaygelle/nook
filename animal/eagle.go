package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	eagleNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Eagle"}
)

var (
	eagleName = nook.Languages{
		language.AmericanEnglish: eagleNameAmericanEnglish}
)

var (
	Eagle = nook.Animal{
		Key:  nook.Key("Eagle"),
		Name: eagleName}
)
