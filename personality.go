package nook

import "golang.org/x/text/language"

// Personality is an Animal Crossing character personality type.
type Personality struct {
	Name Languages
}

var (
	bigSisterAmericanEnglishName = Name{
		Language: language.AmericanEnglish,
		Value:    "Big Sister"}
)

var (
	bigSisterName = Languages{
		language.AmericanEnglish: bigSisterAmericanEnglishName}
)

var (
	BigSister = Personality{
		Name: bigSisterName}
)

var (
	crankyAmericanEnglishName = Name{
		Language: language.AmericanEnglish,
		Value:    "Cranky"}
)

var (
	crankyName = Languages{
		language.AmericanEnglish: crankyAmericanEnglishName}
)

var (
	Cranky = Personality{
		Name: crankyName}
)

var (
	jockAmericanEnglishName = Name{
		Language: language.AmericanEnglish,
		Value:    "Jock"}
)

var (
	jockName = Languages{
		language.AmericanEnglish: jockAmericanEnglishName}
)

var (
	Jock = Personality{
		Name: jockName}
)

var (
	lazyAmericanEnglishName = Name{
		Language: language.AmericanEnglish,
		Value:    "Lazy"}
)

var (
	lazyName = Languages{
		language.AmericanEnglish: lazyAmericanEnglishName}
)

var (
	Lazy = Personality{
		Name: lazyName}
)

var (
	normalAmericanEnglishName = Name{
		Language: language.AmericanEnglish,
		Value:    "Normal"}
)

var (
	normalName = Languages{
		language.AmericanEnglish: normalAmericanEnglishName}
)

var (
	Normal = Personality{
		Name: normalName}
)

var (
	peppyAmericanEnglishName = Name{
		Language: language.AmericanEnglish,
		Value:    "Peppy"}
)

var (
	peppyName = Languages{
		language.AmericanEnglish: peppyAmericanEnglishName}
)

var (
	Peppy = Personality{
		Name: peppyName}
)

var (
	smugAmericanEnglishName = Name{
		Language: language.AmericanEnglish,
		Value:    "Smug"}
)

var (
	smugName = Languages{
		language.AmericanEnglish: smugAmericanEnglishName}
)

var (
	Smug = Personality{
		Name: smugName}
)

var (
	snootyAmericanEnglishName = Name{
		Language: language.AmericanEnglish,
		Value:    "Snooty"}
)

var (
	snootyName = Languages{
		language.AmericanEnglish: snootyAmericanEnglishName}
)

var (
	Snooty = Personality{
		Name: snootyName}
)
