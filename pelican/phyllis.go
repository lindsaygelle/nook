package pelican

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	phyllisBirthday = nook.Birthday{
		Day:   21,
		Month: time.November}
)

var (
	phyllisCode = nook.Code{
		Value: "pgb/plm"}
)

var (
	phyllisAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Phyllis"}

	phyllisCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Élisabec"}

	phyllisDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Phyllis"}

	phyllisFrenchName = nook.Name{
		Language: language.French,
		Value:    "Élisabec"}

	phyllisGermanName = nook.Name{
		Language: language.German,
		Value:    "Peggy"}

	phyllisItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Polly"}

	phyllisJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ぺりみ"}

	phyllisKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "펠리미"}

	phyllisLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Estrella"}

	phyllisRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Филлис"}

	phyllisSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "宋信美"}

	phyllisSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Estrella"}

	phyllisTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "宋信美"}
)

var (
	phyllisName = nook.Languages{
		language.AmericanEnglish:      phyllisAmericanEnglishName,
		language.CanadianFrench:       phyllisCanadianFrenchName,
		language.Dutch:                phyllisDutchName,
		language.French:               phyllisFrenchName,
		language.German:               phyllisGermanName,
		language.Italian:              phyllisItalianName,
		language.Japanese:             phyllisJapaneseName,
		language.Korean:               phyllisKoreanName,
		language.LatinAmericanSpanish: phyllisLatinAmericanSpanishName,
		language.Russian:              phyllisRussianName,
		language.SimplifiedChinese:    phyllisSimplifiedChineseName,
		language.Spanish:              phyllisSpanishName,
		language.TraditionalChinese:   phyllisTraditionalChineseName}
)

var (
	phyllisCharacter = nook.Character{
		Animal:   Pelican,
		Birthday: phyllisBirthday,
		Code:     phyllisCode,
		Gender:   nook.Female,
		Name:     phyllisName}
)

var (
	Phyllis = nook.Resident{
		Character: phyllisCharacter}
)
