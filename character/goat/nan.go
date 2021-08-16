package goat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	nanBirthday = nook.Birthday{
		Day:   24,
		Month: time.August}
)

var (
	nanCode = nook.Code{
		Value: "goa01"}
)

var (
	nanAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nan"}

	nanCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Nana"}

	nanDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Nan"}

	nanFrenchName = nook.Name{
		Language: language.French,
		Value:    "Nana"}

	nanGermanName = nook.Name{
		Language: language.German,
		Value:    "Zenobi"}

	nanItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Nan"}

	nanJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "スミ"}

	nanKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "순이"}

	nanLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pécora"}

	nanRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Нань"}

	nanSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "墨墨"}

	nanSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pécora"}

	nanTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "墨墨"}
)

var (
	nanName = nook.Languages{
		language.AmericanEnglish:      nanAmericanEnglishName,
		language.CanadianFrench:       nanCanadianFrenchName,
		language.Dutch:                nanDutchName,
		language.French:               nanFrenchName,
		language.German:               nanGermanName,
		language.Italian:              nanItalianName,
		language.Japanese:             nanJapaneseName,
		language.Korean:               nanKoreanName,
		language.LatinAmericanSpanish: nanLatinAmericanSpanishName,
		language.Russian:              nanRussianName,
		language.SimplifiedChinese:    nanSimplifiedChineseName,
		language.Spanish:              nanSpanishName,
		language.TraditionalChinese:   nanTraditionalChineseName}
)

var (
	nanCharacter = nook.Character{
		Animal:   animal.Goat,
		Birthday: nanBirthday,
		Code:     nanCode,
		Gender:   gender.Female,
		Name:     nanName}
)

var (
	nanAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "kid"}

	nanCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "fromage"}

	nanDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "geitje"}

	nanFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "fromage"}

	nanGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "mekker"}

	nanItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bimbi"}

	nanJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "っしょ"}

	nanKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "그랬슈"}

	nanLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bibi-bee"}

	nanRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "козленок"}

	nanSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "书"}

	nanSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "peque"}

	nanTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "書"}
)

var (
	nanPhrase = nook.Languages{
		language.AmericanEnglish:      nanAmericanEnglishName,
		language.CanadianFrench:       nanCanadianFrenchName,
		language.Dutch:                nanDutchName,
		language.French:               nanFrenchName,
		language.German:               nanGermanName,
		language.Italian:              nanItalianName,
		language.Japanese:             nanJapaneseName,
		language.Korean:               nanKoreanName,
		language.LatinAmericanSpanish: nanLatinAmericanSpanishName,
		language.Russian:              nanRussianName,
		language.SimplifiedChinese:    nanSimplifiedChineseName,
		language.Spanish:              nanSpanishName,
		language.TraditionalChinese:   nanTraditionalChineseName}
)

var (
	Nan = nook.Villager{
		Character:   nanCharacter,
		Personality: personality.Normal,
		Phrase:      nanPhrase}
)
