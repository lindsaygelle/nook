package duck

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	deenaBirthday = nook.Birthday{
		Day:   27,
		Month: time.June}
)

var (
	deenaCode = nook.Code{
		Value: "duk04"}
)

var (
	deenaAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Deena"}

	deenaCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mina"}

	deenaDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Deena"}

	deenaFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mina"}

	deenaGermanName = nook.Name{
		Language: language.German,
		Value:    "Olivia"}

	deenaItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Dina"}

	deenaJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "まりも"}

	deenaKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "마리모"}

	deenaLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Martita"}

	deenaRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Дина"}

	deenaSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "球藻"}

	deenaSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Martita"}

	deenaTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "毬藻"}
)

var (
	deenaName = nook.Languages{
		language.AmericanEnglish:      deenaAmericanEnglishName,
		language.CanadianFrench:       deenaCanadianFrenchName,
		language.Dutch:                deenaDutchName,
		language.French:               deenaFrenchName,
		language.German:               deenaGermanName,
		language.Italian:              deenaItalianName,
		language.Japanese:             deenaJapaneseName,
		language.Korean:               deenaKoreanName,
		language.LatinAmericanSpanish: deenaLatinAmericanSpanishName,
		language.Russian:              deenaRussianName,
		language.SimplifiedChinese:    deenaSimplifiedChineseName,
		language.Spanish:              deenaSpanishName,
		language.TraditionalChinese:   deenaTraditionalChineseName}
)

var (
	deenaCharacter = nook.Character{
		Animal:   Duck,
		Birthday: deenaBirthday,
		Code:     deenaCode,
		Gender:   nook.Female,
		Name:     deenaName}
)

var (
	deenaAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "woowoo"}

	deenaCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "youhou"}

	deenaDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "joehoe"}

	deenaFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "youhou"}

	deenaGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "juhuu"}

	deenaItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quiproquò"}

	deenaJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "マル"}

	deenaKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "우우"}

	deenaLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "yiuju"}

	deenaRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "кря-кря"}

	deenaSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "球球"}

	deenaSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "miguita"}

	deenaTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "毬毬"}
)

var (
	deenaPhrase = nook.Languages{
		language.AmericanEnglish:      deenaAmericanEnglishName,
		language.CanadianFrench:       deenaCanadianFrenchName,
		language.Dutch:                deenaDutchName,
		language.French:               deenaFrenchName,
		language.German:               deenaGermanName,
		language.Italian:              deenaItalianName,
		language.Japanese:             deenaJapaneseName,
		language.Korean:               deenaKoreanName,
		language.LatinAmericanSpanish: deenaLatinAmericanSpanishName,
		language.Russian:              deenaRussianName,
		language.SimplifiedChinese:    deenaSimplifiedChineseName,
		language.Spanish:              deenaSpanishName,
		language.TraditionalChinese:   deenaTraditionalChineseName}
)

var (
	Deena = nook.Villager{
		Character:   deenaCharacter,
		Personality: nook.Normal,
		Phrase:      deenaPhrase}
)
