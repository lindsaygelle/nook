package nook

import "golang.org/x/text/language"

type Gender struct {
	Name Languages
}

var (
	femaleAmericanEnglishName = Name{
		Language: language.AmericanEnglish,
		Value:    "Female"}
)

var (
	femaleName = Languages{
		language.AmericanEnglish: femaleAmericanEnglishName}
)

var (
	Female = Gender{
		Name: femaleName}
)

var (
	maleAmericanEnglishName = Name{
		Language: language.AmericanEnglish,
		Value:    "Male"}
)

var (
	maleName = Languages{
		language.AmericanEnglish: maleAmericanEnglishName}
)

var (
	Male = Gender{
		Name: maleName}
)
