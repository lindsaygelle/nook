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
	// snakeBirthday represents Snake's birthday.
	snakeBirthday = nook.Birthday{
		Day:   3,
		Month: time.November}
)

var (
	// snakeCode represents Snake's unique code.
	snakeCode = nook.Code{
		Value: "rbt03"}
)

var (
	// snakeAmericanEnglishName represents Snake's name in American English.
	snakeAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Snake"}

	// snakeCanadianFrenchName represents Snake's name in Canadian French.
	snakeCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Jeannot"}

	// snakeDutchName represents Snake's name in Dutch.
	snakeDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Snake"}

	// snakeFrenchName represents Snake's name in French.
	snakeFrenchName = nook.Name{
		Language: language.French,
		Value:    "Jeannot"}

	// snakeGermanName represents Snake's name in German.
	snakeGermanName = nook.Name{
		Language: language.German,
		Value:    "Philip"}

	// snakeItalianName represents Snake's name in Italian.
	snakeItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Razzo"}

	// snakeJapaneseName represents Snake's name in Japanese.
	snakeJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "モモチ"}

	// snakeKoreanName represents Snake's name in Korean.
	snakeKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "닌토"}

	// snakeLatinAmericanSpanishName represents Snake's name in Latin American Spanish.
	snakeLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Rayo"}

	// snakeRussianName represents Snake's name in Russian.
	snakeRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Снейк"}

	// snakeSimplifiedChineseName represents Snake's name in Simplified Chinese.
	snakeSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "百川"}

	// snakeSpanishName represents Snake's name in Spanish.
	snakeSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Rayo"}

	// snakeTraditionalChineseName represents Snake's name in Traditional Chinese.
	snakeTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "百川"}
)

var (
	// snakeName represents Snake's name in different languages.
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
	// snakeCharacter represents Snake's character information.
	snakeCharacter = nook.Character{
		Animal:   animal.Rabbit,
		Birthday: snakeBirthday,
		Code:     snakeCode,
		Key:      character.Snake,
		Gender:   gender.Male,
		Name:     snakeName,
		Special:  false}
)

var (
	// snakeAmericanEnglishPhrase represents Snake's phrase in American English.
	snakeAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bunyip"}

	// snakeCanadianFrenchPhrase represents Snake's phrase in Canadian French.
	snakeCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon lapin"}

	// snakeDutchPhrase represents Snake's phrase in Dutch.
	snakeDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "ninja"}

	// snakeFrenchPhrase represents Snake's phrase in French.
	snakeFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon lapin"}

	// snakeGermanPhrase represents Snake's phrase in German.
	snakeGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "hopplahopp"}

	// snakeItalianPhrase represents Snake's phrase in Italian.
	snakeItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bun bun"}

	// snakeJapanesePhrase represents Snake's phrase in Japanese.
	snakeJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ニンニン"}

	// snakeKoreanPhrase represents Snake's phrase in Korean.
	snakeKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "닌닌"}

	// snakeLatinAmericanSpanishPhrase represents Snake's phrase in Latin American Spanish.
	snakeLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "boin-boing"}

	// snakeRussianPhrase represents Snake's phrase in Russian.
	snakeRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "заяссс"}

	// snakeSimplifiedChinesePhrase represents Snake's phrase in Simplified Chinese.
	snakeSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "忍忍"}

	// snakeSpanishPhrase represents Snake's phrase in Spanish.
	snakeSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "boin-boing"}

	// snakeTraditionalChinesePhrase represents Snake's phrase in Traditional Chinese.
	snakeTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "忍忍"}
)

var (
	// snakePhrase represents Snake's phrases in different languages.
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
	// Snake represents the character Snake with his complete information.
	Snake = nook.Villager{
		Character:   snakeCharacter,
		Personality: personality.Jock,
		Phrase:      snakePhrase}
)
