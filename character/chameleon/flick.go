package chameleon

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"golang.org/x/text/language"
)

var (
	flickBirthday = nook.Birthday{
		Day:   10,
		Month: time.May}
)

var (
	flickCode = nook.Code{
		Value: "chy"}
)

var (
	flickAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Flick"}

	flickCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Djason"}

	flickDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Flick"}

	flickFrenchName = nook.Name{
		Language: language.French,
		Value:    "Djason"}

	flickGermanName = nook.Name{
		Language: language.German,
		Value:    "Carlson"}

	flickItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ivano"}

	flickJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レックス"}

	flickKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "레온"}

	flickLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Kamilo"}

	flickRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Флик"}

	flickSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "龙克斯"}

	flickSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Kamilo"}

	flickTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "龍克斯"}
)

var (
	flickName = nook.Languages{
		language.AmericanEnglish:      flickAmericanEnglishName,
		language.CanadianFrench:       flickCanadianFrenchName,
		language.Dutch:                flickDutchName,
		language.French:               flickFrenchName,
		language.German:               flickGermanName,
		language.Italian:              flickItalianName,
		language.Japanese:             flickJapaneseName,
		language.Korean:               flickKoreanName,
		language.LatinAmericanSpanish: flickLatinAmericanSpanishName,
		language.Russian:              flickRussianName,
		language.SimplifiedChinese:    flickSimplifiedChineseName,
		language.Spanish:              flickSpanishName,
		language.TraditionalChinese:   flickTraditionalChineseName}
)

var (
	flickCharacter = nook.Character{
		Animal:   animal.Chameleon,
		Birthday: flickBirthday,
		Code:     flickCode,
		Gender:   gender.Male,
		Name:     flickName}
)

var (
	Flick = nook.Resident{
		Character: flickCharacter}
)
