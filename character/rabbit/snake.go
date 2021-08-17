package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	snakeBirthday = nook.Birthday{
		Day:   3,
		Month: time.November}
)

var (
	snakeCode = nook.Code{
		Value: "rbt03"}
)

var (
	snakeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Snake"}

	snakeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jeannot"}

	snakeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Snake"}

	snakeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jeannot"}

	snakeGermanName = nook.Name{
		Language: language.German,
		Value:    "Philip"}

	snakeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Razzo"}

	snakeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モモチ"}

	snakeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "닌토"}

	snakeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rayo"}

	snakeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Снейк"}

	snakeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "百川"}

	snakeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rayo"}

	snakeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "百川"}
)

var (
	snakeName = nook.Languages{
		language.AmericanEnglish:      snakeAmericanEnglishName,
		language.CanadianFrench:       snakeCanadianFrenchName,
		language.Dutch:                snakeDutchName,
		language.French:               snakeFrenchName,
		language.German:               snakeGermanName,
		language.Italian:              snakeItalianName,
		language.Japanese:             snakeJapaneseName,
		language.Korean:               snakeKoreanName,
		language.LatinAmericanSpanish: snakeLatinAmericanSpanishName,
		language.Russian:              snakeRussianName,
		language.SimplifiedChinese:    snakeSimplifiedChineseName,
		language.Spanish:              snakeSpanishName,
		language.TraditionalChinese:   snakeTraditionalChineseName}
)

var (
	snakeCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: snakeBirthday,
		Code:     snakeCode,
		Key:      character.Snake,
		Gender:   gender.Male,
		Name:     snakeName}
)

var (
	snakeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bunyip"}

	snakeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon lapin"}

	snakeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ninja"}

	snakeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon lapin"}

	snakeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hopplahopp"}

	snakeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bun bun"}

	snakeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ニンニン"}

	snakeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "닌닌"}

	snakeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "boin-boing"}

	snakeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "заяссс"}

	snakeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "忍忍"}

	snakeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "boin-boing"}

	snakeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "忍忍"}
)

var (
	snakePhrase = nook.Languages{
		language.AmericanEnglish:      snakeAmericanEnglishPhrase,
		language.CanadianFrench:       snakeCanadianFrenchPhrase,
		language.Dutch:                snakeDutchPhrase,
		language.French:               snakeFrenchPhrase,
		language.German:               snakeGermanPhrase,
		language.Italian:              snakeItalianPhrase,
		language.Japanese:             snakeJapanesePhrase,
		language.Korean:               snakeKoreanPhrase,
		language.LatinAmericanSpanish: snakeLatinAmericanSpanishPhrase,
		language.Russian:              snakeRussianPhrase,
		language.SimplifiedChinese:    snakeSimplifiedChinesePhrase,
		language.Spanish:              snakeSpanishPhrase,
		language.TraditionalChinese:   snakeTraditionalChinesePhrase}
)

var (
	Snake = nook.Villager{
		Character:   snakeCharacter,
		Personality: personality.Jock,
		Phrase:      snakePhrase}
)
