package raccoon

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	tomnookBirthday = nook.Birthday{
		Day:   30,
		Month: time.May}
)

var (
	tomnookCode = nook.Code{
		Value: "rcn/rco"}
)

var (
	tomnookAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tom Nook"}

	tomnookCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tom Nook"}

	tomnookDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tom Nook"}

	tomnookFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tom Nook"}

	tomnookGermanName = nook.Name{
		Language: language.German,
		Value:    "Tom Nook"}

	tomnookItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tom Nook"}

	tomnookJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "たぬきち"}

	tomnookKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "너굴"}

	tomnookLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tom Nook"}

	tomnookRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Том Нук"}

	tomnookSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "狸克"}

	tomnookSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tom Nook"}

	tomnookTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "狸克"}
)

var (
	tomnookName = nook.Languages{
		language.AmericanEnglish:      tomnookAmericanEnglishName,
		language.CanadianFrench:       tomnookCanadianFrenchName,
		language.Dutch:                tomnookDutchName,
		language.French:               tomnookFrenchName,
		language.German:               tomnookGermanName,
		language.Italian:              tomnookItalianName,
		language.Japanese:             tomnookJapaneseName,
		language.Korean:               tomnookKoreanName,
		language.LatinAmericanSpanish: tomnookLatinAmericanSpanishName,
		language.Russian:              tomnookRussianName,
		language.SimplifiedChinese:    tomnookSimplifiedChineseName,
		language.Spanish:              tomnookSpanishName,
		language.TraditionalChinese:   tomnookTraditionalChineseName}
)

var (
	tomnookCharacter = nook.Character{
		Animal:   animal.Raccoon,
		Birthday: tomnookBirthday,
		Code:     tomnookCode,
		Gender:   gender.Male,
		Name:     tomnookName}
)

var (
	TomNook = nook.Resident{
		Character: tomnookCharacter}
)
