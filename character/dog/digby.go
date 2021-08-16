package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	digbyBirthday = nook.Birthday{
		Day:   20,
		Month: time.December}
)

var (
	digbyCode = nook.Code{
		Value: "szo"}
)

var (
	digbyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Digby"}

	digbyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Max"}

	digbyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Digby"}

	digbyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Max"}

	digbyGermanName = nook.Name{
		Language: language.German,
		Value:    "Moritz"}

	digbyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Fofò"}

	digbyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ケント"}

	digbyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "켄트"}

	digbyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Candrés"}

	digbyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дигби"}

	digbySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "西施德"}

	digbySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Candrés"}

	digbyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "西施德"}
)

var (
	digbyName = nook.Languages{
		language.AmericanEnglish:      digbyAmericanEnglishName,
		language.CanadianFrench:       digbyCanadianFrenchName,
		language.Dutch:                digbyDutchName,
		language.French:               digbyFrenchName,
		language.German:               digbyGermanName,
		language.Italian:              digbyItalianName,
		language.Japanese:             digbyJapaneseName,
		language.Korean:               digbyKoreanName,
		language.LatinAmericanSpanish: digbyLatinAmericanSpanishName,
		language.Russian:              digbyRussianName,
		language.SimplifiedChinese:    digbySimplifiedChineseName,
		language.Spanish:              digbySpanishName,
		language.TraditionalChinese:   digbyTraditionalChineseName}
)

var (
	digbyCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: digbyBirthday,
		Code:     digbyCode,
		Gender:   gender.Male,
		Name:     digbyName}
)

var (
	Digby = nook.Resident{
		Character: digbyCharacter}
)
