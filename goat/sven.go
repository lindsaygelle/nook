package goat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	svenBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	svenCode = nook.Code{
		Value: ""}
)

var (
	svenAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sven"}

	svenCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	svenDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	svenFrenchName = nook.Name{
		Language: language.French,
		Value:    "Sven"}

	svenGermanName = nook.Name{
		Language: language.German,
		Value:    "Sven"}

	svenItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Giacobbe"}

	svenJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "チャマジ"}

	svenKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	svenLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	svenRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	svenSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "伯伯"}

	svenSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sven"}

	svenTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	svenName = nook.Languages{
		language.AmericanEnglish:      svenAmericanEnglishName,
		language.CanadianFrench:       svenCanadianFrenchName,
		language.Dutch:                svenDutchName,
		language.French:               svenFrenchName,
		language.German:               svenGermanName,
		language.Italian:              svenItalianName,
		language.Japanese:             svenJapaneseName,
		language.Korean:               svenKoreanName,
		language.LatinAmericanSpanish: svenLatinAmericanSpanishName,
		language.Russian:              svenRussianName,
		language.SimplifiedChinese:    svenSimplifiedChineseName,
		language.Spanish:              svenSpanishName,
		language.TraditionalChinese:   svenTraditionalChineseName}
)

var (
	svenCharacter = nook.Character{
		Animal:   Goat,
		Birthday: svenBirthday,
		Code:     svenCode,
		Gender:   nook.Male,
		Name:     svenName}
)

var (
	svenAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	svenCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	svenDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	svenFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "beuh-heu"}

	svenGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "moooh"}

	svenItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "bee eh"}

	svenJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "とほほ"}

	svenKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	svenLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	svenRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	svenSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "活活"}

	svenSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "buuueeeeh"}

	svenTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	svenPhrase = nook.Languages{
		language.AmericanEnglish:      svenAmericanEnglishName,
		language.CanadianFrench:       svenCanadianFrenchName,
		language.Dutch:                svenDutchName,
		language.French:               svenFrenchName,
		language.German:               svenGermanName,
		language.Italian:              svenItalianName,
		language.Japanese:             svenJapaneseName,
		language.Korean:               svenKoreanName,
		language.LatinAmericanSpanish: svenLatinAmericanSpanishName,
		language.Russian:              svenRussianName,
		language.SimplifiedChinese:    svenSimplifiedChineseName,
		language.Spanish:              svenSpanishName,
		language.TraditionalChinese:   svenTraditionalChineseName}
)

var (
	Sven = nook.Villager{
		Character:   svenCharacter,
		Personality: nook.Lazy,
		Phrase:      svenPhrase}
)
