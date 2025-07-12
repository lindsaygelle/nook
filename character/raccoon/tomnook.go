package raccoon

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	// tomnookBirthday represents tomnook birthday.
	tomnookBirthday = nook.Birthday{
		Day:   30,
		Month: time.May}
)

var (
	// tomnookCode represents tomnook code.
	tomnookCode = nook.Code{
		Value: "rcn/rco"}
)

var (
	// tomnookAmericanEnglishName represents tomnook american english name.
	tomnookAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tom Nook"}

	// tomnookCanadianFrenchName represents tomnook canadian french name.
	tomnookCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Tom Nook"}

	// tomnookDutchName represents tomnook dutch name.
	tomnookDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tom Nook"}

	// tomnookFrenchName represents tomnook french name.
	tomnookFrenchName = nook.Name{
		Language: language.French,
		Value:    "Tom Nook"}

	// tomnookGermanName represents tomnook german name.
	tomnookGermanName = nook.Name{
		Language: language.German,
		Value:    "Tom Nook"}

	// tomnookItalianName represents tomnook italian name.
	tomnookItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Tom Nook"}

	// tomnookJapaneseName represents tomnook japanese name.
	tomnookJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "たぬきち"}

	// tomnookKoreanName represents tomnook korean name.
	tomnookKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "너굴"}

	// tomnookLatinAmericanSpanishName represents tomnook latin american spanish name.
	tomnookLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tom Nook"}

	// tomnookRussianName represents tomnook russian name.
	tomnookRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Том Нук"}

	// tomnookSimplifiedChineseName represents tomnook simplified chinese name.
	tomnookSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "狸克"}

	// tomnookSpanishName represents tomnook spanish name.
	tomnookSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tom Nook"}

	// tomnookTraditionalChineseName represents tomnook traditional chinese name.
	tomnookTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "狸克"}
)

var (
	// tomnookName represents tomnook name.
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
	// tomnookCharacter represents tomnook character.
	tomnookCharacter = nook.Character{
		Animal:   animal.Raccoon,
		Birthday: tomnookBirthday,
		Code:     tomnookCode,
		Key:      character.TomNook,
		Gender:   gender.Male,
		Name:     tomnookName,
		Special:  true}
)

var (
	// TomNook represents tom nook.
	TomNook = nook.Resident{
		Character: tomnookCharacter}
)
