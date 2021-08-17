package goat

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
		Value:    ""}

	svenDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

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
		Value:    ""}

	svenLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	svenRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	svenSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "伯伯"}

	svenSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sven"}

	svenTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
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
		Animal:   animal.Goat,
		Birthday: svenBirthday,
		Code:     svenCode,
		Key:      character.Sven,
		Gender:   gender.Male,
		Name:     svenName}
)

var (
	svenAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "buh-uh-ud"}

	svenCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	svenDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

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
		Value:    ""}

	svenLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	svenRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	svenSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "活活"}

	svenSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "buuueeeeh"}

	svenTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	svenPhrase = nook.Languages{
		language.AmericanEnglish:      svenAmericanEnglishPhrase,
		language.CanadianFrench:       svenCanadianFrenchPhrase,
		language.Dutch:                svenDutchPhrase,
		language.French:               svenFrenchPhrase,
		language.German:               svenGermanPhrase,
		language.Italian:              svenItalianPhrase,
		language.Japanese:             svenJapanesePhrase,
		language.Korean:               svenKoreanPhrase,
		language.LatinAmericanSpanish: svenLatinAmericanSpanishPhrase,
		language.Russian:              svenRussianPhrase,
		language.SimplifiedChinese:    svenSimplifiedChinesePhrase,
		language.Spanish:              svenSpanishPhrase,
		language.TraditionalChinese:   svenTraditionalChinesePhrase}
)

var (
	Sven = nook.Villager{
		Character:   svenCharacter,
		Personality: personality.Lazy,
		Phrase:      svenPhrase}
)
