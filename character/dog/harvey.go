package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	harveyBirthday = nook.Birthday{
		Day:   2,
		Month: time.August}
)

var (
	harveyCode = nook.Code{
		Value: "spn"}
)

var (
	harveyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Harvey"}

	harveyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Joe"}

	harveyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Harvey"}

	harveyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Joe"}

	harveyGermanName = nook.Name{
		Language: language.German,
		Value:    "Harvey"}

	harveyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fiorilio"}

	harveyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パニエル"}

	harveyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파니엘"}

	harveyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fauno"}

	harveyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Харви"}

	harveySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "巴猎"}

	harveySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fauno"}

	harveyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "巴獵"}
)

var (
	harveyName = nook.Languages{
		language.AmericanEnglish:      harveyAmericanEnglishName,
		language.CanadianFrench:       harveyCanadianFrenchName,
		language.Dutch:                harveyDutchName,
		language.French:               harveyFrenchName,
		language.German:               harveyGermanName,
		language.Italian:              harveyItalianName,
		language.Japanese:             harveyJapaneseName,
		language.Korean:               harveyKoreanName,
		language.LatinAmericanSpanish: harveyLatinAmericanSpanishName,
		language.Russian:              harveyRussianName,
		language.SimplifiedChinese:    harveySimplifiedChineseName,
		language.Spanish:              harveySpanishName,
		language.TraditionalChinese:   harveyTraditionalChineseName}
)

var (
	harveyCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: harveyBirthday,
		Code:     harveyCode,
		Gender:   gender.Male,
		Name:     harveyName}
)

var (
	Harvey = nook.Resident{
		Character: harveyCharacter}
)
