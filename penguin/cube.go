package penguin

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	cubeBirthday = nook.Birthday{
		Day:   29,
		Month: time.January}
)

var (
	cubeCode = nook.Code{
		Value: "pgn02"}
)

var (
	cubeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Cube"}

	cubeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Cube"}

	cubeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Cube"}

	cubeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cube"}

	cubeGermanName = nook.Name{
		Language: language.German,
		Value:    "Cube"}

	cubeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Cubetto"}

	cubeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ビス"}

	cubeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "빙수"}

	cubeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Cube"}

	cubeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кьюб"}

	cubeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "罗斯"}

	cubeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Cube"}

	cubeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "羅斯"}
)

var (
	cubeName = nook.Languages{
		language.AmericanEnglish:      cubeAmericanEnglishName,
		language.CanadianFrench:       cubeCanadianFrenchName,
		language.Dutch:                cubeDutchName,
		language.French:               cubeFrenchName,
		language.German:               cubeGermanName,
		language.Italian:              cubeItalianName,
		language.Japanese:             cubeJapaneseName,
		language.Korean:               cubeKoreanName,
		language.LatinAmericanSpanish: cubeLatinAmericanSpanishName,
		language.Russian:              cubeRussianName,
		language.SimplifiedChinese:    cubeSimplifiedChineseName,
		language.Spanish:              cubeSpanishName,
		language.TraditionalChinese:   cubeTraditionalChineseName}
)

var (
	cubeCharacter = nook.Character{
		Animal:   Penguin,
		Birthday: cubeBirthday,
		Code:     cubeCode,
		Gender:   nook.Male,
		Name:     cubeName}
)

var (
	cubeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "d-d-dude"}

	cubeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "cococopain"}

	cubeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "k-k-kerel"}

	cubeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "cococopain"}

	cubeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "d-d-dude"}

	cubeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "brivido"}

	cubeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ペンペン"}

	cubeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "땡땡"}

	cubeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "waap"}

	cubeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "ч-чубрик"}

	cubeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "十字"}

	cubeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "waap"}

	cubeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "十字"}
)

var (
	cubePhrase = nook.Languages{
		language.AmericanEnglish:      cubeAmericanEnglishName,
		language.CanadianFrench:       cubeCanadianFrenchName,
		language.Dutch:                cubeDutchName,
		language.French:               cubeFrenchName,
		language.German:               cubeGermanName,
		language.Italian:              cubeItalianName,
		language.Japanese:             cubeJapaneseName,
		language.Korean:               cubeKoreanName,
		language.LatinAmericanSpanish: cubeLatinAmericanSpanishName,
		language.Russian:              cubeRussianName,
		language.SimplifiedChinese:    cubeSimplifiedChineseName,
		language.Spanish:              cubeSpanishName,
		language.TraditionalChinese:   cubeTraditionalChineseName}
)

var (
	Cube = nook.Villager{
		Character:   cubeCharacter,
		Personality: nook.Lazy,
		Phrase:      cubePhrase}
)
