package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	ribbotBirthday = nook.Birthday{
		Day:   13,
		Month: time.February}
)

var (
	ribbotCode = nook.Code{
		Value: "flg01"}
)

var (
	ribbotAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ribbot"}

	ribbotCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Crabotcroac"}

	ribbotDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ribbotkwwwrrrk"}

	ribbotFrenchName = nook.Name{
		Language: language.French,
		Value:    "Crabotcroac"}

	ribbotGermanName = nook.Name{
		Language: language.German,
		Value:    "Robbiquarrrk"}

	ribbotItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ranabotcracrà"}

	ribbotJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ガチャだロボ"}

	ribbotKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "철컥오버"}

	ribbotLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ranobotcrobit"}

	ribbotRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рибботквак-дзынь"}

	ribbotSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "锵锵机器"}

	ribbotSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ranobotcrobit"}

	ribbotTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鏘鏘機器"}
)

var (
	ribbotName = nook.Languages{
		language.AmericanEnglish:      ribbotAmericanEnglishName,
		language.CanadianFrench:       ribbotCanadianFrenchName,
		language.Dutch:                ribbotDutchName,
		language.French:               ribbotFrenchName,
		language.German:               ribbotGermanName,
		language.Italian:              ribbotItalianName,
		language.Japanese:             ribbotJapaneseName,
		language.Korean:               ribbotKoreanName,
		language.LatinAmericanSpanish: ribbotLatinAmericanSpanishName,
		language.Russian:              ribbotRussianName,
		language.SimplifiedChinese:    ribbotSimplifiedChineseName,
		language.Spanish:              ribbotSpanishName,
		language.TraditionalChinese:   ribbotTraditionalChineseName}
)

var (
	ribbotCharacter = nook.Character{
		Animal:   Frog,
		Birthday: ribbotBirthday,
		Code:     ribbotCode,
		Gender:   nook.Male,
		Name:     ribbotName}
)

var (
	ribbotAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "toadyzzrrbbit"}

	ribbotCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "croac"}

	ribbotDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "kwwwrrrk"}

	ribbotFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "croac"}

	ribbotGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "quarrrk"}

	ribbotItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "cracrà"}

	ribbotJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だロボ"}

	ribbotKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "오버"}

	ribbotLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "crobit"}

	ribbotRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "квак-дзынь"}

	ribbotSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "机器"}

	ribbotSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "crobit"}

	ribbotTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "機器"}
)

var (
	ribbotPhrase = nook.Languages{
		language.AmericanEnglish:      ribbotAmericanEnglishName,
		language.CanadianFrench:       ribbotCanadianFrenchName,
		language.Dutch:                ribbotDutchName,
		language.French:               ribbotFrenchName,
		language.German:               ribbotGermanName,
		language.Italian:              ribbotItalianName,
		language.Japanese:             ribbotJapaneseName,
		language.Korean:               ribbotKoreanName,
		language.LatinAmericanSpanish: ribbotLatinAmericanSpanishName,
		language.Russian:              ribbotRussianName,
		language.SimplifiedChinese:    ribbotSimplifiedChineseName,
		language.Spanish:              ribbotSpanishName,
		language.TraditionalChinese:   ribbotTraditionalChineseName}
)

var (
	Ribbot = nook.Villager{
		Character:   ribbotCharacter,
		Personality: nook.Jock,
		Phrase:      ribbotPhrase}
)
