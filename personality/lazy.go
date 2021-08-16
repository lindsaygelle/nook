package personality

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	lazyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lazy"}
)

var (
	lazyName = nook.Languages{
		language.AmericanEnglish: lazyAmericanEnglishName}
)

var (
	Lazy = nook.Personality{
		Name: lazyName}
)
