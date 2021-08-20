package skunk

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"

	"golang.org/x/text/language"
)

var (
	kicksBirthday = nook.Birthday{
		Day:   30,
		Month: time.November}
)

var (
	kicksCode = nook.Code{
		Value: "skk"}
)

var (
	kicksAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Kicks"}

	kicksCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Blaise"}

	kicksDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Kicks"}

	kicksFrenchName = nook.Name{
		Language: language.French,
		Value:    "Blaise"}

	kicksGermanName = nook.Name{
		Language: language.German,
		Value:    "Schubert"}

	kicksItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sciuscià"}

	kicksJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シャンク"}

	kicksKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "패트릭"}

	kicksLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Betunio"}

	kicksRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кикс"}

	kicksSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "薛革"}

	kicksSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Betunio"}

	kicksTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "薛革"}
)

var (
	kicksName = nook.Languages{
		language.AmericanEnglish:      kicksAmericanEnglishName,
		language.CanadianFrench:       kicksCanadianFrenchName,
		language.Dutch:                kicksDutchName,
		language.French:               kicksFrenchName,
		language.German:               kicksGermanName,
		language.Italian:              kicksItalianName,
		language.Japanese:             kicksJapaneseName,
		language.Korean:               kicksKoreanName,
		language.LatinAmericanSpanish: kicksLatinAmericanSpanishName,
		language.Russian:              kicksRussianName,
		language.SimplifiedChinese:    kicksSimplifiedChineseName,
		language.Spanish:              kicksSpanishName,
		language.TraditionalChinese:   kicksTraditionalChineseName}
)

var (
	kicksCharacter = nook.Character{
		Animal:   animal.Skunk,
		Birthday: kicksBirthday,
		Code:     kicksCode,
		Key:      character.Kicks,
		Gender:   gender.Male,
		Name:     kicksName,
		Special:  true}
)

var (
	Kicks = nook.Resident{
		Character: kicksCharacter}
)
