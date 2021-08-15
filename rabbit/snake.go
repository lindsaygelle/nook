package rabbit

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Jeannotmon lapin"}

	snakeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Snakeninja"}

	snakeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jeannotmon lapin"}

	snakeGermanName = nook.Name{
		Language: language.German,
		Value:    "Philiphopplahopp"}

	snakeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Razzobun bun"}

	snakeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モモチニンニン"}

	snakeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "닌토닌닌"}

	snakeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rayoboin-boing"}

	snakeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Снейкзаяссс"}

	snakeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "百川忍忍"}

	snakeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rayoboin-boing"}

	snakeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "百川忍忍"}
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
		Animal:   Rabbit,
		Birthday: snakeBirthday,
		Code:     snakeCode,
		Gender:   nook.Male,
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
	Snake = nook.Villager{
		Character:   snakeCharacter,
		Personality: nook.Jock,
		Phrase:      snakePhrase}
)
