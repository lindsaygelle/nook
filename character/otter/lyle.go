package otter

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	lyleBirthday = nook.Birthday{
		Day:   6,
		Month: time.June}
)

var (
	lyleCode = nook.Code{
		Value: "ott"}
)

var (
	lyleAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Lyle"}

	lyleCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Lionel"}

	lyleDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Lyle"}

	lyleFrenchName = nook.Name{
		Language: language.French,
		Value:    "Lionel"}

	lyleGermanName = nook.Name{
		Language: language.German,
		Value:    "Fred"}

	lyleItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Frodolo"}

	lyleJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ホンマさん"}

	lyleKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "안심해씨"}

	lyleLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sisebuto"}

	lyleRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Лайл"}

	lyleSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿本先生"}

	lyleSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sisebuto"}

	lyleTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿本先生"}
)

var (
	lyleName = nook.Languages{
		language.AmericanEnglish:      lyleAmericanEnglishName,
		language.CanadianFrench:       lyleCanadianFrenchName,
		language.Dutch:                lyleDutchName,
		language.French:               lyleFrenchName,
		language.German:               lyleGermanName,
		language.Italian:              lyleItalianName,
		language.Japanese:             lyleJapaneseName,
		language.Korean:               lyleKoreanName,
		language.LatinAmericanSpanish: lyleLatinAmericanSpanishName,
		language.Russian:              lyleRussianName,
		language.SimplifiedChinese:    lyleSimplifiedChineseName,
		language.Spanish:              lyleSpanishName,
		language.TraditionalChinese:   lyleTraditionalChineseName}
)

var (
	lyleCharacter = nook.Character{
		Animal:   animal.Otter,
		Birthday: lyleBirthday,
		Code:     lyleCode,
		Gender:   gender.Male,
		Name:     lyleName}
)

var (
	Lyle = nook.Resident{
		Character: lyleCharacter}
)
