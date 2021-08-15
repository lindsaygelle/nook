package chicken

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	knoxBirthday = nook.Birthday{
		Day:   23,
		Month: time.November}
)

var (
	knoxCode = nook.Code{
		Value: "chn11"}
)

var (
	knoxAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Knox"}

	knoxCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Wolfram"}

	knoxDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Knox"}

	knoxFrenchName = nook.Name{
		Language: language.French,
		Value:    "Wolfram"}

	knoxGermanName = nook.Name{
		Language: language.German,
		Value:    "Tschiwi"}

	knoxItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Kalin"}

	knoxJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "キンカク"}

	knoxKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "금끼오"}

	knoxLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Tono"}

	knoxRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Нокс"}

	knoxSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "金阁"}

	knoxSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Tono"}

	knoxTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "金閣"}
)

var (
	knoxName = nook.Languages{
		language.AmericanEnglish:      knoxAmericanEnglishName,
		language.CanadianFrench:       knoxCanadianFrenchName,
		language.Dutch:                knoxDutchName,
		language.French:               knoxFrenchName,
		language.German:               knoxGermanName,
		language.Italian:              knoxItalianName,
		language.Japanese:             knoxJapaneseName,
		language.Korean:               knoxKoreanName,
		language.LatinAmericanSpanish: knoxLatinAmericanSpanishName,
		language.Russian:              knoxRussianName,
		language.SimplifiedChinese:    knoxSimplifiedChineseName,
		language.Spanish:              knoxSpanishName,
		language.TraditionalChinese:   knoxTraditionalChineseName}
)

var (
	knoxCharacter = nook.Character{
		Animal:   Chicken,
		Birthday: knoxBirthday,
		Code:     knoxCode,
		Gender:   nook.Male,
		Name:     knoxName}
)

var (
	knoxAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "cluckling"}

	knoxCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "ra-ta-tak"}

	knoxDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "nestei"}

	knoxFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "ra-ta-tak"}

	knoxGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knusprig"}

	knoxItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "uovole"}

	knoxJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "せいッ"}

	knoxKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "하앗"}

	knoxLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "cocorocó"}

	knoxRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "короко"}

	knoxSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "等待"}

	knoxSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "mazorca"}

	knoxTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "等待"}
)

var (
	knoxPhrase = nook.Languages{
		language.AmericanEnglish:      knoxAmericanEnglishName,
		language.CanadianFrench:       knoxCanadianFrenchName,
		language.Dutch:                knoxDutchName,
		language.French:               knoxFrenchName,
		language.German:               knoxGermanName,
		language.Italian:              knoxItalianName,
		language.Japanese:             knoxJapaneseName,
		language.Korean:               knoxKoreanName,
		language.LatinAmericanSpanish: knoxLatinAmericanSpanishName,
		language.Russian:              knoxRussianName,
		language.SimplifiedChinese:    knoxSimplifiedChineseName,
		language.Spanish:              knoxSpanishName,
		language.TraditionalChinese:   knoxTraditionalChineseName}
)

var (
	Knox = nook.Villager{
		Character:   knoxCharacter,
		Personality: nook.Cranky,
		Phrase:      knoxPhrase}
)
