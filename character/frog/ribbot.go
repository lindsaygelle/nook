package frog

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
		Value:    "Crabot"}

	ribbotDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ribbot"}

	ribbotFrenchName = nook.Name{
		Language: language.French,
		Value:    "Crabot"}

	ribbotGermanName = nook.Name{
		Language: language.German,
		Value:    "Robbi"}

	ribbotItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Ranabot"}

	ribbotJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ガチャ"}

	ribbotKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "철컥"}

	ribbotLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Ranobot"}

	ribbotRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Риббот"}

	ribbotSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "锵锵"}

	ribbotSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Ranobot"}

	ribbotTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "鏘鏘"}
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
		Animal:   animal.Frog,
		Birthday: ribbotBirthday,
		Code:     ribbotCode,
		Key:      character.Ribbot,
		Gender:   gender.Male,
		Name:     ribbotName,
		Special:  false}
)

var (
	ribbotAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "toady"}

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
		language.AmericanEnglish:      ribbotAmericanEnglishPhrase,
		language.CanadianFrench:       ribbotCanadianFrenchPhrase,
		language.Dutch:                ribbotDutchPhrase,
		language.French:               ribbotFrenchPhrase,
		language.German:               ribbotGermanPhrase,
		language.Italian:              ribbotItalianPhrase,
		language.Japanese:             ribbotJapanesePhrase,
		language.Korean:               ribbotKoreanPhrase,
		language.LatinAmericanSpanish: ribbotLatinAmericanSpanishPhrase,
		language.Russian:              ribbotRussianPhrase,
		language.SimplifiedChinese:    ribbotSimplifiedChinesePhrase,
		language.Spanish:              ribbotSpanishPhrase,
		language.TraditionalChinese:   ribbotTraditionalChinesePhrase}
)

var (
	Ribbot = nook.Villager{
		Character:   ribbotCharacter,
		Personality: personality.Jock,
		Phrase:      ribbotPhrase}
)
